package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"go-reloaded/tools"
)

func Read_Input(infile_name string) string {
	buffer, err := os.ReadFile(infile_name)
	tools.CheckError(err, "Error: failed to read infile: "+infile_name)
	return string(buffer)
}

func Write_Output(result, outfile_name string) {
	if outfile_name == "main.go" {
		log.Fatal("Error: cannot write to a source code file\n")
	}
	outfile, err := os.Create(outfile_name)
	tools.CheckError(err, "Error: cannot write to the outfile: "+outfile_name)
	outfile.WriteString(result)
}

func main() {
	if len(os.Args[1:]) == 2 {
		text := Read_Input(os.Args[1])
		if text == "" {
			log.Fatal("Error: input file is empty\n")
		}
		result := tools.Parse_Text(text)
		Write_Output(strings.TrimSpace(result), os.Args[2])
	} else {
		fmt.Fprintf(os.Stderr, "Usage: go run . <Infile> <Outfile>\n")
	}
}
