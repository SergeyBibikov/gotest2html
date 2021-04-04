package main

import (
	"fmt"
	"os"

	"github.com/SergeyBibikov/gotest2html/conversions"
	"github.com/SergeyBibikov/gotest2html/generation"
)

func main() {
	generation.CreateStatic()
	defer generation.DeleteStatic()
	os.RemoveAll("reports")
	if len(os.Args) == 2 {
		filename := os.Args[1]
		err := conversions.GenerateFailReportFromFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := conversions.GenerateFailReportFromStdin()
		if err != nil {
			fmt.Println(err)
		}
	}
}
