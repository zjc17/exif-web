package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
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

func TestParser_GojaFunction(t *testing.T) {
	vm := goja.New()
	goCallback := func(result any) {
		fmt.Printf("go_callback: %+v\n", result)
	}
	vm.Set("go_callback", goCallback)
	result, _ := vm.RunString(`
		let result = 0;
		const func = async function() {
			return 1;
		}
		func().then(data => {go_callback(data); return data;});
	`)
	promise := result.Export().(*goja.Promise)
	fmt.Println(promise.Result().Export())
	fmt.Println(vm.Get("result").Export().(int64))
}

func TestParser_Parse(t *testing.T) {
	image, _ := fetcher.GetImagePartial("https://image-hosting.zhangjc.tech/img/20220417210600.png", nil)

	goCallback := func(result any) {
		fmt.Printf("go_callback: %+v\n", result)
	}

	vm := goja.New()
	_, _ = vm.RunString("global={};" + ExifrScript())
	vm.Set("go_callback", goCallback)

	_ = vm.Set("fileData", vm.NewArrayBuffer(image))
	_, err := vm.RunString(`
		let result = exifr.parse(fileData).then(go_callback);
	`)
	time.Sleep(100 * time.Millisecond)
	assert.Nil(t, err)
}

func TestParser_NewParser(t *testing.T) {
	image, _ := fetcher.GetImagePartial("https://image-hosting.zhangjc.tech/img/20220417210600.png", nil)
	fmt.Println(newParse(image))
}

func TestParser_Parse2(t *testing.T) {
	image, _ := fetcher.GetImagePartial("https://image-hosting.zhangjc.tech/img/20220417210600.png", nil)
	parser := NewParser()
	result, err := parser.parse(image)
	fmt.Println(result, err)
}

func newParse(image []byte) (result map[string]any, err error) {
	done := make(chan map[string]any, 1)

	goCallback := func(result map[string]any) {
		done <- result
	}

	func(image []byte) {
		vm := goja.New()
		_, _ = vm.RunString("global={};" + ExifrScript())
		_ = vm.Set("go_callback", goCallback)
		_ = vm.Set("fileData", vm.NewArrayBuffer(image))
		_, err := vm.RunString(`exifr.parse(fileData).then(go_callback)`)
		if err != nil {
			fmt.Println(err)
		}
	}(image)

	select {
	case result := <-done: // If the function finishes before the timeout
		return result, nil
	case <-time.After(100 * time.Millisecond): // If the timeout expires before the function finishes
		err = errors.New("timeout")
		return
	}
}

func TestCancellation(t *testing.T) {
	done := make(chan int, 0)

	go func() {
		defer func() {
			fmt.Println("job defer")
		}()

		fmt.Println("job start")

		time.Sleep(100 * time.Millisecond)
		done <- 0

		fmt.Println("job done")
	}()

	select {
	case <-done: // If the function finishes before the timeout
		fmt.Println("job finished")
	case <-time.After(50 * time.Millisecond): // If the timeout expires before the function finishes
		fmt.Println("job timeout")
		return
	}
}

func TestCancellation_NoBuffer(t *testing.T) {
	done := make(chan int)

	go func() {
		defer func() {
			fmt.Println("job defer")
		}()

		fmt.Println("job start")

		time.Sleep(100 * time.Millisecond)
		done <- 1

		fmt.Println("job done")
	}()

	fmt.Println("block start")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("block end")

	select {
	case value := <-done: // If the function finishes before the timeout
		fmt.Println("job finished", value)
	case <-time.After(50 * time.Millisecond): // If the timeout expires before the function finishes
		fmt.Println("job timeout")
		return
	}
}

func TestCancellation_WithBuffer(t *testing.T) {
	done := make(chan int)

	go func() {
		defer func() {
			fmt.Println("job defer")
		}()

		fmt.Println("job start")

		time.Sleep(100 * time.Millisecond)
		done <- 1

		fmt.Println("job done")
	}()

	fmt.Println("block start")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("block end")

	select {
	case value := <-done: // If the function finishes before the timeout
		fmt.Println("job finished", value)
	case <-time.After(50 * time.Millisecond): // If the timeout expires before the function finishes
		fmt.Println("job timeout")
		return
	}
}

func TestCancellation_WithContext(t *testing.T) {

	job := func(ctx context.Context, ch chan<- int) {
		defer func() {
			fmt.Println("job defer")
			close(ch)
		}()

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("job timeout and exit")
				return
			case ch <- i:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Millisecond)
	ch := make(chan int)

	go job(ctx, ch)

	for {
		select {
		case value, ok := <-ch:
			if !ok {
				fmt.Println("chan closed. (job timeout)")
				return
			} else {
				fmt.Println("job:", value)
			}
		}
	}
}
