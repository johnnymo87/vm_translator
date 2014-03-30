package writer_test

import (
	. "github.com/johnnymo87/vm_translator/parser"
	. "github.com/johnnymo87/vm_translator/writer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	var lines []string
	BeforeEach(func() {
		_ = WriteLines("test.asm", []Command{Command("push constant 7")})
		lines = ReadLines("test.asm")
	})
	It("", func() {
		Ω(len(lines)).Should(Equal(7))
	})
})

var _ = Describe("", func() {
	It("", func() {
		asm := WritePushPop(Command("push constant 7"))
		Ω(len(asm)).Should(Equal(7))
	})
})
