package ytlmetadata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "https://www.youtube.com"

type Metadata struct {
	client *http.Client
	key    string
}

func New() *Metadata {
	return &Metadata{client: &http.Client{}}
}

func (m *Metadata) Fetch(videoID string) (interface{}, error) {
	if m.key == "" {
		if err := m.updateKey(); err != nil {
			return nil, err
		}
	}

	// endpoint := baseURL + "/youtubei/v1/updated_metadata"
	println(m.key)
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
