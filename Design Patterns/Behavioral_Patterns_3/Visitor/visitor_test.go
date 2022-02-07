package visitor

import "testing"

//Creating a type that implements io.Writer package so that we can use it to write tests.
type TestHelper struct { //Implements io.Writer interface
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World", //Giving MessageA struct a value "Hello World"
			Output: testHelper,
		}
		msg.Accept(visitor) //In this method the VisitA(*MessageA) is executed to alter contents of message field
		msg.Print()         //MessageA struct must write contents of Msg to the provided io.Writer interface(testHelper)

		expected := "A: Hello World (Visited A)" // Checks if message is modified as expected
		if testHelper.Received != expected {
			t.Errorf("Expected result was not returned")
		}
	})

	t.Run("MessageB test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World", //Giving MessageA struct a value "Hello World"
			Output: testHelper,
		}
		msg.Accept(visitor) //In this method the VisitA(*MessageA) is executed to alter contents of message field
		msg.Print()         //MessageA struct must write contents of Msg to the provided io.Writer interface(testHelper)

		expected := "B: Hello World (Visited byte)" // Checks if message is modified as expected
		if testHelper.Received != expected {
			t.Errorf("Expected result was not returned")
		}
	})
}
