package libs

import (
	"errors"
	"image/color"

	"gocv.io/x/gocv"
)

func WifuDetect(imagePath string, model string) (gocv.Mat, error) {
	// Read the image
	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		return img, errors.New("image not found")
	}
	defer img.Close()

	// Convert the image to grayscale
	imgGray := gocv.NewMat()
	defer imgGray.Close()
	gocv.CvtColor(img, &imgGray, gocv.ColorBGRToGray)

	// Equalize the histogram
	gocv.EqualizeHist(imgGray, &imgGray)

	// Load the cascade classifier
	faceCascade := gocv.NewCascadeClassifier()
	if !faceCascade.Load(model) {
		return img, errors.New("model not found")
	}
	defer faceCascade.Close()

	// Detect faces
	faces := faceCascade.DetectMultiScale(imgGray)

	// Draw rectangles around detected faces
	for _, r := range faces {
		gocv.Rectangle(&img, r, color.RGBA{255, 0, 255, 0}, 3)
	}

	return img, nil
}
