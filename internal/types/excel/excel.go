package excel

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

type Excel struct {
	file *excelize.File
}

func New() (*Excel, error) {
	reader := bufio.NewReader(os.Stdin)

	fileName, err := reader.ReadString('\r')
	fileName = strings.TrimSpace(fileName)
	if err != nil {
		return nil, err
	}

	file, err := excelize.OpenFile(fileName)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return &Excel{file}, nil
}

func (e *Excel) GetSheetList() string {
	var sheetString string

	sheetList := e.file.GetSheetList()

	if len(sheetList) == 1 {
		sheetString = sheetList[0]
	} else {
		for _, sheet := range sheetList {
			sheetString = fmt.Sprintf("%s ", sheet)
		}
	}

	result := fmt.Sprintf("Обнаружено %d страниц: %s. Выберите какой лист необходимо просканировать.\n"+
		"Для этого необходимо написать название листа или \"ВСЕ\" для сканирования всех листов\n",
		len(sheetList), sheetString)
	return result
}

func (e *Excel) FindWord(word string) (string, error) {
	return word, nil
}

func (e *Excel) FindWords(words []string) ([]string, error) {
	return words, nil
}
