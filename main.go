package main

import (
	"errors"
	"log"
	"monster/client"
	"time"
)

var text, imgUrl, voiceUrl string

func main() {
	var apiKey, token string
	proxy := client.NewMonsterClient(apiKey, token)
	Text2Speech(proxy)
}

// output is url file
func Text2Img(monster *client.MonsterClient) {
	resp, err := monster.Text2Img(client.Text2ImgParam{
		Prompt:        "rainy, storm, grass, horse, rainbow, flying",
		Negprompt:     "",
		GuidanceScale: 7.5,
		Steps:         30,
		AspectRatio:   "landscape",
		Seed:          100,
		Samples:       1,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("Text2Img failed:", err)
	}

	imgUrl, _ = GetFileResult(monster, resp.ProcessID)
}

func Img2Img(monster *client.MonsterClient) {
	resp, err := monster.Img2Img(client.Img2ImgParam{
		Prompt:        "rainy, storm, grass, horse, flying",
		Negprompt:     "rainbow",
		GuidanceScale: 7.5,
		Steps:         30,
		Seed:          100,
		FileURL:       imgUrl,
		Strength:      0.75,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("Img2Img failed:", err)
	}

	imgUrl, _ = GetFileResult(monster, resp.ProcessID)
}

func Pix2Pix(monster *client.MonsterClient) {
	resp, err := monster.Pix2Pix(client.Pix2PixParam{
		Prompt:             "rainy, storm, grass, flying kites",
		Negprompt:          "rainbow, horse",
		GuidanceScale:      8.5,
		ImageGuidanceScale: 3.0,
		Steps:              30,
		Seed:               100,
		FileURL:            imgUrl,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("Pix2Pix failed:", err)
	}

	imgUrl, _ = GetFileResult(monster, resp.ProcessID)
}

func Text2Speech(monster *client.MonsterClient) {
	resp, err := monster.Text2Speech(client.SunoaiBarkParam{
		Prompt:              "hello world, i am a robot, haha",
		Speaker:             "en_speaker_1",
		SampleRate:          25000,
		TextTemperature:     0.3,
		WaveformTemperature: 0.5,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("Text2Speech failed:", err)
	}

	voiceUrl, _ = GetFileResult(monster, resp.ProcessID)
}

// output is text
func Speech2Text(monster *client.MonsterClient) {
	resp, err := monster.Speech2Text(client.WhisperParam{
		Prompt:              "robot, haha",
		Language:            "en",
		FileURL:             voiceUrl,
		TranscriptionFormat: client.TEXT_TYPE_FORMAT_TEXT,
		RemoveSilence:       true,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("Speech2Text failed:", err)
	}

	text, _ = GetTextResult(monster, resp.ProcessID)
}

func TextGeneration(monster *client.MonsterClient) {
	resp, err := monster.TextGeneration(client.TextGenerationParam{
		Prompt: "hello world, tell a story about the robot",
		TopP:   0.5,
		TopK:   10,
	})

	if err != nil || resp.ProcessID == "" {
		log.Println("TextGeneration failed:", err)
	}

	text, _ = GetTextResult(monster, resp.ProcessID)
}

func GetFileResult(monster *client.MonsterClient, processId string) (string, error) {
	for true {
		result, err := monster.TaskStatus(processId)
		if err != nil {
			return "", err
		}
		switch result.Status() {
		case client.TASK_STATUS_COMPLETED:
			url := result.GetOutput()[0]
			log.Println("task completed:", url)
			return url, nil
		case client.TASK_STATUS_FAILED:
			log.Println("check task failed:", result.GetErrMessage())
			return "", errors.New(result.GetErrMessage())
		}
		time.Sleep(time.Second * 2)
	}
	return "", errors.New("fatal error")
}

func GetTextResult(monster *client.MonsterClient, processId string) (string, error) {
	for true {
		result, err := monster.TaskStatus(processId)
		if err != nil {
			return "", err
		}
		switch result.Status() {
		case client.TASK_STATUS_COMPLETED:
			text := result.GetText()
			log.Println("task completed:", text)
			return text, nil
		case client.TASK_STATUS_FAILED:
			log.Println("check task failed:", result.GetErrMessage())
			return "", errors.New(result.GetErrMessage())
		}
		time.Sleep(time.Second * 2)
	}
	return "", errors.New("fatal error")
}
