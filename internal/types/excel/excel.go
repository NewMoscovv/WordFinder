package excel

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

type Excel struct {
	file  *excelize.File
	sheet string
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

	return &Excel{file, ""}, nil
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

func (e *Excel) SetSheet() error {
	var chosenSheet string

	_, err := fmt.Scan(&chosenSheet)
	if err != nil {
		e.sheet = chosenSheet
		return err
	}

	sheetList := e.file.GetSheetList()
	if chosenSheet == "все" {
		e.sheet = "все"
	} else {
		for _, sheet := range sheetList {

			if strings.ToLower(sheet) == strings.ToLower(chosenSheet) {
				e.sheet = sheet
				break
			}
		}
	}

	if e.sheet == "" {
		return errors.New("такого листа нету в списке")
	}

	return nil
}

func (e *Excel) FindWord(word string) (string, error) {

	if e.sheet != "все" {
		rows, err := e.file.GetRows(e.sheet)
		if err != nil {
			return "", err
		}

		for row, _ := range rows {
			for _, cell := range rows[row] {
				if strings.Contains(cell, word) {
					fmt.Print(cell)
				}
			}
		}

	}

	return word, nil

}
