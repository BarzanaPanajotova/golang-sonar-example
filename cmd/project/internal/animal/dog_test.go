package animal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-sonar-example/cmd/project/internal/animal"
)

var _ = Describe("Application", func() {

	It("test", func() {
		app := animal.New("Woff")
		Expect(app.Speak()).To(Equal("Woff"))
	})
})
