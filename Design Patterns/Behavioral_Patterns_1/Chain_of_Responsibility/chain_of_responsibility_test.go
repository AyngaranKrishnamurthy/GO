package chainofresponsibility

import (
	"fmt"
	"strings"
	"testing"
)

//Implementing io.Writer struct to use in our testing
type myTestWriter struct {
	receivedMessage *string
}

//Wwe will pass an instance of the myTestWriter struct to WriterLogger so we can track what's being logged on testing.
func (m *myTestWriter) Write(p []byte) (int, error) {
	if m.receivedMessage == nil {
		m.receivedMessage = new(string)
	}
	tempMessage := fmt.Sprintf("%s%s", *m.receivedMessage, p)
	m.receivedMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s)) //byte(s) represents byte of string
}

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it is found the word 'hello', third writes to some variable if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain \n")

			if myWriter.receivedMessage != nil {
				t.Fatal("Last link should not receive any message")
			}

			chain.Next("Hello\n")
			if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
				t.Fatal("Last link did not receive expected message")
			}
		})

	t.Run("2 loggers, second uses the closure implementation", func(t *testing.T) {
		myWriter = myTestWriter{}
		closureLogger := ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				myWriter.receivedMessage = &s
			},
		}

		writerLogger.NextChain = &closureLogger

		chain.Next("Hello closure logger")

		if *myWriter.receivedMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't received in myWriter")
		}
	})
}
