package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	//To check if the careTaker list has the expected state added
	careTaker := careTaker{}
	mem := originator.Newmemento()
	if mem.state.Description != "Idle" {
		t.Error("Expected state was not found")
	}

	//To check if the careTaker list has increased in size
	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)

	if len(careTaker.mementoList) != currentLen+1 {
		t.Errorf("No new elements was added to the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}

	originator.state = State{"Idle"} //Sending first object to list
	careTaker.Add(originator.Newmemento())

	mem, err := careTaker.memento(0) //Asking list for the first object stored
	if err != nil {
		t.Fatal(err)
	}
	if mem.state.Description != "Idle" {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.memento(-1) //Checking for negative index value from the list
	if err == nil {
		t.Fatal("An error is expected at this point of code")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) { //Must take a memento and extract all its state info to the originator object
	originator := originator{state: State{"Idle"}} //Creating a originator variable with an idle state
	idleMemento := originator.Newmemento()         //Retrieve a newmemento object to use it later

	originator.ExtractAndStoreState(idleMemento) //Storing the state of the originator to the idleMemento's state value.
	if originator.state.Description != "Idle" {
		t.Error("Unexpected State")
	}
}
