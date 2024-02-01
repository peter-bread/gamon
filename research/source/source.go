package source

import (
	"embed"
	"fmt"
	"io/fs"
)

// this is pretty much exactly how `gam script` will work

//go:embed hello.zsh
var content embed.FS

func Execute() {

	// read script
	data, err := fs.ReadFile(content, "hello.zsh")

	// handle error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print script
	fmt.Println(string(data))

}
