package variator

import (
	fndr "Excel_word_checker/internal/finder"
	"Excel_word_checker/internal/types/excel"

	lg "Excel_word_checker/pkg/logger"
)

func ExcelOptions(logger *lg.Logger) fndr.Finder {

	logger.Info.Print("Введите полный путь до файла Excel")

	xls, err := excel.New()
	if err != nil {
		logger.Info.Print(err.Error())
		return ExcelOptions(logger)
	}

	logger.Info.Print(xls.GetSheetList())

	err = xls.SetSheet()
	if err != nil {
		logger.Info.Print(err.Error())
		return ExcelOptions(logger)
	}

	return xls
}
