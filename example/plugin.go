package main // it is required to build package as main
import (
	"fmt"
	"github.com/Pwalne/go-plugin-test/shared"
)

// These cannot be pointers.
type Example struct {}
func (e Example) Init() shared.PluginError {
	fmt.Printf("Plugin is calling Init():pluginError\n")
	return nil
}

var Plugin Example