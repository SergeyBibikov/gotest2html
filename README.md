# gotest2html
A utility to create a minimalistic html report for **failed** test cases from the `go test -json` output

### Build
1. Clone the repo
2. Run `go build`
3. (Optional) For more convinient usage I'd recommend adding the path to the binary to the PATH variable

### Usage
* You may generate the report from a file:\
`$ go test -json > results`\
`$ gotest2html results`
* Or on Linux (and Mac, I guess) pipe the output from `go test -json` directly:\
`$ go test -json | gotest2html`
Once done, the folder "reports" will be created in the current directory. The folder will contain two files: "report_failed.html" and "styles.css".\
Opened in a browser, the report will look like this:
![sample_report](https://user-images.githubusercontent.com/53792559/113504938-55849880-9544-11eb-8231-3673f020fe71.png)

