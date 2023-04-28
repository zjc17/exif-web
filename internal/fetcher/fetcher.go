package fetcher

import (
	"fmt"
	"github.com/zjc17/exif-web/internal/define"
	"io"
	"net/http"
	"strings"
)

var (
	client = &http.Client{}
)

type (
	GetImageParam struct {
		RequestHeader map[string]string
	}
)

var (
	defaultParam = GetImageParam{}
)

func BuildRequestHeader(getHeader func(key string) string) map[string]string {
	return map[string]string{
		"referer": getHeader("origin"),
	}
}

func GetImagePartial(url string, param *GetImageParam) (data []byte, err error) {
	if param == nil {
		param = &defaultParam
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range param.RequestHeader {
		req.Header.Add(strings.ToLower(k), v)
	}
	req.Header.Add("range", fmt.Sprintf("bytes=0-%d", define.ParseDataLength))

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
