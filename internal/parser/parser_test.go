package parser

import (
	"fmt"
	"github.com/zjc17/exif-web/internal/fetcher"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	image, _ := fetcher.GetImagePartial("http://127.0.0.1:8080/DSCF3404.jpg")
	parser := NewParser()
	timestamp := time.Now()
	defer func() {
		fmt.Println("time cost:", time.Since(timestamp))
	}()
	result, _ := parser.Parse(image)
	fmt.Printf("fulfilled: %+v\n", result)
}
