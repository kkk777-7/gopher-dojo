package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageFormat struct {
	Path        string
	PreFormat   string
	AfterFormat string
}

func (i ImageFormat) Convert() error {
	files, err := searchfile(i.Path, i.PreFormat)
	if err != nil {
		return err
	}
	//fmt.Println(files)

	for _, file := range files {
		fs, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fs.Close()

		img, _, err := image.Decode(fs)
		if err != nil {
			return err
		}

		dstpath := changeSuffixFile(file, i.AfterFormat)
		//fmt.Println(dstpath)
		out, err := os.Create(dstpath)
		if err != nil {
			return err
		}
		defer out.Close()

		switch i.AfterFormat {
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
	}
	return nil
}

func CreateImageFormat(path string, pre string, after string) *ImageFormat {
	res := new(ImageFormat)
	res.Path = path
	res.PreFormat = pre
	res.AfterFormat = after
	return res
}

func searchfile(path string, format string) ([]string, error) {
	var files []string

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+format {
				files = append(files, path)
			}
			return nil
		})
	return files, err
}

func changeSuffixFile(file string, format string) string {
	suffix := filepath.Ext(file)
	newfile := file[:len(file)-len(suffix)+1] + format
	return newfile
}
