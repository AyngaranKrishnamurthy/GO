package visitor

import (
	"fmt"
	"io"
	"os"
)

type MessageA struct {
	Msg    string //Message it will print
	Output io.Writer
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

type Visitor interface { //Visitor interface has visit method for Visitable interface's MessageA and MessageB type
	VisitA(*MessageA)
	VisitB(*MessageB)
}

//Message structs implementing Accept(visitor) method
func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

//Implementations of methods declared in Visitor interface
func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

//Print() method for each message type to test the contents of the Msg field on each type
func (m *MessageA) Print() {
	if m.Output == nil { // Checks if the output field is null
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg) //Takes io.Writer package as first argument and the text format as second argument
}
func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

type Visitable interface {
	Accept(Visitor) //This method will execute the decoupled algorithm
}
type MessageVisitor struct{}
