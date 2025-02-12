package finder

type Finder interface {
	FindWord(word string) (string, error)
}
