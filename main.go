package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func generateRandomSquaresAvatar(filename string) error {

	width := 60
	height := 60
	gridSize := 10
	fillChance := 0.2

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	randomColor := color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}

	pixelAmount := 0
	for pixelAmount < 3 {
		pixelAmount = 0
		for x := 0; x < width/2; x += gridSize {
			for y := 0; y < height; y += gridSize {
				if rand.Float64() < fillChance {
					pixelAmount++
					for i := 0; i < gridSize; i++ {
						for j := 0; j < gridSize; j++ {
							img.Set(x+i, y+j, randomColor)
							img.Set(width-x-i-1, y+j, randomColor)
						}
					}
				}
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func openImage(filename string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filename)
	return cmd.Start()
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)

	filename := "avatar.png"
	if err := generateRandomSquaresAvatar(filename); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Done")
	if err := openImage(filename); err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
