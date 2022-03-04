package buflogr_test

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

func ExampleNew() {
	var log logr.Logger = buflogr.New()
	log = log.WithName("MyName")
	log = log.WithValues("module", "example")

	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(1).Info("you should see this")
	log.V(1).V(1).Info("you will also see this")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	log.Error(e{"an error occurred"}, "goodbye", "code", -1)
	helper(log, "thru a helper")

	printBuffer(log)

	// Output:
	// INFO MyName hello module example val1 1 val2 map[k:1]
	// V[1] MyName you should see this module example
	// V[2] MyName you will also see this module example
	// ERROR <nil> MyName uh oh module example trouble true reasons [0.1 0.11 3.14]
	// ERROR an error occurred MyName goodbye module example code -1
	// INFO MyName thru a helper module example
}

// printBuffer breaks the abstraction to print the logged text in the buffer.
func printBuffer(log logr.Logger) {
	if logSink, ok := log.GetSink().(buflogr.Underlier); ok {
		bl := logSink.GetUnderlying()
		fmt.Print(bl.Buf().String())
	}
}
