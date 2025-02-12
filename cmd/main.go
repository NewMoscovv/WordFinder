package main

import (
	fndr "Excel_word_checker/internal/finder"
	"Excel_word_checker/internal/variator"
	lg "Excel_word_checker/pkg/logger"
	"fmt"
)

func main() {
	var finder fndr.Finder
	var word string

	logger := lg.Init()
	logger.Info.Print("Добро пожаловать! Я помогу найти Вам слова в различных файлах.\n" +
		"Сейчас поддерживаются файлы: Excel")
	logger.Info.Print("Выберете, какие файлы желаете отсканировать." +
		"Для этого введите названия, что написаны выше☝️\n")

	fileType, err := variator.SetFileType()
	if err != nil {
		logger.Info.Print(err.Error())
	}

	switch fileType {
	case "excel":
		finder = variator.ExcelOptions(logger)
	default:
		logger.Info.Print("Жругие типы пока не поддерживаются")
		main()
	}

	logger.Info.Print("Введите необходимое слово\n")

	_, err = fmt.Scan(&word)
	if err != nil {
		logger.Err.Print(err.Error())
	}

	_, err = finder.FindWord(word)

}
