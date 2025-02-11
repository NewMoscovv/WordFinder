package finder

type Finder interface {
	OpenFile(filename string) error
	FindWord(word string) (int, error)
	FindWords(words []string) ([]string, error)
}
