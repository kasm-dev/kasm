package kasm

// #include <stdlib.h>
//
// extern int32_t io_get_stdout(void *context);
// extern int32_t resource_write(void *context, int32_t id, int32_t ptr, int32_t len);
import "C"
import (
	"fmt"
	"unsafe"

	// cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	// "github.com/perlin-network/life"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

type Container struct {
	inst wasm.Instance
	Imports *wasm.Imports
}

func (c *Container) AddImport(bytes []byte) error {
	return nil
}

func (c *Container) Load(bytes []byte) error {
	var err error
	imports, _ := wasm.NewImports().Append("io_get_stdout", io_get_stdout, C.io_get_stdout)
	imports, _ = imports.Append("resource_write", resource_write, C.resource_write)
	c.inst, err = wasm.NewInstanceWithImports(bytes, imports)
	// c.inst, err = wasm.NewInstance(bytes)
	fmt.Println(err, c.inst.Exports)
	// sum := c.inst.Exports["sum"]
	// ans := sum(1, 2)
	// fmt.Println(c.inst.Memory.Data())
	return err
}

func (c *Container) LoadPath(path string) error {
	if bytes, err := wasm.ReadBytes(path); err != nil {
		return err
	} else if err = c.Load(bytes); err != nil {
		return err
	} else {
		return nil
	}
}

//export io_get_stdout
func io_get_stdout(context unsafe.Pointer) int32 {
	return 1
}

//export resource_write
func resource_write(context unsafe.Pointer, id, ptr, len int32) int32 {
	return 0
}
