package interfaces

type ApplicationEntry struct {
	Version  float32
	Type     string
	Name     string
	Comment  string
	TryExec  string
	Exec     string
	Icon     string
	MimeType string
	Actions  []*Action
}

type Action struct {
	Name string
	Exec string
	Icon string
}
