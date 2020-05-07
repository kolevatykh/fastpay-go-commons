package logger

type EmptyLogger struct{}

func (logger *EmptyLogger) Error(...interface{}) {}

func (logger *EmptyLogger) Info(...interface{}) {}

func (logger *EmptyLogger) Debug(...interface{}) {}
