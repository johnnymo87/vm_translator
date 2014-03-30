package main

import (
	"flag"
	. "github.com/johnnymo87/vm_translator/io"
	. "github.com/johnnymo87/vm_translator/parser"
	. "github.com/johnnymo87/vm_translator/writer"
)

func main() {
	filename := flag.String("filename", "", "a string *.vm")
	flag.Parse()
	path := ParsePath(*filename)
	lines, err := ReadLines(path + ".vm")
	if err != nil {
		panic(err)
	}
	vm_lines := Parse(lines)
	asm_lines := Translate(vm_lines)
	err = WriteLines(path+".asm", asm_lines)
	if err != nil {
		panic(err)
	}
}
