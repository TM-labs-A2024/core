package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/google/uuid"
)

type DropboxClient struct {
	bearer string
}

type payload struct {
	Path string `json:"path"`
}

func New(apiKey string) (*DropboxClient, error) {
	client := new(DropboxClient)
	client.bearer = fmt.Sprintf("Bearer %s", apiKey)
	return client, nil
}

type uploadArgs struct {
	Autorename     bool   `json:"autorename"`
	Mode           string `json:"mode"`
	Mute           bool   `json:"mute"`
	Path           string `json:"path"`
	StrictConflict bool   `json:"strict_conflict"`
}

type tempLinkResp struct {
	Link string `json:"link"`
}

func (c *DropboxClient) UploadFile(patientId uuid.UUID, body io.Reader) (string, error) {
	fileId := uuid.New()
	path := filepath.Join("/", patientId.String(), fileId.String())
	payload := uploadArgs{
		Autorename:     false,
		Mode:           "add",
		Mute:           false,
		Path:           path,
		StrictConflict: false,
	}

	data, err := json.Marshal(&payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", c.bearer)
	req.Header.Set("Dropbox-Api-Arg", string(data))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return path, nil
}

func (c *DropboxClient) GenerateURL(path string) (*url.URL, error) {
	data := payload{path}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/get_temporary_link", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.bearer)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tempLinkData := tempLinkResp{}
	if err := json.NewDecoder(resp.Body).Decode(&tempLinkData); err != nil {
		return nil, err
	}

	fileUrl, err := url.Parse(tempLinkData.Link)
	if err != nil {
		return nil, err
	}

	return fileUrl, nil
}

func (c *DropboxClient) DeleteFile(path string) error {
	data := payload{path}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/delete_v2", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.bearer)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
