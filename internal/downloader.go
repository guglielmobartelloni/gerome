package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Download struct {
	URL        string
	FilePath   string
	FileSize   int64
	Downloaded int64
}

func DownloadWithResume(d *Download) error {
	// Check if file already exists
	fileInfo, err := os.Stat(d.FilePath)

	if err == nil {
		d.Downloaded = fileInfo.Size()
	} else if !os.IsNotExist(err) {
		return err
	}

	// Open the file for appending (to resume downloads)
	file, err := os.OpenFile(d.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create HTTP client
	client := &http.Client{}

	// Set Range header for resuming downloads
	req, err := http.NewRequest(http.MethodGet, d.URL, nil)
	req.Header.Add("Referer","https://www.erome.com/")
	if err != nil {
		return err
	}

	if d.Downloaded > 0 {
		fmt.Println(fmt.Sprintf("Requesting from range %d", d.Downloaded))
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", d.Downloaded))
	}

	// Send request and handle response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Write downloaded data to file
	writer := io.MultiWriter(file, io.Discard) // Discard extra data if file size changed on server
	n, err := io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}

	d.Downloaded += n

	// Check if download is complete
	if d.FileSize > 0 && d.Downloaded == d.FileSize {
		fmt.Println("Download completed!")
	}

	return nil
}
