package ytlmetadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL       = "https://www.youtube.com"
	clientVersion = "2.20200716.00.00"
)

type Metadata struct {
	client *http.Client
	key    string
}

func New() *Metadata {
	return &Metadata{client: &http.Client{}}
}

func (m *Metadata) fetchMetadata(videoID string) (*metadataResponse, error) {
	endpoint := baseURL + "/youtubei/v1/updated_metadata"

	body := &metadataRequest{
		Context: context{
			Client: client{
				Hl:            "en",
				ClientName:    "WEB",
				ClientVersion: clientVersion,
			},
		},
		VideoID: videoID,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	q := url.Values{"key": {m.key}}
	req.URL.RawQuery = q.Encode()

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r metadataResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, err
}

func (m *Metadata) Fetch(videoID string) (interface{}, error) {
	if m.key == "" {
		if err := m.updateKey(); err != nil {
			return nil, err
		}
	}

	metadata, err := m.fetchMetadata(videoID)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", metadata)

	return nil, nil
}

func getBetween(s string, a string, b string) string {
	splitted := strings.Split(s, a)
	if len(splitted) == 0 {
		return ""
	}

	backward := splitted[len(splitted)-1]
	end := strings.Index(backward, b)
	if end == -1 {
		return ""
	}

	return backward[:end]
}

func (m *Metadata) updateKey() error {
	resp, err := m.client.Get(baseURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	m.key = getBetween(string(b), `"innertubeApiKey":"`, `"`)

	if m.key == "" {
		return fmt.Errorf("failed to update key")
	}

	return nil
}
