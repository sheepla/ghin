//nolint
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/sheepla/ghin/gh"
	"github.com/sheepla/ghin/ui"
)

// TODO: more elegant code
func main() {
	param := gh.NewSearchParam(os.Args[1])

	repos, err := gh.Search(param)
	if err != nil {
		log.Fatalln(err)
	}

	idx, err := ui.SelectRepo(*repos)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s/%s\n", (*repos)[idx].Owner, (*repos)[idx].Name)

	releases, err := gh.GetReleases((*repos)[idx].Owner, (*repos)[idx].Name)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(releases)

	idx, err = ui.SelectTag((*releases))
	if err != nil {
		log.Fatalln(err)
	}

	assets := (*releases)[idx].Assets

	idx, err = ui.SelectAsset(assets)
	if err != nil {
		log.Fatalln(err)
	}

	url := assets[idx].DownloadURL
	fmt.Println(url)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	download(url, pwd)
}

func download(url, destDir string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	fname := path.Base(url)

	f, err := os.Create(path.Join(destDir, fname))
	if err != nil {
		return err
	}

	defer f.Close()

	io.Copy(f, res.Body)

	return nil
}
