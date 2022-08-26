package service

import "os"

type Image struct {
	FileName string
}
type Delete interface {
	Delete() error
}

func (img *Image) Delete() error {
	err := os.Remove(img.FileName)
	if err != nil {
		return err
	}
	return nil
}
