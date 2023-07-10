package client

type Text2ImgRequest struct {
	Model string       `json:"model"`
	Data  Text2ImgData `json:"data"`
}

type Text2ImgData struct {
	Prompt        string  `json:"prompt"`
	Negprompt     string  `json:"negprompt"`
	Samples       int     `json:"samples"`
	Steps         int     `json:"steps"`
	AspectRatio   string  `json:"aspect_ratio"`
	GuidanceScale float64 `json:"guidance_scale"`
	Seed          int     `json:"seed"`
}

type Img2ImgRequest struct {
	Model string      `json:"model"`
	Data  Img2ImgData `json:"data"`
}

type Img2ImgData struct {
	Prompt        string  `json:"prompt"`
	Negprompt     string  `json:"negprompt"`
	Steps         int     `json:"steps"`
	FileURL       string  `json:"init_image_url"`
	Strength      string  `json:"strength"`
	GuidanceScale float64 `json:"guidance_scale"`
	Seed          int     `json:"seed"`
}
