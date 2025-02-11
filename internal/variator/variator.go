package variator

import (
	"errors"
	"fmt"
	"strings"
)

func SetFileType() (string, error) {

	var fileType string

	_, err := fmt.Scan(&fileType)
	if err != nil {
		return "", err
	}

	fileType = strings.ToLower(fileType)

	switch fileType {
	case "excel":
		return fileType, nil
	}
	return "", errors.New("invalid file type")
}
