package embed

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed test.txt
var content embed.FS

func Execute() {

	// reads data from embedded file
	data, err := fs.ReadFile(content, "test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Content: ", string(data))

	destinationDir := filepath.Join(os.Getenv("HOME"), ".config", "gamon")
	destinationPath := filepath.Join(destinationDir, "test.txt")

	// creates necessary directories
	err = os.MkdirAll(destinationDir, 0755)
	if err != nil {
		fmt.Println(err)
	}

	// copies data from embedded file to destination path
	err = os.WriteFile(destinationPath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file copied to: ", destinationPath)

}
