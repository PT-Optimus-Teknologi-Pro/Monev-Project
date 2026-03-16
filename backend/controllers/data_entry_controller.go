package controllers

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/optimus/backend/config"
	"github.com/optimus/backend/dtos"
	"github.com/optimus/backend/models"
)

func GetAllDataEntry(c *gin.Context) {
	var data []models.DataEntry
	config.DB.
		Preload("User").
		Preload("User.Role").
		Preload("SelectedPpk").
		Preload("SelectedPpk.Role").
		Preload("User.PokjaGroup").
		Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    data,
	})
}

func GetDataEntryById(c *gin.Context) {
	idParam := c.Param("id")

	var data models.DataEntry
	result := config.DB.
		Preload("User").
		Preload("User.Role").
		Preload("SelectedPpk").
		Preload("SelectedPpk.Role").
		Preload("User.PokjaGroup").
		First(&data, idParam)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    data,
	})
}

func CreateDataEntry(c *gin.Context) {
	var req dtos.CreateDataEntryRequest
	userId, isNull := c.Get("user_id")

	if !isNull {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Pengguna harus login terlebih dahulu!",
		})
		return
	}

	var user models.User
	config.DB.First(&user, userId)

	err := c.ShouldBindWith(&req, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
		return
	}

	BuktiFile, _ := c.FormFile("bukti_file")

	uploadDir := "uploads/entry"
	_ = os.MkdirAll(uploadDir, os.ModePerm)

	saveUpload := func(file *multipart.FileHeader) *string {
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

	BuktiPath := saveUpload(BuktiFile)

	data := models.DataEntry{
		Tipe:            req.Tipe,
		JenisPengadaan:  req.JenisPengadaan,
		MetodePengadaan: req.MetodePengadaan,
		KodePaket:       req.KodePaket,
		KodeRup:         req.KodeRup,
		TahunAnggaran:   req.TahunAnggaran,
		TanggalMasuk:    req.TanggalMasuk,
		SatuanKerja:     req.SatuanKerja,
		NamaPaket:       req.NamaPaket,
		SumberDana:      req.SumberDana,

		StatusPaket:      req.StatusPaket,
		StatusPengiriman: req.StatusPengiriman,

		NilaiPagu: req.NilaiPagu,
		NilaiHps:  req.NilaiHps,

		NomorKontrak:   req.NomorKontrak,
		NilaiKontrak:   req.NilaiKontrak,
		TanggalKontrak: req.TanggalKontrak,
		NamaPpk:        req.NamaPpk,
		JabatanPpk:     req.JabatanPpk,
		Pemasukan:      req.Pemasukan,
		Pendaftar:      req.Pendaftar,

		NamaPimpinanPerusahaan: req.NamaPimpinanPerusahaan,
		JabatanPimpinan:        req.JabatanPimpinan,

		Pemenang:       req.Pemenang,
		NilaiPenawaran: req.NilaiPenawaran,
		NilaiNegosiasi: req.NilaiNegosiasi,
		NomorTelp:      req.NomorTelp,
		Email:          req.Email,
		Npwp:           req.Npwp,

		AlamatPemenang:  req.AlamatPemenang,
		LokasiPekerjaan: req.LokasiPekerjaan,

		Catatan:   req.Catatan,
		BuktiFile: BuktiPath,

		SelectedPpkId: req.SelectedPpkId,
		UserId:        user.Id,
	}

	err = config.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Membuat data gagal!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Membuat data berhasil",
		"data":    data,
	})
}

func UpdateDataEntry(c *gin.Context) {
	var req dtos.CreateDataEntryRequest
	id := c.Param("id")

	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
		return
	}
	BuktiFile, _ := c.FormFile("bukti_file")

	uploadDir := "uploads/entry"
	_ = os.MkdirAll(uploadDir, os.ModePerm)

	saveUpload := func(file *multipart.FileHeader) *string {
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

	BuktiPath := saveUpload(BuktiFile)

	var data models.DataEntry
	if err := config.DB.First(&data, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	data.Tipe = req.Tipe
	data.JenisPengadaan = req.JenisPengadaan
	data.MetodePengadaan = req.MetodePengadaan
	data.KodePaket = req.KodePaket
	data.KodeRup = req.KodeRup
	data.TahunAnggaran = req.TahunAnggaran
	data.TanggalMasuk = req.TanggalMasuk
	data.SatuanKerja = req.SatuanKerja
	data.NamaPaket = req.NamaPaket
	data.SumberDana = req.SumberDana

	data.StatusPaket = req.StatusPaket
	data.StatusPengiriman = req.StatusPengiriman

	data.NilaiPagu = req.NilaiPagu
	data.NilaiHps = req.NilaiHps

	data.NomorKontrak = req.NomorKontrak
	data.NilaiKontrak = req.NilaiKontrak
	data.TanggalKontrak = req.TanggalKontrak
	data.NamaPpk = req.NamaPpk
	data.JabatanPpk = req.JabatanPpk
	data.Pemasukan = req.Pemasukan
	data.Pendaftar = req.Pendaftar

	data.NamaPimpinanPerusahaan = req.NamaPimpinanPerusahaan
	data.JabatanPimpinan = req.JabatanPimpinan

	data.Pemenang = req.Pemenang
	data.NilaiPenawaran = req.NilaiPenawaran
	data.NilaiNegosiasi = req.NilaiNegosiasi
	data.NomorTelp = req.NomorTelp
	data.Email = req.Email
	data.Npwp = req.Npwp
	data.AlamatPemenang = req.AlamatPemenang
	data.LokasiPekerjaan = req.LokasiPekerjaan

	if req.Catatan != nil && *req.Catatan != "" {
		data.Catatan = req.Catatan
	}

	if req.SelectedPpkId != nil {
		data.SelectedPpkId = req.SelectedPpkId
	}
	if BuktiPath != nil {
		data.BuktiFile = BuktiPath
	}

	if err := config.DB.Save(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "mengubah data gagal!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "mengubah data berhasil",
		"data":    data,
	})
}

func DeleteDataEntry(c *gin.Context) {
	id := c.Param("id")

	var data models.DataEntry
	config.DB.First(&data, id)

	var rab []models.RabHeader
	err := config.DB.Where("data_entry_id = ?", id).Find(&rab).Error
	if len(rab) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Paket Pekerjaan Konstruksi sudah di-assign ke user PPK",
		})
		return
	}

	err = config.DB.Delete(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Mengahapus dara gagal!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Menghapus data berhasil",
		"data":    data,
	})
}
