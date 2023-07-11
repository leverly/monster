package client

type WhisperRequest struct {
	Model string       `json:"model"`
	Param WhisperParam `json:"data"`
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

const (
	TEXT_TYPE_FORMAT_TEXT    = "text"
	TEXT_TYPE_FORMAT_SRT     = "srt"
	TEXT_TYPE_FORMAT_WORD    = "word"
	TEXT_TYPE_FORMAT_VERBOSE = "verbose"
)

type WhisperParam struct {
	Prompt              string `json:"prompt"`
	FileURL             string `json:"file"`
	RemoveSilence       bool   `json:"remove_silence"`
	TranscriptionFormat string `json:"transcription_format"`
	Language            string `json:"language"`
}
