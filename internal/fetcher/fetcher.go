package fetcher

import (
	"fmt"
	"github.com/zjc17/exif-web/internal/define"
	"io"
	"net/http"
)

var (
	client = &http.Client{}
)

func GetImagePartial(url string) (data []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=0-%d", define.ParseDataLength))
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
