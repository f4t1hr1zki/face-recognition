
package main
import (
	"image/color"
	"gocv.io/x/gocv"
)

func main () {
	webcam, err := gocv.OpenVideoCapture(0)
	//Window Untuk Webcam
		window := gocv.NewWindow("Webcam")
		if err != nil {
					println("Kamera Gagal")
					return
				}
				defer webcam.Close()
	//
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
		img := gocv.NewMat()

	//LoadXMLClassifier
		classifier.Load("haarcascade_frontalcatface.xml")

			defer window.Close()

			defer img.Close()
		for {
					webcam.Read(&img)
		
					if img.Empty() {
						continue
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


/*func main() {
	img := gocv.IMRead("foto_sendiri.jpeg", gocv.IMReadGrayScale)
	defer img.Close()

	window := gocv.NewWindow("Image")
	defer window.Close()

	window.IMShow(img)
	window.WaitKey(0)
	
}
*/
