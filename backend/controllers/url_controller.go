package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/optimus/backend/config"
	"github.com/optimus/backend/dtos"
	"github.com/optimus/backend/models"
	"gorm.io/gorm"
)

func GetUrl(c *gin.Context) {
	query := config.DB
	var data []models.Url

	if err := query.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Mengambil data gagal",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    data,
	})
}

func CreateUrl(c *gin.Context) {
	query := config.DB
	var req dtos.CreateUrlRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Membuat data gagal",
			"error":   err.Error(),
		})
		return
	}

	data := models.Url{
		Url:   req.Url,
		Tahun: req.Tahun,
	}

	if err := query.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Membuat data gagal",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Membuat data berhasil",
		"data":    data,
	})
}

func UpdateUrl(c *gin.Context) {
	query := config.DB
	var req dtos.CreateUrlRequest
	var data models.Url

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update data gagal",
			"error":   err.Error(),
		})
		return
	}

	err := query.First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		data = models.Url{
			Url:   req.Url,
			Tahun: req.Tahun,
		}
		query.Create(&data)
	} else {
		query.Model(&data).Updates(models.Url{
			Url:   req.Url,
			Tahun: req.Tahun,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update URL berhasil",
		"data":    data,
	})
}

func getAccessToken() (string, error) {
	var data models.Url
	err := config.DB.First(&data).Error
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s/api2/token", data.Url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("admin", "Sdc2_30Lsk6_df4.,K")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token request failed, status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(result.AccessToken)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func GetDecodedAccessToken(c *gin.Context) {
	token, err := getAccessToken()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": token,
	})
}

func GetBaseUrl() (string, error) {
	var data models.Url
	err := config.DB.First(&data).Error
	if err != nil {
		return "", err
	}

	return data.Url, nil
}

func CallLPSE(endpoint string, query string) ([]byte, error) {
	baseUrl, err := GetBaseUrl()
	if err != nil {
		return nil, err
	}

	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	fullUrl := fmt.Sprintf("https://%s/api2/%s%s", baseUrl, endpoint, query)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func ProxyLPSE(c *gin.Context) {

	endpoint := c.Param("endpoint") 
	endpoint = strings.TrimPrefix(endpoint, "/")
	query := c.Request.URL.RawQuery

	var queryString string
	if query != "" {
		queryString = "?" + query
	}

	body, err := CallLPSE(endpoint, queryString)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(200, "application/json", body)
}
