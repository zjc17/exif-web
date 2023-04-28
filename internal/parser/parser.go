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
		ParseRaw(data []byte) (result map[string]any, err error)
		Parse(data []byte) (parseResult *ParseResult, err error)
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
	result, err := p.ParseRaw(data)
	if err != nil {
		return
	}
	parseResult = &ParseResult{}
	bytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bytes, &parseResult); err != nil {
		return
	}
	return
}

func (p parser) ParseRaw(data []byte) (result map[string]any, err error) {
	data = util.TrimByteArray(data, define.ParseDataLength)
	result, err = p.parse(p.vm.NewArrayBuffer(data))
	if err != nil {
		return
	}
	return
}

func (p parser) parse(data goja.ArrayBuffer) (result map[string]any, err error) {
	_ = p.vm.Set("fileData", data)
	value, err := p.vm.RunString(fmt.Sprintf("exifr.parse(fileData)"))
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
				err = errors.New("rejected")
				return
			}
			time.Sleep(10)
		}
		return
	}()
	return
}
