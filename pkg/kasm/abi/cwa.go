package abi

// #include <stdlib.h>
//
// extern int32_t io_get_stdout(void *ctxp);
// extern int32_t resource_write(void *ctxp, int32_t id, int32_t ptr, int32_t len);
import "C"
import (
	"fmt"
  "unsafe"

  wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

//export io_get_stdout
func io_get_stdout(ctxp unsafe.Pointer) int32 {
	fmt.Println("io_get_stdout")
	return 1
}

//export resource_write
func resource_write(ctxp unsafe.Pointer, fd, ptr, len int32) int32 {
	ctx := wasm.IntoInstanceContext(ctxp)
	fmt.Println("resource_write", fd, ptr, len)
	if fd == 1 {
		out := ctx.Memory().Data()[ptr:]
		fmt.Print(string(out[:len]), out[:len])
	}
	return 0
}

func AppendCWAImports(imports *wasm.Imports) (*wasm.Imports, error) {
	if _, err := imports.Append("io_get_stdout", io_get_stdout, C.io_get_stdout); err != nil {
		return nil, err
	} else if _, err = imports.Append("resource_write", resource_write, C.resource_write); err != nil {
		return nil, err
	}
	return imports, nil
}
