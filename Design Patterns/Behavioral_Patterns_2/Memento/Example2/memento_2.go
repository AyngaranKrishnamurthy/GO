package main

import "fmt"

//Example using Command Pattern and Facade
/*The idea is to use Command pattern to encapsulate a set of different states and provide a small facade
to automate the insertion in the careTaker object*/
//An example of an Audio Mixer is implemented here
/*
We will use memento pattern to save two types of states: Volume and Mute
Volume is state of byte type and Mute is of boolean type
*/

type Command interface {
	GetValue() interface{} //Returns an interface to a a value (represents any-type can be type-casted later)
}

type Volume byte //Volume state
func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool //Mute state
func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type Originator struct {
	command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{memento: o.command}
}
func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.command = m.memento
}

type careTaker struct { //Here we use stack instead of a list
	mementoList []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}
func (c *careTaker) Pop() Memento {
	if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0 : len(c.mementoList)-1]
		return tempMemento
	}
	return Memento{}
}

//The Facade pattern will hold contents of the originator and provide two easy to use methods: save and restore
type MementoFacade struct {
	originator Originator
	careTaker  careTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.command = s
	m.careTaker.Add(m.originator.NewMemento())
}
func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.command
}

func main() {
	m := MementoFacade{} //Gets a variable with facade pattern

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}

//Type Assertion - Checks the type and prints appropriate value
func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
