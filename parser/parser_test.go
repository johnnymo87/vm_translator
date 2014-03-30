package parser_test

import (
	. "github.com/johnnymo87/vm_translator/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IO", func() {
	var lines = ReadLines("../data/StackArithmetic/SimpleAdd/SimpleAdd.vm")
	It("Parses 3 commands from SimpleAdd.vm", func() {
		Ω(len(Parse(lines))).Should(Equal(3))
	})
})
