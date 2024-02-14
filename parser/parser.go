package parser

import (
	"errors"
	"strconv"
	"strings"
)

func GetInstructionType(instruction string) string {
	if instruction[0] == '@' {
		return "A"
	}
	if instruction[0] == '(' {
		return "L"
	}
	return "C"
}

func GetSymbol(instruction string) (string, error) {
	instructionType := GetInstructionType(instruction)
	value := strings.Trim(instruction, "@()")
	if value == "" {
		err := errors.New("A-Instruction or L-Instruction cannot be an empty string")
		return "", err
	}

	if instructionType == "A" {
		if decimalValue, err := strconv.Atoi(value); err == nil {
			if decimalValue < 0 || decimalValue > 32767 {
				err := errors.New(
					"A-Instruction must be a signed non negative 16-bit decimal number or a string",
				)
				return "", err
			}
		}
	}
	return value, nil
}

func GetComputeInstructionParts(cInstruction string) (dest, comp, jump string, err error) {
	parts := strings.Split(cInstruction, "=")
	if len(parts) == 1 {
		dest = "null"
	} else {
		dest = parts[0]
	}
	parts = strings.Split(parts[len(parts)-1], ";")
	comp = parts[0]
	if comp == "" {
		return "", "", "", errors.New("comp is mandatory in a C-Instruction")
	}
	if len(parts) == 1 {
		jump = "null"
	} else {
		jump = parts[1]
	}
	return
}
