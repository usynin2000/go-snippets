package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func ConvertMP4ToMP3( mp4Path string) (string, error) {
	ext := filepath.Ext(mp4Path)

	mp3Path := strings.TrimSuffix(mp4Path, ext) + ".mp3"

	cmd := exec.Command(
		"ffmpeg", "-i", mp4Path, "-vn", "-acodec",  "libmp3lame", "-q:a", "2", mp3Path,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ffmpeg error: %v, output: %s", err, string(output))
	}

	return mp3Path, nil

}

func main(){
	mp3File, err := ConvertMP4ToMP3("test_video.mp4")

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("MP3-файл сохранён как:", mp3File)
}