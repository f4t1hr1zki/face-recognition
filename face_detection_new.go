package main
import (
	"image/color"
	"gocv.io/x/gocv"
	"github.com/Kagami/go-face"
)
func main () {
	webcam, err := gocv.OpenVideoCapture(0)
		//window webcam
		window := gocv.NewWindow("Webcam")
		if err != nil {
			println("Kamera Gagal Terbuka!")
			return 
		}
		defer webcam.Close()
//ClassifierForCascade
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
		img := gocv.NewMat()
		
//FaceRecoginzerWithKagamiGo-Face
rec, err := face.NewRecognizer("models")
	if err != nil {
		println("Gagal Menginilisiasi Go-Face", err.Error())
		return
	}
	defer rec.Close()
	
//LoadXMLFile
	classifier.Load("haarcascade_frontalface_alt.xml")
		defer window.Close()
		defer img.Close()
//DataFotoSampel
	fotoSampel, err := rec.RecognizeSingleFile("foto/foto_sendiri.jpeg")
//	fotoSampel, err := rec.RecognizeSingleFile("foto/kayaba_akihiko.jpg")
		if err != nil || fotoSampel == nil {
			println("Foto Sampel tidak mendeteksi wajah")
			return
		}
wajahAsli := fotoSampel.Descriptor
//PerulanganFor
	for {
		webcam.Read(&img)
			if img.Empty() {
				continue
			}
		//ConvertImgToByte(encode-decode)
		nativeBuf, _ := gocv.IMEncode(".jpg", img)
		if err != nil {
			continue
		}
		buf := nativeBuf.GetBytes()
		defer nativeBuf.Close()
		//Recognize 'buf'
		wajah, err := rec.Recognize(buf)
		if err != nil {
			continue
		}
		if len (wajah) > 0 {
			println("Ada Wajah Terdeteksi!")
		}

		for _, f := range wajah {
			jarak := face.SquaredEuclideanDistance(wajahAsli, f.Descriptor)
		if jarak  < 0.6 {
			println("Akses Diterima,Jarak kemiripan:", jarak)
		} else {
			println("Akses Ditolak,Jarak kemiripan:", jarak)
		}
		
		}
		
		faces := classifier.DetectMultiScale(img)
		for _, kotak := range faces {
			gocv.Rectangle(&img, kotak, color.RGBA{0, 255, 0, 0}, 3)
		}
			window.IMShow(img)
			if window.WaitKey(1) >= 0 {
				break
			}
		
		
	}
	
}
