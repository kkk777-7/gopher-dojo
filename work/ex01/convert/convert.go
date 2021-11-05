package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

type Options struct {
	Quality int
}

func Convert() {

	word := os.Args[1]
	fmt.Println(word)
	file, err := os.Open(word)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println(*file)

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	dst := os.Args[2]
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	jpeg.Encode(out, img, nil)

}
