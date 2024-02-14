package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"HackAssembler/encoder"
)

func GetFileName() (fileName string) {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Please, specify the .asm file to proceed")
	}
	fileName = args[0]
	extension := filepath.Ext(fileName)
	if extension != ".asm" {
		log.Fatal("Please, provide a valid .asm file")
	}
	return
}

func main() {
	path := GetFileName()
	assemblyCode := encoder.SanitizeAssemblyCode(path)
	finalBinaryString, err := encoder.Encode(assemblyCode)
	if err != nil {
		log.Fatal(err)
	}
	basepath := filepath.Base(path)
	outputFile := strings.TrimSuffix(basepath, ".asm") + ".hack"
	os.WriteFile(outputFile, []byte(finalBinaryString), 0644)
}
