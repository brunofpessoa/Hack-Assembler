package encoder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"HackAssembler/parser"
)

type Code struct {
	lineNumber  int
	instruction string
}

func symbolsTableContains(symbol string) bool {
	_, contains := SymbolTable[symbol]
	return contains
}

func SanitizeAssemblyCode(fileName string) (code []Code) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "//")
		line = parts[0]
		line = strings.ReplaceAll(line, " ", "")
		if len(line) == 0 {
			continue
		}

		instructionType := parser.GetInstructionType(line)
		if instructionType == "L" {
			symbol, err := parser.GetSymbol(line)
			if err != nil {
				log.Fatal(err)
			}
			SymbolTable[symbol] = fmt.Sprint(lineNumber)
		} else {
			lineNumber++
			c := Code{lineNumber: lineNumber, instruction: line}
			code = append(code, c)
		}
	}

	return
}

func Encode(code []Code) (finalBinaryString string, err error) {
	variableCount := 0
	for _, c := range code {
		instructionType := parser.GetInstructionType(c.instruction)

		if instructionType == "A" {
			symbol, err := parser.GetSymbol(c.instruction)
			if err != nil {
				return "", err
			}
			_, isNaN := strconv.Atoi(symbol)
			if !symbolsTableContains(symbol) && isNaN != nil {
				SymbolTable[symbol] = fmt.Sprint(16 + variableCount)
				variableCount++
			}
			hasSymbol := symbolsTableContains(symbol)
			var binary string
			if hasSymbol {
				symbolValue := SymbolTable[symbol]
				binary = decimalToBinary(symbolValue, c.lineNumber)
			} else {
				binary = decimalToBinary(symbol, c.lineNumber)
			}
			finalBinaryString += binary + "\n"
			continue
		}
		if instructionType == "C" {
			dest, comp, jump, err := parser.GetComputeInstructionParts(c.instruction)
			if err != nil {
				return "", err
			}
			cInstructionBin := fmt.Sprintf(
				"111%s%s%s\n",
				CompInstruction[comp],
				DestInstruction[dest],
				JumpInstruction[jump],
			)
			finalBinaryString += cInstructionBin
			continue
		}
	}
	return
}

func decimalToBinary(decimal string, lineNumber int) string {
	n, err := strconv.Atoi(decimal)
	if err != nil {
		msg := fmt.Sprintf("Cannot convert %s to Int. Line: %v", decimal, lineNumber)
		log.Fatal(msg)
	}
	binaryString := strconv.FormatInt(int64(n), 2)
	padding := 16 - len(binaryString)
	for i := 0; i < padding; i++ {
		binaryString = "0" + binaryString
	}
	return binaryString
}
