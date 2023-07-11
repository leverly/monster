package client

// text2img
type Text2ImgRequest struct {
	Model string        `json:"model"`
	Param Text2ImgParam `json:"data"`
}

type Text2ImgParam struct {
	Prompt        string  `json:"prompt"`
	Negprompt     string  `json:"negprompt"`
	Samples       int     `json:"samples"`
	Steps         int     `json:"steps"`
	AspectRatio   string  `json:"aspect_ratio"`
	GuidanceScale float64 `json:"guidance_scale"`
	Seed          int     `json:"seed"`
}

// img2img
type Img2ImgRequest struct {
	Model string       `json:"model"`
	Param Img2ImgParam `json:"data"`
}

type Img2ImgParam struct {
	Prompt        string  `json:"prompt"`
	Negprompt     string  `json:"negprompt"`
	Steps         int     `json:"steps"`
	FileURL       string  `json:"init_image_url"`
	Strength      float64 `json:"strength"`
	GuidanceScale float64 `json:"guidance_scale"`
	Seed          int     `json:"seed"`
}
