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

		newParse(image []byte) (result map[string]any, err error)
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
	result, err = p.newParse(data)
	if err != nil {
		return
	}
	return
}

func (p parser) newParse(image []byte) (result map[string]any, err error) {
	done := make(chan map[string]any, 1)

	goCallback := func(result map[string]any) {
		done <- result
	}

	_ = p.vm.Set("go_callback", goCallback)
	_ = p.vm.Set("fileData", p.vm.NewArrayBuffer(image))
	_, err = p.vm.RunString(`exifr.parse(fileData).then(go_callback)`)
	if err != nil {
		return
	}

	select {
	case result := <-done: // If the function finishes before the timeout
		return result, nil
	case <-time.After(100 * time.Millisecond): // If the timeout expires before the function finishes
		err = errors.New("timeout")
		return
	}
}

func (p parser) parse(data goja.ArrayBuffer) (result map[string]any, err error) {
	_ = p.vm.Set("fileData", data)
	value, err := p.vm.RunString(fmt.Sprintf("exifr.parse(fileData)"))
	if err != nil {
		return
	}
	promise := value.Export().(*goja.Promise)
	result, _ = func() (result map[string]any, err error) {
		for i := 0; i < 10; i++ {
			switch promise.State() {
			case goja.PromiseStatePending:
				fmt.Println("PromiseStatePending", promise.Result())
			case goja.PromiseStateFulfilled:
				result = promise.Result().Export().(map[string]any)
				return
			case goja.PromiseStateRejected:
				fmt.Println("PromiseStateRejected", promise.Result())
				err = errors.New("rejected")
				return
			}
			time.Sleep(10 * time.Microsecond)
		}
		return
	}()
	return
}
