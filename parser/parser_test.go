package parser

import (
	"testing"
)

func TestGetInstructionType(t *testing.T) {
	t.Run("Should return an A-instruction type", func(t *testing.T) {
		var instructionType string
		instructionType = GetInstructionType("@2")
		if instructionType != "A" {
			t.Fail()
		}
		instructionType = GetInstructionType("@R0")
		if instructionType != "A" {
			t.Fail()
		}
	})

	t.Run("Should return an L-instruction type", func(t *testing.T) {
		instructionType := GetInstructionType("(LOOP)")
		if instructionType != "L" {
			t.Fail()
		}
	})

	t.Run("Should return an C-instruction type", func(t *testing.T) {
		var instructionType string
		instructionType = GetInstructionType("D=D+1;JLE")
		if instructionType != "C" {
			t.Fail()
		}
		instructionType = GetInstructionType("D+1;JLE")
		if instructionType != "C" {
			t.Fail()
		}
		instructionType = GetInstructionType("M=D+1")
		if instructionType != "C" {
			t.Fail()
		}
	})
}

func TestGetSymbol(t *testing.T) {
	t.Run("Should return an A-instruction Symbol", func(t *testing.T) {
		symbol, err := GetSymbol("@2")
		if err != nil {
			t.Error(err)
		}
		if symbol != "2" {
			t.Fail()
		}
	})

	t.Run("Should return an error if A-instruction value is LT 0 or GT 32767", func(t *testing.T) {
		_, errGreater := GetSymbol("@32768")
		if errGreater == nil {
			t.Fail()
		}
		_, errNegative := GetSymbol("@-1")
		if errNegative == nil {
			t.Fail()
		}
	})

	t.Run("Should return an L-instruction Symbol", func(t *testing.T) {
		symbol, err := GetSymbol("(LOOP)")
		if err != nil {
			t.Error(err)
		}
		if symbol != "LOOP" {
			t.Fail()
		}
	})
}

func TestGetComputeInstructionParts(t *testing.T) {
	t.Run("Should parse dest=comp;jump instruction", func(t *testing.T) {
		cInstruction := "D=M;JGT"
		dest, comp, jump, err := GetComputeInstructionParts(cInstruction)
		if err != nil {
			t.Error(err)
		}
		if dest != "D" || comp != "M" || jump != "JGT" {
			t.Fail()
		}
	})

	t.Run("Should parse dest=comp instruction", func(t *testing.T) {
		cInstruction := "D=M"
		dest, comp, jump, err := GetComputeInstructionParts(cInstruction)
		if err != nil {
			t.Error(err)
		}
		if dest != "D" || comp != "M" || jump != "null" {
			t.Fail()
		}
	})

	t.Run("Should parse comp;jump instruction", func(t *testing.T) {
		cInstruction := "0;JMP"
		dest, comp, jump, err := GetComputeInstructionParts(cInstruction)
		if err != nil {
			t.Error(err)
		}
		if dest != "null" || comp != "0" || jump != "JMP" {
			t.Fail()
		}
	})

	t.Run("Should return an error if comp is ommited instruction", func(t *testing.T) {
		cInstruction := "M=;JMP"
		_, _, _, err := GetComputeInstructionParts(cInstruction)
		if err == nil {
			t.Error(err)
		}
	})
}
