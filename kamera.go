package main

import (
	"fmt"
	"image/color"
	"runtime"

	"gocv.io/x/gocv"
)

func kamera() (*gocv.VideoCapture, *gocv.Window, *gocv.CascadeClassifier, error) {
	webcam, errCam := gocv.OpenVideoCapture(0)
	if errCam != nil {
		return nil, nil, nil, fmt.Errorf("Kamera Gagal Terbuka!")
	}
	window := gocv.NewWindow("WebCam")

	classifier := gocv.NewCascadeClassifier()
	if !classifier.Load("haarcascade_frontalface_alt.xml") {
		webcam.Close()
		window.Close()
		return nil, nil, nil, fmt.Errorf("Gagal Memuat file .xml")

	}
	return webcam, window, &classifier, nil
}
func konversiGambar(img gocv.Mat) ([]byte, error) {
	nativeBuf, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return nil, err
	}
	buf := nativeBuf.GetBytes()
	runtime.GC()
	return buf, nil
}
func gambarKotak(img gocv.Mat, classifier *gocv.CascadeClassifier) {
	faces := classifier.DetectMultiScale(img)
	for _, kotak := range faces {
		gocv.Rectangle(&img, kotak, color.RGBA{0, 255, 0, 0}, 3)
	}
}
