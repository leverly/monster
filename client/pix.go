package client

// image editing
type Pix2PixRequest struct {
	Model string       `json:"model"`
	Param Pix2PixParam `json:"data"`
}

const (
	IMG_FILE_TYPE_JPG  = "jpg"
	IMG_FILE_TYPE_JPEG = "jpeg"
	IMG_FILE_TYPE_PNG  = "png"
)

type Pix2PixParam struct {
	Prompt             string  `json:"prompt"`
	Negprompt          string  `json:"negprompt"`
	Steps              int     `json:"steps"`
	FileURL            string  `json:"init_image_url"`
	GuidanceScale      float64 `json:"guidance_scale"`
	ImageGuidanceScale float64 `json:"image_guidance_scale"`
	Seed               int     `json:"seed"`
}
