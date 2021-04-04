package custom_types

type TestCase struct {
	RunTimestamp string
	Name         string
	Outputs      []string
	Result       string
}

func NewTestCase(runTimestamp string, name string, outputs []string, result string) *TestCase {
	return &TestCase{RunTimestamp: runTimestamp,
		Name:    name,
		Outputs: outputs,
		Result:  result}
}

type GeneralResults struct {
	FailedTests      map[string]struct{}
	TotalFailed      int
	TotalTests       int
	TotalWithoutSub  int
	FailedWithoutSub int
	TestCases        map[string]*TestCase
}

type ForTemplate struct {
	FailedTests    []*TestCase
	GeneralResults *GeneralResults
}

func NewForTemplate(tc []*TestCase, gr *GeneralResults) *ForTemplate {
	return &ForTemplate{tc, gr}
}
