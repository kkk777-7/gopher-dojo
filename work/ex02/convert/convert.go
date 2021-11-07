//Package convert image extension
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

// Target Image Format
type ImageFormat struct {
	// Target Path
	Path string
	// Previous Image Extension
	PreFormat string
	// After Image Extension
	AfterFormat string
}

// Convert convert the suffix of target files
//
// Input Value:  i ImageFormat
//
// Output Value:  error
func (i *ImageFormat) Convert() error {
	files, err := searchfile(i.Path, i.PreFormat)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("there is no target file : %s", i.PreFormat)
	}

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

// CreateImageFormat creates struct of pointer for Image Format.
//
// Input Value:  path (string): Starting path, pre (string): Previous converted suffix, after (string): After converted suffix
//
// Output Value:  (*ImageFormat): Target Image Format
func CreateImageFormat(path string, pre string, after string) *ImageFormat {
	res := new(ImageFormat)
	res.Path = path
	res.PreFormat = pre
	res.AfterFormat = after
	return res
}

// searchfile searches the path of target files.
//
// Input Value:  path (string): Starting path, format (string): Previous converted suffix
//
// Outout Value:  ([]string): Target files, error
func searchfile(path string, pre string) ([]string, error) {
	var files []string

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+pre {
				files = append(files, path)
			}
			return nil
		})
	return files, err
}

// changeSuffixFile changes suffix of the target path.
//
// Input Value:  file (string): The Target Path of File, format (string): The specified suffix
//
// Output Value:  (string): Changed Path of File
func changeSuffixFile(file string, format string) string {
	suffix := filepath.Ext(file)
	newfile := file[:len(file)-len(suffix)+1] + format
	return newfile
}
