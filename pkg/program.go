package main

import (
	"flag"
	"fmt"
	"gosing-document/pkg/signer"
	"log"
	"os"
	"path"
)

var docF, outputF, imageF, searchTextF string

func init() {
	flag.StringVar(&docF, "docpath", "nil", "flag -docpath required")
	flag.StringVar(&outputF, "outputpath", "nil", "flag -outputpath required")
	flag.StringVar(&imageF, "imagepath", "nil", "flag -imagepath required")
	flag.StringVar(&searchTextF, "searchtext", "nil", "flag -searchtext required")
	flag.Parse()

	if docF == "nil" || outputF == "nil" || imageF == "nil" || searchTextF == "nil" {
		fmt.Println("All flags required")
		flag.Usage()
		os.Exit(1)
	}
}

// TODO ESTAN FUNCIONANDO BIEN, AHORA TOCA PROBARO JUNTO CON EL EL ACONEX API
func main() {
	/**
	basePath, _ := os.Getwd()
	const (
		docPath    = "renuncia-single-sign.pdf"
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
		**/

	signerDoc := signer.NewSigner(&signer.SignOp{
		InputPath:     docF,
		OutputPath:    outputF,
		SignaturePath: imageF,
		SearchText:    searchTextF,
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
	signerDoc.SetSignStrategy(strategyMap[path.Ext(docF)])

	err := signerDoc.Sign()
	if err != nil {
		log.Fatalf("Error al agregar la imagen: %v", err)
	}

	fmt.Println("✅ Imagen agregada correctamente en:", outputF)

	println("***EXIT PROGRAM***")
}
