package paypal

import (
	"log"
	"os"
)

func init() {
	SetLogger(log.New(os.Stdout, "", log.LstdFlags|log.Llongfile))
}

type Logger interface {
	SetPrefix(prefix string)
	Prefix() string
	Println(args ...interface{})
	Printf(format string, args ...interface{})
	Output(calldepth int, s string) error
}

var logger Logger

func SetLogger(l Logger) {
	if l == nil {
		l = &nilLogger{}
	}
	if l.Prefix() == "" {
		l.SetPrefix("[paypal] ")
	}
	logger = l
}

type nilLogger struct {
}

func (log *nilLogger) SetPrefix(prefix string) {
}

func (log *nilLogger) Prefix() string {
	return ""
}

func (log *nilLogger) Println(args ...interface{}) {
}

func (log *nilLogger) Printf(format string, args ...interface{}) {
}

func (log *nilLogger) Output(calldepth int, s string) error {
	return nil
}
