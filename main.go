package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	mainurlbeg    = os.Getenv("URLHEAD")
	mainurlend    = os.Getenv("URLTAIL")
	filepath      = "./damp/"
	startitter, _ = strconv.Atoi(os.Getenv("HEADITTER"))
	enditter, _   = strconv.Atoi(os.Getenv("TAILITTER"))
	fileformat    = os.Getenv("FILEFORMAT")
)

func main() {

	for i := startitter; i < enditter; i++ {
		url := mainurlbeg + strconv.Itoa(i) + mainurlend
		DownloadFile(i, url)
		fmt.Println("Current itter is: ", i)
	}
	fmt.Println("Finished...")
}

func DownloadFile(i int, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get Error", err)
		return err
	}
	defer resp.Body.Close()
	if strings.Contains(resp.Status, "404") {
		return err
	} else {
		// Create the file
		s := filepath + strconv.Itoa(i) + fileformat
		out, err := os.Create(s)
		if err != nil {
			return err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return err
	}

}
