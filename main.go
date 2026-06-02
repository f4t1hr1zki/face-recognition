package main

import (
	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
)

func main() {

	//Fungsi GoCV/Kamera
	webcam, window, classifier, errKam := kamera()
	if errKam != nil {
		println(errKam.Error())
		return
	}
	defer window.Close()
	defer webcam.Close()
	defer classifier.Close()

	//Fungsi  Ai/go-face
	rec, wajahAsli, errAi := siapkanAI()
	if errAi != nil {
		println(errAi.Error())
		return
	}
	defer rec.Close()

	//ForCam
	img := gocv.NewMat()
	defer img.Close()

	//FOR-IF CASE
	for {
		//GoCV
		webcam.Read(&img) //pengambilan gambar
		if img.Empty() {
			continue
		}
		//Gocv(perulangan untuk koversi)
		buf, errEncode := konversiGambar(img)
		if errEncode != nil {
			continue
		}
		//buf(gocv) -> rec(Go-Face)
		wajah, errFace := rec.Recognize(buf)
		if errFace == nil {
			for _, f := range wajah {
				jarak := face.SquaredEuclideanDistance(wajahAsli, f.Descriptor)
				if jarak < 0.6 {
					println("Akses Diterima", jarak)
				} else {
					println("Akses ditolak", jarak)
				}
			}
		}
		gambarKotak(img, classifier)

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
