package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	ADD_TASK_URL = "https://api.monsterapi.ai/apis/add-task"
	GET_TASK_URL = "https://api.monsterapi.ai/apis/task-status"
)

type MonsterClient struct {
	apikey string
	token  string
}

func NewMonsterClient(apikey, token string) *MonsterClient {
	return &MonsterClient{apikey: apikey, token: token}
}

func (m *MonsterClient) TaskStatus(processId string) (*TaskResultResponse, error) {
	request := TaskStatusRequest{
		ProcessID: processId,
	}
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := m.request(GET_TASK_URL, payload)
	if err != nil {
		return nil, err
	}

	var resp TaskResultResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *MonsterClient) Text2Img(param Text2ImgParam) (*AddTaskResponse, error) {
	request := Text2ImgRequest{
		Model: "txt2img",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) Img2Img(param Img2ImgParam) (*AddTaskResponse, error) {
	request := Img2ImgRequest{
		Model: "img2img",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) Pix2Pix(param Pix2PixParam) (*AddTaskResponse, error) {
	request := Pix2PixRequest{
		Model: "pix2pix",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) Speech2Text(param WhisperParam) (*AddTaskResponse, error) {
	request := WhisperRequest{
		Model: "whisper",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) Text2Speech(param SunoaiBarkParam) (*AddTaskResponse, error) {
	request := SunoaiBarkRequest{
		Model: "sunoai-bark",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) TextGeneration(param TextGenerationParam) (*AddTaskResponse, error) {
	request := TextGenerationRequest{
		Model: "falcon-7b-instruct",
		Param: param,
	}
	return m.addTask(request)
}

func (m *MonsterClient) addTask(param any) (*AddTaskResponse, error) {
	payload, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	body, err := m.request(ADD_TASK_URL, payload)
	if err != nil {
		return nil, err
	}
	var resp AddTaskResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *MonsterClient) request(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
