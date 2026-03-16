package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/optimus/backend/config"
	"github.com/optimus/backend/dtos"
	"github.com/optimus/backend/models"
	"github.com/optimus/backend/utils"
)

func GetAllUser(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Role").Preload("PokjaGroup").Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data",
		"data":    users,
	})
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func CreateUser(c *gin.Context) {
	var req dtos.CreateUserRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var existing models.User
	err = config.DB.Where("email = ?", req.Email).First(&existing).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email sudah digunakan pengguna lain!",
		})
		return
	}

	skFile, _ := c.FormFile("sk_file")
	pbjFile, _ := c.FormFile("pbj_file")
	competenceFile, _ := c.FormFile("competence_file")
	photoFile, _ := c.FormFile("file_photo")

	uploadDir := "uploads/users"
	_ = os.MkdirAll(uploadDir, os.ModePerm)

	saveUploaded := func(file *multipart.FileHeader) *string {
		if file == nil {
			return nil
		}

		filename := uuid.New().String() + "_" + filepath.Base(file.Filename)
		path := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, path); err != nil {
			return nil
		}
		return &path
	}

	skPath := saveUploaded(skFile)
	pbjPath := saveUploaded(pbjFile)
	competencePath := saveUploaded(competenceFile)
	photoPath := saveUploaded(photoFile)

	if !isValidPassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password minimal 8 karakter dan harus mengandung huruf besar, huruf kecil, angka, dan simbol.",
		})
		return
	}

	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: utils.HashSHA512(req.Password),
		IsActive: req.IsActive,

		Nik:           req.Nik,
		Nip:           req.Nip,
		Group:         req.Group,
		RoleId:        req.RoleId,
		PokjaGroupsId: req.PokjaGroupsId,

		PhoneNumber:     req.PhoneNumber,
		OpdOrganization: req.OpdOrganization,

		SkNumber:         req.SkNumber,
		SkFile:           skPath,
		PbjNumber:        req.PbjNumber,
		PbjFile:          pbjPath,
		CompetenceNumber: req.CompetenceNumber,
		CompetenceFile:   competencePath,
		PhotoFile:        photoPath,
		SatkerCode:       req.SatkerCode,
		GpId:             req.GpId,
		Address:          req.Address,
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var createdUser models.User

	if err := config.DB.
		Preload("Role").First(&createdUser, user.Id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Membuat data berhasil",
		"data":    createdUser,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	err := config.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Mengambil data gagal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	var req dtos.CreateUserRequest
	id := c.Param("id")

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	if req.Email != "" {
		var count int64
		config.DB.Model(&models.User{}).
			Where("email = ? AND id <> ?", req.Email, id).
			Count(&count)

		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Email sudah digunakan!",
			})
			return
		}
	}

	saveUploaded := func(field string) *string {
		file, err := c.FormFile(field)
		if err != nil {
			return nil
		}

		_ = os.MkdirAll("uploads/users", os.ModePerm)
		filename := uuid.New().String() + "_" + filepath.Base(file.Filename)
		path := filepath.Join("uploads/users", filename)

		if err := c.SaveUploadedFile(file, path); err != nil {
			return nil
		}
		return &path
	}

	if p := saveUploaded("sk_file"); p != nil {
		user.SkFile = p
	}
	if p := saveUploaded("pbj_file"); p != nil {
		user.PbjFile = p
	}
	if p := saveUploaded("competence_file"); p != nil {
		user.CompetenceFile = p
	}
	if p := saveUploaded("file_photo"); p != nil {
		user.PhotoFile = p
	}

	user.FullName = req.FullName
	user.Email = req.Email
	user.IsActive = req.IsActive
	user.Nik = req.Nik
	user.Nip = req.Nip
	user.Group = req.Group
	user.RoleId = req.RoleId
	user.PokjaGroupsId = req.PokjaGroupsId
	user.PhoneNumber = req.PhoneNumber
	user.OpdOrganization = req.OpdOrganization
	user.SkNumber = req.SkNumber
	user.PbjNumber = req.PbjNumber
	user.CompetenceNumber = req.CompetenceNumber
	user.SatkerCode = req.SatkerCode
	user.GpId = req.GpId
	user.Address = req.Address

	if req.Password != "" {
		user.Password = utils.HashSHA512(req.Password)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	config.DB.Preload("Role").First(&user, user.Id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengubah data berhasil",
		"data":    user,
	})
}

func UpdateStatus(c *gin.Context) {
	id := c.Param("user_id")

	var user models.User
	config.DB.First(&user, id)

	if user.IsActive == nil {
		active := true
		user.IsActive = &active
	} else {
		*user.IsActive = !*user.IsActive
	}

	err := config.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Mengubah status gagal!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status berhasil diubah!",
		"data":    user,
	})
}

func RequestResetPassword(c *gin.Context) {
	db := config.DB

	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email tidak valid",
		})
		return
	}

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Jika email terdaftar, link reset akan dikirim",
		})
		return
	}

	token, err := utils.GenerateResetToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal generate token",
		})
		return
	}

	origin := c.GetHeader("Origin")
	resetURL := fmt.Sprintf(
		"%s/reset-kata-sandi?token=%s&email=%s",
		origin,
		token,
		url.QueryEscape(user.Email),
	)

	emailBody := `
<table style="width: 100%; border-collapse: collapse; background: linear-gradient(135deg, #ff6600 0%, #ff8c42 100%); padding: 40px 20px;" role="presentation">
<tbody>
<tr>
<td align="center">
<table style="max-width: 600px; width: 100%; border-collapse: collapse; background: #ffffff; border-radius: 20px; overflow: hidden; box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3); font-family: system-ui, sans-serif, Arial;" role="presentation">
<tbody>
<tr>
<td style="background: linear-gradient(135deg, #ff6600 0%, #ff8c42 100%); padding: 50px 40px; text-align: center;">
<h1 style="margin: 0; color: #ffffff; font-size: 32px; font-weight: bold; letter-spacing: -0.5px;">Monev Monalisa</h1>
<p style="margin: 10px 0 0; color: rgba(255, 255, 255, 0.9); font-size: 14px; font-weight: 500;">Sistem Monitoring &amp; Evaluasi Proyek</p>
</td>
</tr>
<tr>
<td style="padding: 50px 40px;">
<h2 style="margin: 0 0 20px; color: #1a202c; font-size: 28px; font-weight: bold;">Reset Kata Sandi Anda</h2>
<p style="margin: 0 0 20px; color: #4a5568; font-size: 16px; line-height: 1.6;">Halo,</p>
<p style="margin: 0 0 20px; color: #4a5568; font-size: 16px; line-height: 1.6;">Kami menerima permintaan untuk mereset kata sandi akun&nbsp;<strong>Monev Monalisa</strong> Anda. Klik tombol di bawah ini untuk membuat password baru.</p>
<table style="width: 100%; margin: 30px 0;" role="presentation">
<tbody>
<tr>
<td align="center"><a style="display: inline-block; padding: 16px 40px; background: linear-gradient(135deg, #ff6600 0%, #ff8c42 100%); color: #ffffff; text-decoration: none; border-radius: 50px; font-size: 16px; font-weight: 600; box-shadow: 0 10px 25px rgba(255, 102, 0, 0.4); transition: all 0.3s ease;" href="` + resetURL + `"> Reset Password </a></td>
</tr>
</tbody>
</table>
<div style="background: #fff7ed; border-left: 4px solid #ff6600; padding: 20px; margin: 30px 0; border-radius: 8px;">
<p style="margin: 0 0 10px; color: #2d3748; font-size: 14px; font-weight: 600;">Catatan Keamanan:</p>
<p style="margin: 0; color: #4a5568; font-size: 14px; line-height: 1.5;">Link reset password ini akan kadaluarsa dalam <strong>24 jam</strong> untuk alasan keamanan. Jika Anda tidak meminta reset ini, abaikan email ini.</p>
</div>
</td>
</tr>
<tr>
<td style="background: #fef3c7; padding: 30px 40px; text-align: center; border-top: 1px solid #fde68a;">
<p style="margin: 20px 0 0; color: #b45309; font-size: 12px;">&copy; 2026 Monev Monalisa. Hak Cipta Dilindungi.</p>
<p style="margin: 10px 0 0; color: #d97706; font-size: 11px;">Ini adalah pesan otomatis, mohon jangan membalas email ini.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
`

	err = utils.SendEmail(user.Email, "Reset Password", emailBody)
	if err != nil {
		fmt.Println("SMTP ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengirim email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Jika email terdaftar, link reset akan dikirim",
	})
}

func ResetPassword(c *gin.Context) {
	var req struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request tidak valid"})
		return
	}

	userID, err := utils.ParseJWT(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User tidak ditemukan"})
		return
	}

	user.Password = utils.HashSHA512(req.Password)

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
}
