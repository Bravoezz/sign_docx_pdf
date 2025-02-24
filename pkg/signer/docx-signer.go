package signer

import (
	"baliance.com/gooxml/common"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"gosing-document/pkg/utils"
	"image"
	"os"
	"os/exec"
	"path"
	"strings"
)

type DocxSigner struct{}

func NewDocxSigner() SignStrategy {
	return new(DocxSigner)
}

func (sg DocxSigner) SignDocument(op SignOp) error {
	//aqui convertir el word a pdf
	basePath, _ := os.Getwd()
	convertPath := path.Join(basePath, "assets")
	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", op.InputPath, "--outdir", convertPath)
	outputStr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	println("convert to pdf: ", string(outputStr))

	splitStr := strings.Split(op.InputPath, "/")
	newNamePdf := strings.Replace(splitStr[len(splitStr)-1], ".docx", ".pdf", -1)
	println("location new pdf convert: ", path.Join(basePath, "assets", newNamePdf))
	txtFounded, err := utils.GetCoordinates(path.Join(basePath, "assets", newNamePdf), op.SearchText)
	if err != nil {
		return err
	}

	doc, _ := document.Open(op.InputPath)

	img, _ := doc.AddImage(common.Image{
		Size: image.Point{
			X: 120,
			Y: 70,
		},
		Format: path.Ext(op.SignaturePath)[1:],
		Path:   op.SignaturePath,
	})

	prg := doc.AddParagraph()
	run := prg.AddRun()
	drawing, _ := run.AddDrawingAnchored(img)

	x, y := sg.convertPDFToWordCoordinates(txtFounded.X, txtFounded.Y)

	drawing.SetOffset(x, y-75)
	drawing.SetTextWrapNone()

	if err := doc.SaveToFile(op.OutputPath); err != nil {
		return err
	}
	return nil
}

func (DocxSigner) convertPDFToWordCoordinates(pdfX, pdfY float64) (wordX, wordY measurement.Distance) {
	// Altura típica de una página A4 en puntos (11.69 pulgadas * 72 puntos/pulgada)
	const PAGE_HEIGHT_POINTS = 841.89
	// 1. Convertir coordenada X
	// La coordenada X en PDF ya está en el sistema correcto (izquierda a derecha)
	// Solo necesitamos convertirla a measurement.Distance
	wordX = measurement.Distance(pdfX)

	// 2. Convertir coordenada Y
	// Necesitamos:
	// a) Invertir el eje Y (porque en PDF crece hacia arriba y en Word hacia abajo)
	// b) Ajustar el origen (de abajo a arriba)
	invertedY := PAGE_HEIGHT_POINTS - pdfY
	wordY = measurement.Distance(invertedY)

	return wordX, wordY
}
