package logger

type Factory interface {
	New(component string) Logger
}
