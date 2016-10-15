package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	DEFAULT_PORT = "9000"
)

var (
	port     string
	filename = "https://stream-alfa.dropcam.com:443/nexus_aac/b8fbe1918dd5470e913d5780445bc66b/playlist.m3u8"
	width    = 1920
	height   = 1080
)

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":9000", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		panic("could not get frame" + err.Error())
	}
	r := bytes.NewReader(buffer.Bytes())
	//img, _ := jpeg.Decode(r)
	file, err := os.Create("img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, r)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Fprintf(w, "we made it")
}
