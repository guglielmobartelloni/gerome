package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	outputDir string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func getLinks(url string) map[string]struct{} {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := regexp.MustCompile(`<source src="(.*)"`)
	matches := r.FindAllStringSubmatch(string(bodyText), -1)
	links := make(map[string]struct{})
	for _, v := range matches {
		links[v[1]] = struct{}{}
	}

	return links
}

func (a *App) DownloadVideos(url string) {
	message := ""
	if url == "" {
		alert("You must provide a url", a)
	}
	if a.outputDir == "" {
		alert("You must set an output dir", a)
	}

	links := getLinks(url)
	nLinks := len(links)
	if nLinks == 0 {
		alert("There are no videos on the provided url", a)
	}

	alert(fmt.Sprintf("Found %d videos!", nLinks), a)

	var wg sync.WaitGroup
	for link := range links {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()
			fmt.Println(fmt.Sprintf("Processing %s", link))

		}(link)
	}
	wg.Wait()
	message = "Finished downloading videos!"
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{Message: message})

}

func alert(message string, a *App) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{Message: message})
}

func (a *App) SetOutputDir() string {
	outputDir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		panic("Something went wrong with the file picker.")
	}
	a.outputDir = outputDir
	return outputDir
}
