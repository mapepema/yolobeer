#!/bin/bash

# Create folder if not exists
mkdir -p data

# Download names
curl -o data/beers.name https://opencv-tutorial.readthedocs.io/en/latest/_downloads/a9fb13cbea0745f3d11da9017d1b8467/coco.names

# Download config file
curl -o data/beers.cfg https://opencv-tutorial.readthedocs.io/en/latest/_downloads/10e685aad953495a95c17bfecd1649e5/yolov3.cfg

# Download weights
curl -o data/beers.weights https://pjreddie.com/media/files/yolov3.weights
