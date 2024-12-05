package gogger

import (
	"fmt"
	"time"
)

func (l *Logger) print(mtype string, message string, err ...error) {
	var typeColor Color

	switch mtype {
	case "Info":
		typeColor = l.Config.Colors.Cyan

	case "Error":
		typeColor = l.Config.Colors.Red

	case "Warning":
		typeColor = l.Config.Colors.Yellow

	case "Ok":
		typeColor = l.Config.Colors.Green
	}

	if err != nil {
		fmt.Printf("%s%s%s | %s%s:%s %s%s%s -> %s%s%s\n",
			l.Config.Colors.Gray,
			time.Now().Format("2022-12-04 10:41 PM"),
			l.Config.Colors.Reset,
			/////////////////////
			typeColor,
			mtype,
			l.Config.Colors.Reset,
			/////////////////////
			l.Config.Colors.White,
			message,
			l.Config.Colors.Reset,
			/////////////////////
			l.Config.Colors.Red,
			err,
			l.Config.Colors.Reset,
			/////////////////////
		)
	} else {
		fmt.Printf("%s%s%s | %s%s:%s %s%s%s\n",
			l.Config.Colors.Gray,
			time.Now().Format("2022-12-04 10:41 PM"),
			l.Config.Colors.Reset,
			/////////////////////
			typeColor,
			mtype,
			l.Config.Colors.Reset,
			/////////////////////
			l.Config.Colors.White,
			message,
			l.Config.Colors.Reset,
			/////////////////////
		)
	}
}

func (l *Logger) Info(message string) {
	l.print("Info", message)
}

func (l *Logger) Err(err error, message string) {
	if err == nil {
		return
	}
	l.print("Error", message, err)
}

func (l *Logger) Warn(message string) {
	l.print("Warning", message)
}

func (l *Logger) Ok(message string) {
	l.print("Ok", message)
}

func NewLogger() *Logger {
	logger := Logger{
		Config: Config{
			Time: true,
			Bold: true,
			Colors: Colors{
				Yellow: "\033[33m",
				Red:    "\033[31m",
				Green:  "\033[32m",
				Cyan:   "\033[36m",
				Gray:   "\033[37m",
				White:  "\033[97m",
				Reset:  "\033[0m",
			},
		},
	}

	return &logger
}
