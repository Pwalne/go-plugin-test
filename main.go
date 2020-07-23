package main

import (
	"fmt"
	"github.com/Pwalne/plugin-test/shared"
	"path/filepath"
	"plugin"
)

func main() {
	files, err := filepath.Glob("plugins/*.so")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		fmt.Printf("Loading file %s\n", f)
		pl, err := plugin.Open(f)
		if err != nil {
			fmt.Printf("Error %s\n Reason: %s\n", f, err)
			continue
		}
		sym, err := pl.Lookup("Plugin")
		if err != nil {
			fmt.Printf("Could not find sym `Plugin`. Reason: %s\n", err)
			continue
		}
		if plgin, ok := sym.(shared.Plugin); ok {
			plgin.Init()
		} else {
			fmt.Println("sym not of type shared.Plugin")
		}
	}
	fmt.Println("Finished loading plugins.")
}
