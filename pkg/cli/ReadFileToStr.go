//=================ReadFileToStr=======================
package cli

import (
	"fmt"
	"io/fs"
	"os"
)

func ReadFileToStr(dir, filename string) (string, error) {

	// Open the file system (os.DirFS is a common implementation for the local file system)
	//dir = "."
	fsys := os.DirFS(dir)

	// Read the entire content of "data.txt" using io/fs.ReadFile
	data, err := fs.ReadFile(fsys, filename)
	if err != nil {
		//fmt.Println("Error reading file:", err)
		fmt.Println("Error : standard.txt does not exist")
		return "", err
	}

	return string(data), nil
}

//==================================================
