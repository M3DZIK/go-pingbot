package api

type URLType struct {
	URL string `json:"url" binding:"required"`
}
