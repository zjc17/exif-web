package main

import (
	"fmt"
	"github.com/zjc17/exif-web/internal/fetcher"
	"github.com/zjc17/exif-web/internal/parser"
	"time"
)

func main() {
	GojaExample()
}

func GojaExample() {
	image, _ := fetcher.GetImagePartial("http://127.0.0.1:8080/DSCF3404.jpg")
	p := parser.NewParser()
	timestamp := time.Now()
	defer func() {
		fmt.Println("time cost:", time.Since(timestamp))
	}()
	result, _ := p.Parse(image)
	fmt.Printf("fulfilled: %+v\n", result)
	return
}
