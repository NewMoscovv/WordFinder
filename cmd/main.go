package main

import (
	fndr "Excel_word_checker/internal/finder"
	"Excel_word_checker/internal/variator"
	lg "Excel_word_checker/pkg/logger"
)

func main() {
	var finder fndr.Finder

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
	}

	_, err = finder.FindWord("o")

}
