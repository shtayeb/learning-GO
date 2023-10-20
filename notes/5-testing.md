## Testing
- Vanilla GO includes testing
- A test is a file with suffix_test.go
- You define functions with prefix Test and with an special signature receiving a *testing.T argument
- You can create subtests as goroutines
- You use CLI with `go test`
- TableDrivenTest Design Pattern
- **Fuzzing**
    Automated testing that manipulates inputs to find bugs. Go fuzzing uses converage guaidance to find failures and is valuable in detection security exploits and vulnerabilities.