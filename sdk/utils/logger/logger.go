package logger

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
	PrintInfo(str)
}

func (logger *Logger) Warn(str string) {
	if logger.Enable == false {
		return
	}
	PrintWarn(str)
}

func (logger *Logger) Error(str string) {
	if logger.Enable == false {
		return
	}
	PrintError(str)
}
