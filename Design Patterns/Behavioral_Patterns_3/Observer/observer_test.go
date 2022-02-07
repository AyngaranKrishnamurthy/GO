package observer

import (
	"fmt"
	"testing"
)

//Root test structure to configure the publisher
type TestObserver struct {
	Id      int
	Message string
}

func (p *TestObserver) Notify(m string) { //Implements Observer pattern by defining Notify() method.
	fmt.Printf("Observer: %d, Message Received: %s \n", p.Id, m)
	p.Message = m
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, ""}
	testObserver2 := &TestObserver{2, ""}
	testObserver3 := &TestObserver{3, ""}
	publisher := Publisher{}

	t.Run("AddSubscriber", func(t *testing.T) {
		publisher.AddSubscriber(testObserver1)
		publisher.AddSubscriber(testObserver2)
		publisher.AddSubscriber(testObserver3)
		if len(publisher.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveSubscriber", func(t *testing.T) {
		publisher.RemoveSubscriber(testObserver2)
		if len(publisher.ObserverList) != 3 {
			t.Errorf("The Observer size didn't return the expected value")
		}
		for _, observer := range publisher.ObserverList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}
			if testObserver.Id == 2 {
				t.Fail()
			}
		}
	})

	t.Run("NotifySubscriber", func(t *testing.T) {

		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}
			if printObserver.Message != "" {
				t.Errorf("The Observer's message field wasn't empty")
			}

		}
		if len(publisher.ObserverList) == 0 {
			t.Errorf("Empty Subscriber List")
		}

		message := "Hello World"
		publisher.NotifySubscriber(message)
		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}
			if printObserver.Message != message {
				t.Errorf("Expected message was not received by the observer")
			}
		}
	})
}
