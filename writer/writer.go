package writer

import (
	"fmt"
	"regexp"
	. "github.com/johnnymo87/vm_translator/parser"
)

var push, pop = regexp.MustCompile(`(?i)push`), regexp.MustCompile(`(?i)pop`)

func Translate(lines []string) []string {
	var asm_lines []string
	for _, line := range lines {
		typ, _ := Type(line)
		switch typ {
		case "ARITHMETIC":
			asm_lines = append(asm_lines, WriteArithmetic(line)...)
		case "LOGICAL":
			asm_lines = append(asm_lines, WriteLogic(line)...)
		case "PUSHPOP":
			asm_lines = append(asm_lines, WritePushPop(line)...)
		default:
			fmt.Println("Unknown command type for line '%v'\n", line)
			panic(line)
		}
	}
	return append(asm_lines, Finally()...)
}

func WritePushPop(command string) []string {
	comm, _, num, err := ReadPushPop(command)
	cons := fmt.Sprintf("@%d", num)
	if err != nil {
		panic(err)
	}
	switch {
	case push.MatchString(comm):
		return []string{cons, "D=A", "@SP", "A=M", "M=D", "@SP", "M=M+1"}
	case pop.MatchString(comm):
		fmt.Println("Support for POP command not implemented!\n")
		panic(comm)
	default:
		fmt.Println("Command '%v' is not a valid PUSHPOP command\n", comm)
		panic(comm)
	}
}

func WriteArithmetic(command string) []string {
	switch command {
	case "add":
		return append(PopTwo(), "D=D+M")
	case "sub":
		return append(PopTwo(), "D=D-M")
	case "neg":
		return append(PopOne(), []string{"@0", "D=A-D"}...)
	default:
		fmt.Println("Command '%v' is not a valid ARITHMETIC command\n", command)
		panic(command)
	}
}

func WriteLogic(command string) []string {
	base := []string{"D=D-M", "@TRUE", "", "@FALSE", "0;JMP", "(TRUE)", "@-1", "D=A", "@END", "0;JMP", "(FALSE)", "@0", "D=A"}
	switch command {
	case "eq":
		base[2] = "D;JEQ"
		return base
	case "lt":
		base[2] = "D;JGT"
		return base
	case "gt":
		base[2] = "D;JLT"
		return base
	default:
		fmt.Println("Command '%v' is not a valid LOGICAL command\n", command)
		panic(command)
	}
}

func WriteBitWise(command string) []string {
	switch command {
	case "and":
		return append(PopTwo(), "D=D&A")
	case "or":
		return append(PopTwo(), "D=D|A")
	case "not":
		return append(PopOne(), "D=!D")
	default:
		fmt.Println("Command '%v' is not a valid LOGICAL command\n", command)
		panic(command)
	}
}

// Pop one to D
func PopOne() []string {
	return []string{"@SP", "M=M-1", "A=M", "D=M"}
}

// Pop one to D, then one to M
func PopTwo() []string {
	return []string{"@SP", "M=M-1", "A=M", "D=M", "@SP", "M=M-1", "A=M"}
}

// Push D to stack, increment pointer
func Finally() []string {
	return []string{"@END", "0;JMP", "(END)", "@SP", "A=M", "M=D", "@SP", "M=M+1"}
}
