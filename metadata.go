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

type MetadataClient struct {
	client   *http.Client
	key      string
	Language string
}

type Metadata struct {
	ViewCount      string
	ShortViewCount string
	IsLive         bool
	LikeCount      string
	DislikeCount   string
	Date           string
	Title          string
	Description    string
}

func New() *MetadataClient {
	return &MetadataClient{
		client:   &http.Client{},
		Language: "en",
	}
}

func (m *MetadataClient) fetchMetadata(videoID string) (*metadataResponse, error) {
	endpoint := baseURL + "/youtubei/v1/updated_metadata"

	body := &metadataRequest{
		Context: context{
			Client: client{
				Hl:            m.Language,
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

func (m *MetadataClient) Fetch(videoID string) (*Metadata, error) {
	if m.key == "" {
		if err := m.updateKey(); err != nil {
			return nil, err
		}
	}

	resp, err := m.fetchMetadata(videoID)
	if err != nil {
		return nil, err
	}

	meta := Metadata{}

	for _, action := range resp.Actions {
		viewCount := action.UpdateViewershipAction.ViewCount.VideoViewCountRenderer
		toggleButton := action.UpdateToggleButtonTextAction
		date := action.UpdateDateTextAction
		title := action.UpdateTitleAction
		description := action.UpdateDescriptionAction.Description

		if runs := viewCount.ViewCount.Runs; len(runs) != 0 {
			meta.ViewCount = runs[0].Text
			meta.ShortViewCount = viewCount.ExtraShortViewCount.SimpleText
			meta.IsLive = viewCount.IsLive
			continue
		}
		if toggleButton.ButtonID == "TOGGLE_BUTTON_ID_TYPE_LIKE" {
			meta.LikeCount = toggleButton.DefaultText.SimpleText
			continue
		}
		if toggleButton.ButtonID == "TOGGLE_BUTTON_ID_TYPE_DISLIKE" {
			meta.DislikeCount = toggleButton.DefaultText.SimpleText
			continue
		}
		if date.DateText.SimpleText != "" {
			meta.Date = date.DateText.SimpleText
			continue
		}
		if title.Title.SimpleText != "" {
			meta.Title = title.Title.SimpleText
			continue
		}
		if runs := description.Runs; len(runs) != 0 {
			for _, descriptionRun := range runs {
				meta.Description += descriptionRun.Text
			}
			continue
		}
	}

	return &meta, nil
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

func (m *MetadataClient) updateKey() error {
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
