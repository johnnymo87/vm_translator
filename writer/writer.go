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
			break
		case "PUSHPOP":
			asm_lines = append(asm_lines, WritePushPop(line)...)
		default:
			fmt.Println("Unknown command type for line '%v'\n", line)
			panic(line)
		}
	}
	return asm_lines
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
		return []string{"@SP", "M=M-1", "A=M", "D=M", "@SP", "M=M-1", "A=M", "D=D+M", "@SP", "A=M", "M=D", "@SP", "M=M+1"}
	case "sub":
		return []string{"@SP", "M=M-1", "A=M", "D=M", "@SP", "M=M-1", "A=M", "D=D-M", "@SP", "A=M", "M=D", "@SP", "M=M+1"}
	case "neg":
		return []string{"@SP", "M=M-1", "A=M", "D=M", "@0", "D=A-D", "@SP", "A=M", "M=D", "@SP", "M=M+1"}
	default:
		fmt.Println("Command '%v' is not a valid ARITHMETIC command\n", command)
		panic(command)
	}
}
