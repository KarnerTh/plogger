package presentation

type logList []string

// PERF: use fixed array?
func (s *logList) Push(value string) {
	*s = append(*s, value)
	if len(*s) > 5 {
		*s = (*s)[1:len(*s)]
	}
}
