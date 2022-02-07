package chainofresponsibility

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string) //Takes the message we want to log and passes it to the following chain in the link.
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First Logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second Logger: %s\n", s)
		if se.NextChain != nil {
			se.NextChain.Next(s)
		}
		return
	}
	fmt.Println("Finishing Second Logging \n ")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("Writer Logger: " + s))
	}
	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

//When implementing the chain we usually start with the last piece on the link and in our case it is a WriterLogger.
//The WriterLogger writes to an io.Writer so we can pass myWriter as io.Writer interface.

//Adding a closure function to the chain. A more flexible link in the chain for quick debugging.
type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string) //It is a function that takes string and returns nothing.
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
