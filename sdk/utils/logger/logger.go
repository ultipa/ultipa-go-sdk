package logger

import "ultipa-go-sdk/sdk/printers"

type Logger struct {
	Enable bool
}

func NewLogger(enable bool) *Logger {
	return &Logger{
		Enable: enable,
	}
}

func (logger *Logger) Log(str string) {
	if logger.Enable == false {
		return
	}
	printers.PrintInfo(str)
}

func (logger *Logger) Warn(str string) {
	if logger.Enable == false {
		return
	}
	printers.PrintWarn(str)
}

func (logger *Logger) Error(str string) {
	if logger.Enable == false {
		return
	}
	printers.PrintError(str)
}
