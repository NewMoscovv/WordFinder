package main

import (
	"Excel_word_checker/internal/variator"
	lg "Excel_word_checker/pkg/logger"
)

func main() {

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
		variator.ExcelLogic(logger)
	}

}
