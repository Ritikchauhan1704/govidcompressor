package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func compressVideo(inputPath, outputPath string) error {
	// ffmpeg command:
	// -vcodec libx264: use H.264 encoder
	// -crf 28: lower = better quality, higher = more compression
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vcodec", "libx264", "-crf", "28", outputPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	// change this to your video folder path
	videoDir := "./videos"
	outputDir := "./compressed"

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(videoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// skip folders
		if info.IsDir() {
			return nil
		}

		// process only video files
		ext := filepath.Ext(path)
		if ext == ".mp4" || ext == ".mov" || ext == ".mkv" {
			outputPath := filepath.Join(outputDir, info.Name())
			fmt.Println("Compressing:", path)
			err := compressVideo(path, outputPath)
			if err != nil {
				log.Printf("Error compressing %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
