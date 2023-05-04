package yolobeer

import (
	"fmt"
	"time"
)

// YOLOResponse Neural net's response
type YOLOResponse struct {
	Beers   int
	Elapsed time.Duration
}

func (resp *YOLOResponse) String() string {
	result := ""
	result += fmt.Sprintf("Number of beers found : %d", resp.Beers)
	result += fmt.Sprintf("\nElapsed to find beers and read symbols: %v", resp.Elapsed)
	return result
}
