package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/zjc17/exif-web/internal/define"
	"github.com/zjc17/exif-web/internal/util"
	"time"
)

type (
	Parser interface {
		Parse(data []byte) (parseResult *ParseResult, err error)
		ParseJsonString(data []byte) (jsonStr string, err error)
		ParseRaw(data []byte) (result map[string]any, err error)

		parse(image []byte) (result map[string]any, err error)
	}

	parser struct {
		vm *goja.Runtime
	}
)

func NewParser() Parser {
	vm := goja.New()
	_, _ = vm.RunString("global={};" + ExifrScript())
	return &parser{vm: vm}
}

func (p parser) Parse(data []byte) (parseResult *ParseResult, err error) {
	jsonStr, err := p.ParseJsonString(data)
	if err != nil {
		return
	}
	parseResult = new(ParseResult)
	err = parseResult.From(jsonStr)
	if err != nil {
		return
	}
	return
}

func (p parser) ParseJsonString(data []byte) (jsonStr string, err error) {
	result, err := p.ParseRaw(data)
	if err != nil {
		return
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	jsonStr = string(bytes)
	return
}

func (p parser) ParseRaw(data []byte) (result map[string]any, err error) {
	data = util.TrimByteArray(data, define.ParseDataLength)
	result, err = p.parse(data)
	if err != nil {
		return
	}
	return
}

func (p parser) parse(image []byte) (result map[string]any, err error) {
	// use buffered channel to prevent race condition
	done := make(chan map[string]any, 1)

	go func() {
		timestamp := time.Now()
		defer func() {
			fmt.Println("parse time:", time.Since(timestamp))
		}()
		_ = p.vm.Set("go_callback", func(result map[string]any) {
			done <- result
		})
		_ = p.vm.Set("fileData", p.vm.NewArrayBuffer(image))
		_, err = p.vm.RunString(`exifr.parse(fileData).then(go_callback)`)
		if err != nil {
			return
		}
	}()

	select {
	case result := <-done: // If the function finishes before the timeout
		return result, nil
	case <-time.After(100 * time.Millisecond): // If the timeout expires before the function finishes
		err = errors.New("timeout")
		return
	}
}
