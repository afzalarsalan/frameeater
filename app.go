package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

var (
	port     string
	filename = "https://stream-alfa.dropcam.com:443/nexus_aac/b8fbe1918dd5470e913d5780445bc66b/playlist.m3u8"
	width    = 1920
	height   = 1080
)

type imageResponse struct {
	ImageURL string `json:"imageurl"`
}

func main() {
	http.HandleFunc("/", homepage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9001", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "img.jpg")
	//var buffer bytes.Buffer
	//cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		panic("could not get frame" + err.Error())
	}
	/*r := bytes.NewReader(buffer.Bytes())
	//img, _ := jpeg.Decode(r)
	file, err := os.Create("static/img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, r)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()*/
	resp := &imageResponse{ImageURL: "13.84.145.193:9000/static/img.jpg"}
	json.NewEncoder(w).Encode(resp)
}
