package ytlmetadata

import (
	"os"
	"testing"
)

func TestMetadata(t *testing.T) {
	m := New()
	if _, err := m.Fetch(os.Getenv("VIDEO_ID")); err != nil {
		t.Fatal(err)
	}
}
