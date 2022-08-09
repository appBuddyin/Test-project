package ifelseStatment

func gettingPassingMarks(marks int) string {
	if marks <= 33 {
		return "fail"
	} else {
		return "pass"
	}
}
