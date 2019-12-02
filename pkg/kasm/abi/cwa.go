package abi

// #include <stdlib.h>
//
// extern int32_t io_get_stdout(void *context);
// extern int32_t resource_write(void *context, int32_t id, int32_t ptr, int32_t len);
import "C"
import (
  "unsafe"
  
  wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

//export io_get_stdout
func io_get_stdout(context unsafe.Pointer) int32 {
	return 1
}

//export resource_write
func resource_write(context unsafe.Pointer, id, ptr, len int32) int32 {
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
