package controllers

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"

// 	// "log"
// 	// "time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/optimus/backend/config"
// 	"github.com/optimus/backend/models"
// )

// func GetKatalogV5(c *gin.Context) {
// 	tahun := c.Query("tahun")

// 	var cache models.KatalogV5Cache

// 	bytes, err := os.ReadFile("cache/katalogv5s.json")
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

// 	bytes, err := os.ReadFile("cache/katalogv6s.json")
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

// 	bytes, err := os.ReadFile("cache/tenders.json")
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

// 	bytes, err := os.ReadFile("cache/tenderselesais.json")
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
// 	var data models.Url
// 	err := config.DB.First(&data).Error
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	url := fmt.Sprintf("https://%s/api2/token", data.Url)

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

// func GetDecodedAccessToken(c *gin.Context) {
// 	token, err := getAccessToken()
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"access_token": token,
// 	})
// }

// func fetchAndCacheLPSE(url, filePath string) error {
// 	token, err := getAccessToken()
// 	if err != nil {
// 		return err
// 	}

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

// 	return os.WriteFile(filePath, body, 0644)
// }

// func FetchAndCacheKatalogV5() error {
// 	return fetchAndCacheLPSE(
// 		"https://malutprov.lpse.info/api2/v1/katalogv5?tahun=2025",
// 		"cache/katalogv5s.json",
// 	)
// }

// func FetchAndCacheKatalogV6() error {
// 	return fetchAndCacheLPSE(
// 		"https://malutprov.lpse.info/api2/v1/katalogv6?tahun=2025",
// 		"cache/katalogv6s.json",
// 	)
// }

// func FetchAndCacheTender() error {
// 	return fetchAndCacheLPSE(
// 		"https://malutprov.lpse.info/api2/v1/tender?tahun=2025",
// 		"cache/tenders.json",
// 	)
// }

// func FetchAndCacheTenderSelesai() error {
// 	return fetchAndCacheLPSE(
// 		"https://malutprov.lpse.info/api2/v1/tenderselesai?tahun=2025",
// 		"cache/tenderselesais.json",
// 	)
// }

// func FetchAndCacheNonTenderSelesai() error {
// 	return fetchAndCacheLPSE(
// 		"https://malutprov.lpse.info/api2/v1/nontenderselesai?tahun=2025",
// 		"cache/nontenderselesais.json",
// 	)
// }

func StartLPSECron() {
	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c := cron.New(
	// 	cron.WithLocation(loc),
	// )

	// // 14:30 WIB
	// _, err = c.AddFunc("10 15 * * *", func() {
	// 	log.Println("Cron: Fetch Semua Data Selesai")

	// 	if err := FetchAndCacheKatalogV5(); err != nil {
	// 		log.Println("KatalogV5 error:", err)
	// 	}

	// 	if err := FetchAndCacheKatalogV6(); err != nil {
	// 		log.Println("KatalogV6 error:", err)
	// 	}

	// 	if err := FetchAndCacheTender(); err != nil {
	// 		log.Println("Tender error:", err)
	// 	}

	// 	if err := FetchAndCacheTenderSelesai(); err != nil {
	// 		log.Println("Tender Selesai error:", err)
	// 	}

	// 	if err := FetchAndCacheNonTenderSelesai(); err != nil {
	// 		log.Println("Non-Tender error:", err)
	// 	}
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c.Start()
}

// func StartNonTenderSelesaiCron() {
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
// 		if err := FetchAndCacheNonTenderSelesai(); err != nil {
// 			log.Println("Cron error:", err)
// 		}
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.Start()
// }

// func FetchAndCacheNonTenderSelesai() error {
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
// 	if err := FetchAndCacheNonTenderSelesai(); err != nil {
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

// func FetchAndCacheTenderSelesai() error {
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

// 	err = os.WriteFile("cache/tenderselesais.json", body, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SyncNonTenderSelesai(c *gin.Context) {
// 	if err := FetchAndCacheTenderSelesai(); err != nil {
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
