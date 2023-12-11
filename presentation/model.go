package presentation

type model struct {
	logLines logList
	values   []float64
}

var InitialModel = model{
	logLines: logList{},
	values:   []float64{},
}
