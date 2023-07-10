package client

type Pix2PixRequest struct {
	Model string      `json:"model"`
	Data  Pix2PixData `json:"data"`
}

const (
	IMG_FILE_TYPE_JPG  = "jpg"
	IMG_FILE_TYPE_JPEG = "jpeg"
	IMG_FILE_TYPE_PNG  = "png"
)

type Pix2PixData struct {
	Prompt             string  `json:"prompt"`
	Negprompt          string  `json:"negprompt"`
	Steps              int     `json:"steps"`
	FileURL            string  `json:"init_image_url"`
	GuidanceScale      float64 `json:"guidance_scale"`
	ImageGuidanceScale float64 `json:"image_guidance_scale"`
	Seed               int     `json:"seed"`
}
