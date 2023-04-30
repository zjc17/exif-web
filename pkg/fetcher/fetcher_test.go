package fetcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildRequestHeader(t *testing.T) {
	getHeader := func(key string) string {
		switch key {
		case "origin":
			return "https://www.google.com"
		}
		return ""
	}
	header := BuildRequestHeader(getHeader)
	assert.Equal(t, "https://www.google.com", header["referer"])
}

func TestGetImagePartial(t *testing.T) {

	url := "https://raw.githubusercontent.com/zjc17/exif-web/main/.github/preview-01.png"
	param := &GetImageParam{
		RequestHeader: BuildRequestHeader(func(key string) string {
			switch key {
			case "origin":
				return "https://www.google.com"
			}
			return ""
		}),
	}
	_, err := GetImagePartial(url, param)
	assert.Nil(t, err)
}
