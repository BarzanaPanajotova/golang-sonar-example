package animal

type animal struct {
	words string
}

func New(words string) animal {
	return animal{words}
}

func (a animal) Speak() string {
	return a.words
}
