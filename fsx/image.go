package fsx

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif" // use init to support decode jpeg,jpg,png,gif
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/webp"
)

const (
	maxImageSize = 8192 * 8192
)

// IsSupportedImageFile currently answers support image type is
// `image/jpeg, image/jpg, image/png, image/gif, image/webp`
func IsSupportedImageFile(localFilePath string) (bool, error) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(localFilePath), "."))
	switch ext {
	case "jpg", "jpeg", "png", "gif":
		// only allow for `image/jpeg,image/jpg,image/png, image/gif`
		ok, err := decodeAndCheckImageFile(localFilePath, standardImageConfigCheck)
		if !ok {
			return false, err
		}
		ok, err = decodeAndCheckImageFile(localFilePath, standardImageCheck)
		if !ok {
			return false, err
		}
	case "ico":
		// TODO: There is currently no good Golang library to parse whether the image is in ico format.
		return true, nil
	case "webp":
		ok, err := decodeAndCheckImageFile(localFilePath, webpImageConfigCheck)
		if !ok {
			return false, err
		}
		ok, err = decodeAndCheckImageFile(localFilePath, webpImageCheck)
		if !ok {
			return false, err
		}
	default:
		return false, errors.New("not support file type")
	}
	return true, nil
}

func decodeAndCheckImageFile(localFilePath string, checker func(io.Reader) error) (bool, error) {
	file, err := os.Open(localFilePath)
	if err != nil {
		return false, fmt.Errorf("open file error: %v", err)
	}
	defer file.Close()

	if err = checker(file); err != nil {
		return false, fmt.Errorf("check image format error: %v", err)
	}
	return true, nil
}

func standardImageConfigCheck(file io.Reader) error {
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("decode image config error: %v", err)
	}
	if imageSizeTooLarge(config) {
		return fmt.Errorf("image size too large")
	}
	return nil
}

func standardImageCheck(file io.Reader) error {
	_, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("decode image error: %v", err)
	}
	return nil
}

func webpImageConfigCheck(file io.Reader) error {
	config, err := webp.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("decode webp image config error: %v", err)
	}
	if imageSizeTooLarge(config) {
		return fmt.Errorf("image size too large")
	}
	return nil
}

func webpImageCheck(file io.Reader) error {
	_, err := webp.Decode(file)
	if err != nil {
		return fmt.Errorf("decode webp image error: %v", err)
	}
	return nil
}

func imageSizeTooLarge(config image.Config) bool {
	return config.Width*config.Height > maxImageSize
}
