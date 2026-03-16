package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/optimus/backend/config"
	"github.com/optimus/backend/models"
)

func GetAllNonTenderKontractIsb(c *gin.Context) {
	var data []models.NonTenderKontractIsb
	config.DB.Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    data,
	})
}

// func GetByIdNonTenderKontractIsb(c *gin.Context) {
// 	id := c.Param("id")

// 	var data models.NonTenderKontractIsb
// 	if err := config.DB.First(&data, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Data tidak ditemukan",
// 		})
// 		return
// 	}

// 	err := config.DB.Delete(&data).Error
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Menghapus data gagal!",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Menghapus data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllNonTenderIsb(c *gin.Context) {
// 	var data []models.NonTenderIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllNonTenderSelesaiIsb(c *gin.Context) {
// 	year := c.Query("tahun")

// 	var data []models.NonTenderSelesaiIsb
// 	config.DB.Where("tahun_anggaran = ?", year).Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllNonTenderTahapIsb(c *gin.Context) {
// 	var data []models.NonTenderTahapIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllTenderIsb(c *gin.Context) {
// 	var data []models.TenderIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllTenderKontrakIsb(c *gin.Context) {
// 	var data []models.TenderKontrakIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllTenderSelesaiIsb(c *gin.Context) {
// 	var data []models.TenderSelesaiIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllTenderTahapIsb(c *gin.Context) {
// 	var data []models.TenderTahapIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// // func GetAllPaketPurchasing(c *gin.Context) {
// // 	var data []models.PaketPurchasing
// // 	config.DB.Find(&data)

// // 	c.JSON(http.StatusOK, gin.H{
// // 		"message": "Mengambil data berhasil",
// // 		"data":    data,
// // 	})
// // 	return
// // }

// func GetAllPaketPurchasing(c *gin.Context) {
// 	kd_paket := c.Query("kd_paket")
// 	year := c.Query("tahun")
// 	var data []models.PaketPurchasing

// 	if kd_paket != "" {
// 		config.DB.Where("tahun_anggaran = ?", year).Where("kd_paket = ?", kd_paket).Find(&data)
// 	} else {
// 		config.DB.Find(&data)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllPencatatanNonTenderIsb(c *gin.Context) {
// 	var data []models.PencatatanNonTenderIsb
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllRupPaketPenyedia(c *gin.Context) {
// 	var data []models.RupPaketPenyedia
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

// func GetAllRupPaketSwakelola(c *gin.Context) {
// 	var data []models.RupPaketSwakelola
// 	config.DB.Find(&data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Mengambil data berhasil",
// 		"data":    data,
// 	})
// }

func GetAllRupPenyediaTerumumkan(c *gin.Context) {
	var data []models.RupPenyediaTerumumkan
	config.DB.Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Mengambil data berhasil",
		"data":    data,
	})
}

// func GetKatalogV5(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.KatalogV5Cache

// 	bytes, err := os.ReadFile("cache/katalogv5.json")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal membaca cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	if err := json.Unmarshal(bytes, &cache); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal parse cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result := make([]models.KatalogV5, 0)

// 	if tahun != "" {
// 		for _, item := range cache.Data {
// 			if item.TahunAnggaran == tahun {
// 				result = append(result, item)
// 			}
// 		}
// 	} else {
// 		result = cache.Data
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 	})
// }

// func GetKatalogV6(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.KatalogV6Cache

// 	bytes, err := os.ReadFile("cache/katalogv6.json")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal membaca cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	if err := json.Unmarshal(bytes, &cache); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal parse cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result := make([]models.KatalogV6, 0)

// 	if tahun != "" {
// 		for _, item := range cache.Data {
// 			if strconv.Itoa(item.TahunAnggaran) == tahun {
// 				result = append(result, item)
// 			}
// 		}
// 	} else {
// 		result = cache.Data
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 	})
// }

// func GetNonTenderSelesai(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.NonTenderSelesaiCahce

// 	bytes, err := os.ReadFile("cache/nontenderselesais.json")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal membaca cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	if err := json.Unmarshal(bytes, &cache); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal parse cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result := make([]models.NonTenderSelesai, 0)

// 	if tahun != "" {
// 		for _, item := range cache.Data {
// 			if strconv.Itoa(item.TahunAnggaran) == tahun {
// 				result = append(result, item)
// 			}
// 		}
// 	} else {
// 		result = cache.Data
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 	})
// }

// func GetTender(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.TenderCache

// 	bytes, err := os.ReadFile("cache/tender.json")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal membaca cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	if err := json.Unmarshal(bytes, &cache); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal parse cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result := make([]models.Tender, 0)

// 	if tahun != "" {
// 		for _, item := range cache.Data {
// 			if strconv.Itoa(item.TahunAnggaran) == tahun {
// 				result = append(result, item)
// 			}
// 		}
// 	} else {
// 		result = cache.Data
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 	})
// }

// func GetTenderSelesai(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.TenderSelesaiCache

// 	bytes, err := os.ReadFile("cache/tenderselesai.json")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal membaca cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	if err := json.Unmarshal(bytes, &cache); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal parse cache katalog",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	result := make([]models.TenderSelesai, 0)

// 	if tahun != "" {
// 		for _, item := range cache.Data {
// 			if strconv.Itoa(item.TahunAnggaran) == tahun {
// 				result = append(result, item)
// 			}
// 		}
// 	} else {
// 		result = cache.Data
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": result,
// 	})
// }


// func getAccessToken() (string, error) {
// 	url := "https://malutprov.lpse.info/api2/token"

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	req.SetBasicAuth("admin", "Sdc2_30Lsk6_df4.,K")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("token request failed, status: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	var result struct {
// 		AccessToken string `json:"access_token"`
// 	}

// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return "", err
// 	}

// 	decoded, err := base64.StdEncoding.DecodeString(result.AccessToken)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(decoded), nil
// }

// func StartTenderSelesaiCron() {
// 	loc, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c := cron.New(
// 		cron.WithLocation(loc),
// 	)

// 	// Menit Jam Hari Bulan HariDalamMinggu
// 	// 24 14 * * *  => jam 14:24
// 	_, err = c.AddFunc("30 14 * * *", func() {
// 		log.Println("Running cron: Fetch Tender Selesai")
// 		if err := fetchAndCacheTenderSelesai(); err != nil {
// 			log.Println("Cron error:", err)
// 		}
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.Start()
// }


// func fetchAndCacheTenderSelesai() error {
// 	token, err := getAccessToken()
// 	if err != nil {
// 		return err
// 	}

// 	url := "https://malutprov.lpse.info/api2/v1/nontenderselesai?tahun=2025"

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+token)
// 	req.Header.Set("Accept", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("data request failed, status: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	if err := os.MkdirAll("cache", 0755); err != nil {
// 		return err
// 	}

// 	err = os.WriteFile("cache/nontenderselesais.json", body, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SyncTenderSelesai(c *gin.Context) {
// 	if err := fetchAndCacheTenderSelesai(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Gagal sync data",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Cache berhasil diperbarui",
// 	})
// }
