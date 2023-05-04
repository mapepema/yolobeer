package yolobeer

import (
	"image"
	"time"
)

func (net *YOLONetwork) ReadBeers(imgSrc image.Image, saveCrop bool) (*YOLOResponse, error) {
	resp := YOLOResponse{}
	st := time.Now()
	beers, err := net.detectBeers(imgSrc)
	if err != nil {
		return nil, err
	}
	resp.Beers = beers
	resp.Elapsed = time.Since(st)
	return &resp, nil
}
