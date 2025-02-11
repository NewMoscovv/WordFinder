package finder

type Finder interface {
	FindWord(word string) (string, error)
	FindWords(words []string) ([]string, error)
}
