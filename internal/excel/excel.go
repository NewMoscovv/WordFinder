package excel

import (
	lg "Excel_word_checker/pkg/logger"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	file *excelize.File
}

func New(fileName string) (*Excel, error) {

	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	return &Excel{file}, nil
}

func (e *Excel) GetSheetList(logger lg.Logger) {
	var sheetString string

	sheetList := e.file.GetSheetList()

	if len(sheetList) == 1 {
		sheetString = sheetList[0]
	} else {
		for _, sheet := range sheetList {
			sheetString = fmt.Sprintf("%s ", sheet)
		}
	}

	logger.Info.Printf("Обнаружено %d страниц: %s. Выберите какой лист необходимо просканировать.\n"+
		"Для этого необходимо написать название листа или \"ВСЕ\" для сканирования всех листов",
		len(sheetList), sheetString)

}

func (e *Excel) FindWord(word string) (string, error) {
	return word, nil
}

func (e *Excel) FindWords(words []string) ([]string, error) {
	return words, nil
}
