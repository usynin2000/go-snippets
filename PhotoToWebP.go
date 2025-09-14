package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
    "github.com/nfnt/resize"
)

func ToWebPWithResize(inputPath, outputPath string, lossless bool, quality float32, maxWidth uint) error {
	file, err := os.Open(inputPath)

	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Если ширина больше maxWidth — уменьшаем пропорционально
	if maxWidth > 0 && uint(img.Bounds().Dx()) > maxWidth {
		img = resize.Resize(maxWidth, 0, img, resize.Lanczos3)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create ouput file: %w", err)
	}

	defer outFile.Close()

	err = webp.Encode(outFile, img, &webp.Options{
		Lossless: lossless,
		Quality:  quality,
	})

	if err != nil {
		return fmt.Errorf("failed to encode to WebP: %w", err)
	}

	fmt.Printf("✅ Converted and resized %s -> %s\n", inputPath, outputPath)
	return nil
}

func main() {
	input := "Вывески.png"
	name := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
	output := name + ".webp"

	// Уменьшаем ширину до 1920 пикселей, качество 85%

	err := ToWebPWithResize(input, output, false, 85, 1920)

	if err != nil {
		fmt.Println("❌ Error:", err)
	}
}
