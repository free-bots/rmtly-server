package interfaces

type ApplicationEntry struct {
	version   float32
	entryType string // type
	name      string
	comment   string
	tryExec   string
	exec      string
	icon      string
	mimeType  string
	actions   []Action
}

type Action struct {
	name string
	exec string
	icon string
}
