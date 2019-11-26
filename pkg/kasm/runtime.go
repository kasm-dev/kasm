package kasm

import (
	"fmt"

	// "github.com/perlin-network/life"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

type Container struct {
	inst wasm.Instance
}

func (c *Container) Load(bytes []byte) error {
	var err error
	c.inst, err = wasm.NewInstance(bytes)
	fmt.Println(err)
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
