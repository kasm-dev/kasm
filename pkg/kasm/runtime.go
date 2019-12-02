package kasm

import (
	"fmt"
	// "unsafe"

	"github.com/arccoza/kasm/pkg/kasm/abi"
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
	imports, _ := abi.AppendCWAImports(wasm.NewImports())
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
