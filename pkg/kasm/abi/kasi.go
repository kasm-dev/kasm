package abi

// #include <stdlib.h>
//
// extern int32_t _host_read(void *ctxp, int32_t fd, int32_t len, int32_t ptr);
// extern int32_t _host_write(void *ctxp, int32_t fd, int32_t len, int32_t ptr);
// extern int32_t _http_route_add(void *ctxp, int32_t id, int32_t ptr, int32_t len);
import "C"
import (
	"fmt"
  "unsafe"
  
  wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

//export _host_read
func _host_read(ctxp unsafe.Pointer, fd, ptr, len int32) int32 {
	fmt.Println("_host_read")
	return 1
}

//export _host_write
func _host_write(ctxp unsafe.Pointer, fd, ptr, len int32) int32 {
	ctx := wasm.IntoInstanceContext(ctxp)
	fmt.Println("_host_write", fd, len, ptr)
	in := ctx.Memory().Data()[ptr:]
	fmt.Println(string(in[:len]), in[:len])
	return 1
}

//export _http_route_add
func _http_route_add(ctxp unsafe.Pointer, id, ptr, len int32) int32 {
	ctx := wasm.IntoInstanceContext(ctxp)
	in := ctx.Memory().Data()[ptr:]
	fmt.Println(string(in[:len]), in[:len])
	return 0
}

func AppendKasiImports(imports *wasm.Imports) (*wasm.Imports, error) {
	var err error
	importer := func(name string, impl interface{}, cptr unsafe.Pointer) {
		if err != nil {
			return
		}
		_, err = imports.Append(name, impl, cptr)
	}

	importer("_host_read", _host_read, C._host_read)
	importer("_host_write", _host_write, C._host_write)
	importer("_http_route_add", _http_route_add, C._http_route_add)
	return imports, err
}
