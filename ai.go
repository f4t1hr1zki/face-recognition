package main

import (
	"fmt"

	"github.com/Kagami/go-face"
)

func siapkanAI() (*face.Recognizer, face.Descriptor, error) {
	rec, err := face.NewRecognizer("models/")
	if err != nil {
		return nil, face.Descriptor{}, err
	}
	fotoSampel, errPh := rec.RecognizeSingleFile("foto_sendiri.jpeg")
	if errPh != nil || fotoSampel == nil {
		rec.Close()
		return nil, face.Descriptor{}, fmt.Errorf("Foto sampel tidak mendeteksi wajah")

	}
	return rec, fotoSampel.Descriptor, nil
}
