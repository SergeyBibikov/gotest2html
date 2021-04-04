# gotest2html
A utility to create a minimalistic html report for failed test cases from the `go test -json` output

### Build
1. Clone the repo
2. Run `go build`
3. (Optional) For more convinient usage I'd recommend adding the path to the binary to the PATH variable

### Usage
* You may generate report from a file:\
`$ go test -json > results`\
`$ gotest2html results`
* Or on Linux (and Mac, I guess) pipe the output from `go test -json` directly:\
`$ go test -json | gotest2html`
