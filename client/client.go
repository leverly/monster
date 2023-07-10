package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.monsterapi.ai/apis/add-task"
)

type MonsterClient struct {
	apikey string
	token  string
}

func NewMonsterClient(apikey, token string) *MonsterClient {
	return &MonsterClient{apikey: apikey, token: token}
}

func (m *MonsterClient) Text2Img(param Text2ImgData) (*TaskResponse, error) {
	request := Text2ImgRequest{
		Model: "txt2img",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) Img2Img(param Img2ImgData) (*TaskResponse, error) {
	request := Img2ImgRequest{
		Model: "Img2img",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) Pix2Pix(param Pix2PixData) (*TaskResponse, error) {
	request := Pix2PixRequest{
		Model: "pix2pix",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) Speech2Text(param WhisperData) (*TaskResponse, error) {
	request := WhisperRequest{
		Model: "whisper",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) Text2Speech(param SunoaiBarkData) (*TaskResponse, error) {
	request := SunoaiBarkRequest{
		Model: "sunoai-bark",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) TextGeneration(param TextGenerationData) (*TaskResponse, error) {
	request := TextGenerationRequest{
		Model: "falcon-7b-instruct",
		Data:  param,
	}
	return m.postRequest(request)
}

func (m *MonsterClient) postRequest(param any) (*TaskResponse, error) {
	payload, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Add("x-api-key", m.apikey)
	req.Header.Add("Authorization", m.token)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp TaskResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
