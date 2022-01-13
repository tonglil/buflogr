package main

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/tonglil/buflogr"
)

type e struct {
	str string
}

func (e e) Error() string {
	return e.str
}

func helper(log logr.Logger, msg string) {
	helper2(log, msg)
}

func helper2(log logr.Logger, msg string) {
	log.WithCallDepth(2).Info(msg)
}

func main() {
	var log logr.Logger = buflogr.New()

	log = log.WithName("MyName")
	example(log.WithValues("module", "example"))

	log.Info("print the log")
	printBuffer(log)
}

// example only depends on logr.
func example(log logr.Logger) {
	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(1).Info("you should see this")
	log.V(1).V(1).Info("you should NOT see this")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	log.Error(e{"an error occurred"}, "goodbye", "code", -1)
	helper(log, "thru a helper")
}

// printBuffer breaks the abstraction to print the logged text in the buffer.
func printBuffer(log logr.Logger) {
	if logSink, ok := log.GetSink().(buflogr.Underlier); ok {
		bl := logSink.GetUnderlying()
		fmt.Print(bl.Buf().String())
	}
}
