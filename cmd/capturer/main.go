package main

import (
	"flag"
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

var (
	deviceID int
	filename string
)

func init() {
	flag.IntVar(&deviceID, "d", 0, "Number of Device ID")
	flag.StringVar(&filename, "f", "image.png", "Name of output image filename")
	flag.Parse()
}

func main() {
	webcam, err := gocv.OpenVideoCapture(deviceID)
	handleError(err)
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		handleError(fmt.Errorf("cannot read device %d", deviceID))
		return
	}
	if img.Empty() {
		handleError(fmt.Errorf("no image on device %d", deviceID))
		return
	}

	if ok := gocv.IMWrite(filename, img); !ok {
		handleError(fmt.Errorf("failed to write image %s", filename))
		return
	}

	log.Println(filename)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
