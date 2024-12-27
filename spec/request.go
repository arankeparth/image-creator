package spec

type GenerateImageRequest struct {
	IsUpdate bool   `json:"isUpdate"`
	Code     string `json:"code"`
}
