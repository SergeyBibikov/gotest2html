package conversions

import "github.com/SergeyBibikov/gotest2html/custom_types"

func GetGeneralResults(m *[]map[string]interface{}) *custom_types.GeneralResults {
	var failedTests map[string]struct{} = make(map[string]struct{})
	var testCases map[string]*custom_types.TestCase = make(map[string]*custom_types.TestCase)
	totalFailed := 0
	totalTests := 0
	var runTimestamp string
	var name string
	var outputs []string
	var result string
	var allowOutput bool
	runCounter := 0
loop:
	for _, v := range *m {
		action := v["Action"].(string)
		switch action {
		case "run":
			totalTests++
			runCounter++
			if _, ok := v["Time"]; ok {
				runTimestamp = v["Time"].(string)
			} else {
				runTimestamp = "----"
			}
			if name == "" {
				name = v["Test"].(string)
			}
			allowOutput = true
		case "output":
			if !allowOutput {
				break loop
			}
			outputs = append(outputs, v["Output"].(string))
		case "fail":
			failedTests[name] = struct{}{}
			totalFailed++
			result = action
			fallthrough
		case "pass":
			runCounter--
			if result == "" {
				result = action
			}
			if runCounter == 0 {
				testCases[name] = custom_types.NewTestCase(runTimestamp, name, outputs, result)
				runTimestamp = ""
				name = ""
				outputs = make([]string, 0)
				result = ""
				allowOutput = false
			}
		}
	}
	return &custom_types.GeneralResults{
		TotalFailed:      totalFailed,
		TotalTests:       totalTests,
		TotalWithoutSub:  len(testCases),
		FailedWithoutSub: len(failedTests),
		FailedTests:      failedTests,
		TestCases:        testCases}
}

func GetFailedTestsSlice(gr *custom_types.GeneralResults) []*custom_types.TestCase {
	var result []*custom_types.TestCase
	for v := range gr.FailedTests {
		result = append(result, gr.TestCases[v])
	}
	return result
}
