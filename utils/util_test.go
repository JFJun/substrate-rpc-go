package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func Test_GetFilePath(t *testing.T){
	root:=os.Getenv("GOPATH")
	fmt.Println(root)
	filepath.Walk(fmt.Sprintf("%s/pkg/mod/github.com/itering/scale.go@v1.1.39/source/",root), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
}
