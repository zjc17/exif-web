package parser

import (
	"encoding/json"
	"fmt"
	"github.com/zjc17/exif-web/pkg/fetcher"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	image, _ := fetcher.GetImagePartial("https://image-hosting.zhangjc.tech/img/20220417210600.png", nil)
	parser := NewParser()
	timestamp := time.Now()
	defer func() {
		fmt.Println("time cost:", time.Since(timestamp))
	}()
	result, _ := parser.Parse(image)
	fmt.Printf("fulfilled: %+v\n", result)
}

func TestParser_ParseRaw(t *testing.T) {
	image, _ := fetcher.GetImagePartial("https://image-hosting.zhangjc.tech/img/20220417210600.png", nil)
	parser := NewParser()
	timestamp := time.Now()
	defer func() {
		fmt.Println("time cost:", time.Since(timestamp))
	}()
	result, _ := parser.ParseRaw(image)
	data, _ := json.Marshal(result)
	fmt.Printf("fulfilled: %+s\n", string(data))
}
