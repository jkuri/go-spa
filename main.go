package main

import (
	"log"
	"net/http"
	"os"
	"path"

	_ "go-spa/statik" // UI files

	"github.com/rakyll/statik/fs"
)

type index struct {
	assets http.FileSystem
}

func (i *index) Open(name string) (http.File, error) {
	ret, err := i.assets.Open(name)
	if !os.IsNotExist(err) || path.Ext(name) != "" {
		return ret, err
	}

	return i.assets.Open("/index.html")
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(&index{statikFS}))
	http.ListenAndServe(":8080", nil)
}
