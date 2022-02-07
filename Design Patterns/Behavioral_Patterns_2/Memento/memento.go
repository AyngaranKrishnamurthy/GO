package memento

import "fmt"

type State struct {
	Description string
}

type memento struct { //Our states will be containerized into this type before storing them onto the care taker
	state State
}

type originator struct {
	state State
}

func (o *originator) Newmemento() memento { //Takes states from mementos and create new mementos with their stored state
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found")
	}
	return c.mementoList[i], nil
}
