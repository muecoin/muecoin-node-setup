package main

import (
	"archive/zip"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cavaliercoder/grab"
)

func main() {

	homepath := ""

	if runtime.GOOS == "linux" {
		homepath = "~/.monetaryunit"
	}

	fmt.Println(homepath)
	DownloadBinaries()
	untar()

}

func DownloadBinaries() {

	url := "https://github.com/muecoin/MUE/releases/download/v2.1.4/mon-2.1.4-x86_64-linux-gnu.tar.gz"

	client := grab.NewClient()
	req, _ := grab.NewRequest(".", url)

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("Downloading MonetaryUnit binaries (%.2f%%)\n",
				100*resp.Progress())

		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	r, err := zip.OpenReader(resp.Filename)
	if err != nil {
		//return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()
}

func untar() {
	exec.Command("/bin/sh", "untar.sh").Run()
}
