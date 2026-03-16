package dtos

type CreateUrlRequest struct {
	Url      string  `form:"url"`
	Tahun    *string `form:"tahun"`
}