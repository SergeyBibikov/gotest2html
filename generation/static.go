package generation

import (
	"io"
	"os"
	"sync"
)

func CreateStatic() {
	os.Mkdir("static", 0755)
	os.Chdir("static")
	var wg sync.WaitGroup
	wg.Add(2)
	go createBase(&wg)
	go createCase(&wg)

	wg.Wait()
	os.Chdir("..")
}
func DeleteStatic() {
	os.RemoveAll("static")
}
func createBase(wg *sync.WaitGroup) {
	defer wg.Done()
	f, _ := os.Create("base.html")
	defer f.Close()

	towr := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Test Results</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <h2> Number of failed test cases</h2>
    <p>W/out subtests: {{.GeneralResults.FailedWithoutSub}} of {{.GeneralResults.TotalWithoutSub}} </p>
    <p>Total: {{.GeneralResults.TotalFailed}} of {{.GeneralResults.TotalTests}}</p>
    {{range $tc := .FailedTests}}
    <br>
    {{template "case.html" $tc}}
    {{end}}


<script>
    var coll = document.getElementsByClassName("collapsible");
    var i;
    
    for (i = 0; i < coll.length; i++) {
        coll[i].addEventListener("click", function() {
        var content = this.nextElementSibling;
        if (content.style.display === "block"|| !content.style.display) {
            content.style.display = "none";
        } else {
            content.style.display = "block";
        }
        });
    }
</script>
</body>
</html>
	`
	io.WriteString(f, towr)
}
func createCase(wg *sync.WaitGroup) {
	defer wg.Done()
	f, _ := os.Create("case.html")
	defer f.Close()
	toWr := `<button type="button" class="collapsible"> {{.Name}} </button>
	<div class="content">
		<h4>Started at: {{.RunTimestamp}}</h4>
		<p class="fail"> Result: {{.Result}}</p>
		<p> Output: </p>
		<code>{{range $val := .Outputs}}
			<p>{{$val}}</p>
			{{end}}
		</code>
	</div>`
	io.WriteString(f, toWr)
}
