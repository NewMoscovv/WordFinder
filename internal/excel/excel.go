package excel

type Excel struct {
	fileName string
}

func New(fileName string) *Excel {
	return &Excel{fileName: fileName}
}

func (e *Excel) FindWord(word string) (string, error) {
	return word, nil
}

func (e *Excel) FindWords(words []string) ([]string, error) {
	return words, nil
}
