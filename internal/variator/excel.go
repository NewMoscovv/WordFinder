package variator

import (
	fndr "Excel_word_checker/internal/finder"
	"Excel_word_checker/internal/types/excel"
	lg "Excel_word_checker/pkg/logger"
)

func ExcelLogic(logger *lg.Logger) fndr.Finder {

	var finder fndr.Finder

	logger.Info.Print("Введите путь до файла:\n")

	xls, err := excel.New()
	if err != nil {
		logger.Err.Println(err.Error())
	}

	logger.Info.Print(xls.GetSheetList())

	finder = xls

	return finder
}
