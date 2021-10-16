package main

import (
	"fmt"
	"golang-sonar-example/cmd/project/internal/animal"
)

type SpeakingAnimal interface {
	Speak() string
}

func main() {
	yorkie := animal.NewDog("I am a Yorkie, woff, woff!")
	golden := animal.NewDog("I am a Golden Retriever, woff, woff!")
	persian := animal.NewCat("I am a Persian Cat, meow!")

	animals := []SpeakingAnimal{yorkie, golden, persian}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
