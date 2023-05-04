package yolobeer

import (
	"fmt"
	"image"

	"github.com/LdDl/go-darknet"
)

func (net *YOLONetwork) detectBeers(imgSrc image.Image) (int, error) {
	img, err := darknet.Image2Float32(imgSrc)
	if err != nil {
		return 0, err
	}
	dr, err := net.Beers.Detect(img)
	if err != nil {
		return 0, err
	}
	img.Close()
	var rects []image.Rectangle
	for _, d := range dr.Detections {
		for i := range d.ClassIDs {
			if d.ClassNames[i] != "car" && d.ClassNames[i] != "motorbike" && d.ClassNames[i] != "bus" && d.ClassNames[i] != "train" && d.ClassNames[i] != "truck" {
				// I think this is excess condition...
				// continue
			}
			bBox := d.BoundingBox
			minX, minY := float64(bBox.StartPoint.X), float64(bBox.StartPoint.Y)
			maxX, maxY := float64(bBox.EndPoint.X), float64(bBox.EndPoint.Y)
			rect := image.Rect(round(minX), round(minY), round(maxX), round(maxY))
			rects = append(rects, rect)
		}
	}
	fmt.Println(rects)
	return len(dr.Detections), nil
}
