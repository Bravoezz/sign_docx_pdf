package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

type TextElement struct {
	NumPage int
	Text    string
	X       float64 // Coordenada de la primera letra
	Y       float64 // Coordenada de la primera letra
}

func GetCoordinates(docPath, searchText string) (*TextElement, error) {
	f, r, err := pdf.Open(docPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var textElm TextElement
	foundText := false

	// Recorrer cada página del PDF
	for pageNum := 1; pageNum <= r.NumPage(); pageNum++ {
		page := r.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		// Obtener contenido de la página
		texts := page.Content().Text
		words := groupIntoWords(texts) // Agrupar caracteres en palabras

		// Buscar la palabra en la página
		for _, word := range words {
			fmt.Printf("Palabra: '%s' en (X: %.2f, Y: %.2f)\n", word.Text, word.X, word.Y)

			// Verificar si coincide con la palabra buscada
			if strings.Contains(strings.ReplaceAll(word.Text, " ", ""), strings.ReplaceAll(searchText, " ", "")) {
				println("hola: ", strings.ReplaceAll(word.Text, " ", ""), strings.ReplaceAll(searchText, " ", ""))
				fmt.Printf("¡Encontrado '%s' en (X: %.2f, Y: %.2f) en la página %d!\n", word.Text, word.X, word.Y, pageNum)
				textElm = TextElement{
					NumPage: pageNum,
					Text:    word.Text,
					X:       word.X,
					Y:       word.Y,
				}
				foundText = true
				//break // Detener la búsqueda con la primera coincidencia -- no detener para que llega a la ultima hoja siempre
			}
		}
	}

	if !foundText {
		return nil, errors.New("TEXTO NO ENCONTRADO: " + searchText)
	}

	return &textElm, nil
}

// 🔹 Función para agrupar caracteres en palabras basadas en la proximidad
func groupIntoWords(texts []pdf.Text) []TextElement {
	if len(texts) == 0 {
		return nil
	}

	var words []TextElement
	var currentWord TextElement
	currentWord.Text = ""
	isNewWord := true

	for i := 0; i < len(texts); i++ {
		t := texts[i]

		// Si es una nueva palabra, guardar coordenadas de la primera letra
		if isNewWord {
			currentWord = TextElement{
				Text: t.S,
				X:    t.X,
				Y:    t.Y,
			}
			isNewWord = false
		} else {
			// Si está cerca en X y en la misma línea, es parte de la palabra
			if abs(t.X-texts[i-1].X) < 10 && abs(t.Y-texts[i-1].Y) < 5 {
				currentWord.Text += t.S
			} else {
				// Guardar la palabra actual y empezar una nueva
				words = append(words, currentWord)
				currentWord = TextElement{
					Text: t.S,
					X:    t.X,
					Y:    t.Y,
				}
			}
		}
	}

	// Guardar la última palabra
	if len(currentWord.Text) > 0 {
		words = append(words, currentWord)
	}

	return words
}

// 🔹 Función auxiliar para valor absoluto
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
