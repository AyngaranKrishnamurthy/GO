package observer

import "fmt"

type Observer interface {
	Notify(string) //Accepts a string that will contain the message to spread.
}

type Publisher struct {
	ObserverList []Observer //List of subscribed observers
}

func (s *Publisher) AddSubscriber(o Observer) {
	s.ObserverList = append(s.ObserverList, o)
}
func (s *Publisher) RemoveSubscriber(o Observer) {
	var indexToRemove int
	for i, observer := range s.ObserverList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	s.ObserverList = append(s.ObserverList[:indexToRemove], s.ObserverList[indexToRemove+1:]...) //Creates two new slices one till the index to remove and another after index to remove and append both.
}
func (s *Publisher) NotifySubscriber(m string) {
	fmt.Printf("Publisher Received the message")
	for _, observer := range s.ObserverList { //Iterates over the ObserverList to notify fallthrough
		observer.Notify(m)
	}
}
