package excel

import (
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	fileName *excelize.File
}

func New(fileName string) (*Excel, error) {

	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	return &Excel{file}, nil
}

func (e *Excel) FindWord(word string) (string, error) {
	return word, nil
}

func (e *Excel) FindWords(words []string) ([]string, error) {
	return words, nil
}
