package utils

import (
	"errors"
	"fmt"
	"rsc.io/pdf"
	"sort"
	"strings"
)

type TextElement struct {
	NumPage int
	Text    string
	X       float64
	Y       float64
}

func GetCoordinates(docPath, searchText string) (*TextElement, error) {
	var textElm TextElement
	pdfFile, _ := pdf.Open(docPath)

	foundText := false
	for pageNum := 1; pageNum <= pdfFile.NumPage(); pageNum++ {
		page := pdfFile.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		content := page.Content()
		words := groupIntoWords(content.Text)

		for _, word := range words {
			fmt.Printf("Palabra: '%s' en (X: %.2f, Y: %.2f)\n", word.Text, word.X, word.Y)

			// Buscar coincidencias
			if strings.Contains(strings.ToLower(word.Text), strings.ToLower(strings.Replace(searchText, " ", "", -1))) {
				fmt.Printf("¡Encontrado '%s' en (X: %.2f, Y: %.2f)!\n", word.Text, word.X, word.Y)
				textElm.NumPage, textElm.X, textElm.Y, textElm.Text = pageNum, word.X, word.Y, word.Text
				foundText = true
			}
		}
	}
	if !foundText {
		return nil, errors.New("TEXTO NO ENCONTRADO " + searchText)
	}

	return &textElm, nil
}

// Función para agrupar caracteres en palabras basándose en su posición
func groupIntoWords(texts []pdf.Text) []TextElement {
	if len(texts) == 0 {
		return nil
	}

	// Ordenar los textos por Y (línea) y luego por X (posición horizontal)
	sort.Slice(texts, func(i, j int) bool {
		// Si están en la misma línea (aproximadamente)
		if abs(texts[i].Y-texts[j].Y) < 1 {
			return texts[i].X < texts[j].X
		}
		return texts[i].Y > texts[j].Y
	})

	var words []TextElement
	currentWord := TextElement{
		X: texts[0].X,
		Y: texts[0].Y,
	}

	for i := 0; i < len(texts); i++ {
		// Si es el primer carácter o está cerca del anterior (misma palabra)
		if i == 0 || (abs(texts[i].Y-texts[i-1].Y) < 1 && texts[i].X-texts[i-1].X < 15) {
			currentWord.Text += texts[i].S
		} else {
			// Guardar la palabra actual y empezar una nueva
			if len(currentWord.Text) > 0 {
				words = append(words, currentWord)
			}
			currentWord = TextElement{
				Text: texts[i].S,
				X:    texts[i].X,
				Y:    texts[i].Y,
			}
		}
	}

	// Agregar la última palabra
	if len(currentWord.Text) > 0 {
		words = append(words, currentWord)
	}

	return words
}

// Función auxiliar para valor absoluto
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
