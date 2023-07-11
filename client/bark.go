package client

type SunoaiBarkRequest struct {
	Model string          `json:"model"`
	Param SunoaiBarkParam `json:"data"`
}

type SunoaiBarkParam struct {
	Prompt              string  `json:"prompt"`
	Speaker             string  `json:"speaker"`
	SampleRate          int     `json:"sample_rate"`
	TextTemperature     float64 `json:"text_temp"`
	WaveformTemperature float64 `json:"waveform_temp"`
}
