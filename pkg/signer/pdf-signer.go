package signer

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"gosing-document/pkg/utils"
	"log"
	"strconv"
)

type PdfSigner struct{}

func NewPdfSigner() SignStrategy {
	return new(PdfSigner)
}

func (PdfSigner) SignDocument(op SignOp) error {
	txtElemet, err := utils.GetCoordinates(op.InputPath, op.SearchText)
	if err != nil {
		return err
	}
	// Configurar la imagen como marca de agua
	confStr := fmt.Sprintf("pos:bl, scale: 0.3, off:%.2f %.2f, rot:0", txtElemet.X, txtElemet.Y)
	wm, err := api.ImageWatermark(op.SignaturePath, confStr, true, false, types.POINTS)
	if err != nil {
		log.Fatalf("Error creando marca de agua: %v", err)
	}

	// Aplicar la imagen al PDF
	err = api.AddWatermarksFile(op.InputPath, op.OutputPath, []string{strconv.Itoa(txtElemet.NumPage)}, wm, nil)
	return nil
}
