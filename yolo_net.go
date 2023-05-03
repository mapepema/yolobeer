package yolobeer

import (
	"github.com/LdDl/go-darknet"
)

const (
	gpuIndex = 0
)

// NewYOLONetwork return pointer to YOLONetwork
func NewYOLONetwork(beersCfg, beersWeights string, beersThreshold float32) (*darknet.YOLONetwork, error) {
	beers := darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		WeightsFile:              beersWeights,
		NetworkConfigurationFile: beersCfg,
		Threshold:                beersThreshold,
	}
	if err := beers.Init(); err != nil {
		return nil, err
	}
	return &beers, nil
}
