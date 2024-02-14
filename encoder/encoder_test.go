package encoder

import (
	"os"
	"testing"
)

var expectedRectAssemblyCode = []Code{
	{1, "@R0"},
	{2, "D=M"},
	{3, "@END"},
	{4, "D;JLE"},
	{5, "@n"},
	{6, "M=D"},
	{7, "@SCREEN"},
	{8, "D=A"},
	{9, "@addr"},
	{10, "M=D"},
	{11, "@addr"},
	{12, "A=M"},
	{13, "M=-1"},
	{14, "@addr"},
	{15, "D=M"},
	{16, "@32"},
	{17, "D=D+A"},
	{18, "@addr"},
	{19, "M=D"},
	{20, "@n"},
	{21, "M=M-1"},
	{22, "D=M"},
	{23, "@LOOP"},
	{24, "D;JGT"},
	{25, "@END"},
	{26, "0;JMP"},
}

func TestSanitizeAssemblyCode(t *testing.T) {
	code := SanitizeAssemblyCode("../hack-files/Rect.asm")

	if !symbolsTableContains("LOOP") || !symbolsTableContains("END") {
		t.Fail()
	}

	for i, v := range code {
		expected := expectedRectAssemblyCode[i]
		if v.lineNumber != expected.lineNumber || v.instruction != expected.instruction {
			t.Fail()
		}
	}
}

func TestEncoder(t *testing.T) {
	t.Run("Should encode correctly Rect.asm file", func(t *testing.T) {
		finalBinaryString, err := Encode(expectedRectAssemblyCode)
		if err != nil {
			t.Error(err)
		}
		file, err := os.ReadFile("../hack-files/Rect.hack")
		if err != nil {
			t.Error(err)
		}
		expectedBinary := string(file)
		if finalBinaryString != expectedBinary {
			t.Fail()
		}
	})
}
