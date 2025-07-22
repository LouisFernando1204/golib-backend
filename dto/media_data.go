package dto

type MediaData struct {
	Id   string `json:"id"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

type CreateMediaRequest struct {
	Path string
}
