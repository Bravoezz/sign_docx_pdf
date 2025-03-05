package main

import (
	"fmt"
	"gosing-document/pkg/signer"
	"log"
	"os"
	"path"
)

func main() {
	basePath, _ := os.Getwd()
	const (
		docPath    = "renuncia.pdf"
		outputPath = "output-sign.pdf"
		imagePath  = "firma-final.png"
		searchText = "Bravo Ramos Joel Brayan"
	)

	signerDoc := signer.NewSigner(&signer.SignOp{
		InputPath:     path.Join(basePath, "assets", docPath),
		OutputPath:    path.Join(basePath, "store", outputPath),
		SignaturePath: path.Join(basePath, "assets", imagePath),
		SearchText:    searchText,
	})

	/**
	if path.Ext(docPath) == ".docx" {
		signerDoc.SetSignStrategy(signer.NewDocxSigner())
	} else if path.Ext(docPath) == ".pdf" {
		signerDoc.SetSignStrategy(signer.NewPdfSigner())
	}
	**/

	strategyMap := map[string]signer.SignStrategy{
		".docx": signer.NewDocxSigner(),
		".pdf":  signer.NewPdfSigner(),
	}
	signerDoc.SetSignStrategy(strategyMap[path.Ext(docPath)])

	err := signerDoc.Sign()
	if err != nil {
		log.Fatalf("Error al agregar la imagen: %v", err)
	}

	fmt.Println("✅ Imagen agregada correctamente en:", outputPath)

	println("***EXIT PROGRAM***")
}
