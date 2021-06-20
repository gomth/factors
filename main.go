package factors

// Generate - if you dont know what to call - call this
// error != nil in case of error. Returned array []int is undefined in this case.
func Get(n int) ([]int, error) {
	return GetUsingTrialDivisionForInt(n)
}
