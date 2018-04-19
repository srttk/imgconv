//package will be controll convert.
package converter

import (
	"os"
	"image"
	"image/png"
	"image/jpeg"
	"image/gif"
	"path/filepath"
	"strings"
	"errors"
)

/*
Images contains filepath without extension and image.Image

Execution is very slow, so please avoid using it on actual programs
 */
type Images map[string]image.Image //key is fipepath without extension, val is image.Image

var images Images

func init() {
	images = Images{}
}

/*
"srcType" is file Extension without dot. It supports only "jpg", "jpeg", "png", "gif" extension.
Convert only images matching "srcType"

"srcDir" is directory in which the file you want to convert exists. ex, "./test_images"

NewImages is a function that returns images.
If problems occur during generation it returns err.
 */
func NewImages(srcType string, srcDir string) (Images, error) {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != "."+srcType {
			return nil
		}
		fileName := strings.TrimSuffix(path, filepath.Ext(path))
		img, err := getImg(path)
		if err != nil {
			return err
		}
		images[fileName] = img
		return nil
	})
	return images, err
}

func getImg(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

/*
"outType" is file Extension without dot. it supports only "jpg", "jpeg", "png", "gif" extension.

ConvertTo is a method to convert images.
It returns err if a problem occurs during processing.

This creates a new image in the same directory without destroying the original image.
 */
func (i Images) ConvertTo(outType string) error {
	for fileName, img := range i {
		out, err := os.Create(fileName+"."+outType)
		if err != nil {
			return err
		}
		switch outType {
		case "jpeg", "jpg":
			err = jpeg.Encode(out, img, nil)
			if err != nil {
				return err
			}
		case "png":
			err = png.Encode(out, img)
			if err != nil {
				return err
			}
		case "gif":
			err = gif.Encode(out, img, nil)
			if err != nil {
				return err
			}
		default:
			return errors.New("sorry. not support this outType extend")
		}
	}
	return nil
}
