package main

import (
	lg "Excel_word_checker/pkg/logger"
	"fmt"
)

func main() {

	logger := lg.Init()
	logger.Info.Print("Добро пожаловать! Я помогу найти Вам слова в различных файлах.\n\n" +
		"Сейчас поддерживаются файлы: Excel")
	logger.Info.Print("Выберете, какие файлы желаете отсканировать." +
		"Для этого введите названия, что написаны выше☝️\n")

	var finderType string

	_, err := fmt.Scanln(&finderType)
	if err != nil {
		logger.Err.Print(err.Error())
	}

}
