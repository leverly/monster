package client

type SunoaiBarkRequest struct {
	Model string         `json:"model"`
	Data  SunoaiBarkData `json:"data"`
}

type SunoaiBarkData struct {
	Prompt       string  `json:"prompt"`
	Speaker      string  `json:"speaker"`
	SampleRate   int     `json:"sample_rate"`
	TextTemp     float64 `json:"text_temp"`
	WaveformTemp float64 `json:"waveform_temp"`
}
