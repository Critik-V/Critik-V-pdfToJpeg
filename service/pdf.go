package service

import (
	"errors"
	"fmt"
	"go-pdf2jpeg/utils"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

const jpegQuality int = 7    // Quality of the JPEG image
const imgExt string = ".jpg" // Extension of
const docExt string = ".pdf" // Extension of the document

var (
	ErrPickingPage    = errors.New("error picking page")              // Error when picking the first page of the PDF
	ErrCreatingJpeg   = errors.New("error creating jpeg image")       // Error when creating the jpeg image
	ErrEncodingJpeg   = errors.New("error encoding jpeg image")       // Error when encoding the jpeg image
	ErrPdfDirNotExist = errors.New("pdf directory does not exist")    // Error when the pdf directory does not exist
	ErrImgDirCreation = errors.New("image directory creation failed") // Error when the image directory creation fails
)

/*
outPutFileName returns the name of the output file
by appending the image extension to the fileName
*/
func outPutFileName(fileName string) string {
	return fmt.Sprintf("%v%v", fileName, imgExt)
}

/*
pdfPath returns the path to the PDF file
by appending the document extension to the fileName
*/
func pdfPath(dir, fileName string) string {
	return fmt.Sprintf("%v/%v%v", dir, fileName, docExt)
}

/*
convert converts the first page of the PDF to a JPEG image
and saves it in the imgDir directory
*/
func convert(doc *fitz.Document, imgDir string, fileName string) error {
	img, err := doc.Image(0)
	if err != nil {
		return errors.Join(ErrPickingPage, err)
	}

	// Create a new file in the img directory
	f, err := os.Create(filepath.Join(imgDir, outPutFileName(fileName)))
	if err != nil {
		return errors.Join(ErrCreatingJpeg, err)
	}
	defer f.Close()

	// Encode the image to JPEG
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: int(jpegQuality)})
	if err != nil {
		return errors.Join(ErrEncodingJpeg, err)
	}
	return nil
}

/**
* PdfToJpeg converts the first page of the PDF to a JPEG image
* and saves it in the img directory
 */
func PdfToJpeg(fileName string) error {
	pdfDir := utils.GetPdfDir() // Path to the directory where the PDFs are stored
	imgDir := utils.GetImgDir() // Path to the directory where the images will be stored

	doc, err := fitz.New(pdfPath(pdfDir, fileName))
	if err != nil {
		return errors.Join(ErrPdfDirNotExist, err)
	}
	defer doc.Close()

	// Check if the img directory exists
	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		// img directory does not exist
		err = os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			return errors.Join(ErrImgDirCreation, err)
		}
		return convert(doc, imgDir, fileName)
	}
	// img directory already exists
	return convert(doc, imgDir, fileName)
}
