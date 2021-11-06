package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func Convert(st string, dst string, pre string, after string) error {

	file, err := os.Open(st)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	switch after {
	case "jpg":
		err = jpeg.Encode(out, img, nil)
	case "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, nil)
	}
	if err != nil {
		return err
	}
	fmt.Printf("convert: %s\n", out.Name())

	return nil
}
