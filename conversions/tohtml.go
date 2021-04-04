package conversions

import (
	"encoding/json"
	"html/template"
	"io"
	"os"

	"github.com/SergeyBibikov/gotest2html/custom_types"
)

func GenerateFailReportFromFile(rawFile string) error {
	toUnmarsh, err := FromFile2JsonArray(rawFile)
	if err != nil {
		return err
	}

	return parseAndCreate(toUnmarsh)
}
func GenerateFailReportFromStdin() error {
	toUnmarsh, err := FromStout2JsonArray(os.Stdin)
	if err != nil {
		return err
	}
	return parseAndCreate(toUnmarsh)
}
func parseAndCreate(toUnm []byte) error {
	var a []map[string]interface{}
	json.Unmarshal(toUnm, &a)
	results := GetGeneralResults(&a)
	sl := GetFailedTestsSlice(results)
	fort := custom_types.NewForTemplate(sl, results)
	tmp, err := template.New("base.html").ParseFiles("static/base.html", "static/case.html")
	if err != nil {
		return err
	}
	err = os.Mkdir("reports", 0755)
	if err != nil {
		return err
	}
	err = os.Chdir("reports")
	if err != nil {
		return err
	}
	fr, _ := os.Create("report_failed.html")
	defer fr.Close()

	tmp.Execute(fr, fort)
	createCss()
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

func createCss() {
	f, _ := os.Create("style.css")
	defer f.Close()
	toWr := `html{
		background-color: rgb(228, 225, 233);
		font-family: 'Courier New', Courier, monospace;
	}
	.fail{
		font-weight: bold;
		color: red;
	}
	.collapsible {
		background-color: rgb(233, 103, 81);
		color: white;
		cursor: pointer;
		padding: 18px;
		width: 100%;
		border: none;
		text-align: left;
		outline: none;
		font-size: 15px;
	  }

	  .active, .collapsible:hover {
		background-color: #555;
	  }

	  .content {
		padding: 0 10px 10px 10px;
		display: block;
		overflow: hidden;
		background-color: #f1f1f1;
	  }`
	io.WriteString(f, toWr)
}
