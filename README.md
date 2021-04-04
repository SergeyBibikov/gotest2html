# gotest2html
A utility to create a minimalistic html report for failed test cases from the `go test -json` output

### Usage
* You may generate report from a file:\
`$ go test -json > results`\
`$ gotest2html results`
* Or on Linux (and Mac, I guess) pipe the output from `go test -json` directly:\
`$ go test -json | gotest2html`
