package client

type WhisperRequest struct {
	Model string      `json:"model"`
	Data  WhisperData `json:"data"`
}

const (
	VOICE_FILE_TYPE_M4A  = "m4a"
	VOICE_FILE_TYPE_MP3  = "mp3"
	VOICE_FILE_TYPE_MP4  = "mp4"
	VOICE_FILE_TYPE_MPEG = "mpeg"
	VOICE_FILE_TYPE_MPGA = "mpga"
	VOICE_FILE_TYPE_WAV  = "wav"
	VOICE_FILE_TYPE_WEBM = "webm"
	VOICE_FILE_TYPE_OGG  = "ogg"
)

type WhisperData struct {
	FileURL             string `json:"file"`
	TranscriptionFormat string `json:"transcription_format"`
	Prompt              string `json:"prompt"`
	Language            string `json:"language"`
}
