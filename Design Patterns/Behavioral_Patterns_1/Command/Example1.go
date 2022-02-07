//There are a total of 3 different examples each handles differently

package main

import "fmt"

// Example1 -> Command Handler - Executes contents of the command

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

//Command interface that accepts a message string and returns the command interface
func CreateCommand(s string) Command {
	fmt.Println("Creating Command")
	return &ConsoleOutput{
		message: s,
	}
}

//Command Queue to store in queue any type implementing the command interface
type CommandQueue struct {
	queue []Command //stores an array of command interface
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 { //When queue reaches 3 items it executes all the command stored in the queue field.
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}
func main() {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First Message"))
	queue.AddCommand(CreateCommand("Second Message"))
	queue.AddCommand(CreateCommand("Third Message"))

	queue.AddCommand(CreateCommand("Fourth Message"))
	queue.AddCommand(CreateCommand("Fifth Message"))
}

/*
Example 2: Delegation(passing) of information to a different object
//Instead of printing to console we create a command that extracts information.

type Command interface{
	Info() string
}
type TimePassed struct{
	start time.Time
}
func (t *TimePassed) Info() string{
	return time.Since(t.start).String() //Returns time elapsed since the start
}

type HelloMessage struct{}
func (h HelloMessage) Info() string{
	return "Hello World"
}

func main() {
	var timeCommand Command
	timeCommand = &TimePassed{time.Now()}

	var helloCommand Command
	helloCommand = &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}

*/

/*
Example 3: Chain of Responsibility of Commands
//Log the link information retrieved to the console
type Command interface{
Info() string
}
type TimePassed struct{
start time.Time
}
func (t *TimePassed) Info() string{
return time.Since(t.start).String() //Returns time elapsed since the start
}

type ChainLogger interface{
Next(Command)
}

type Logger struct{
NextChain ChainLogger
}
func (f* Logger) Next(c Command){
time.Sleep(time.Second)
fmt.Println("Elapsed time from creation: %s\n",c.Info())
if f.NextChain!= nil{
f.NextChain.Next(c)
}
}

func main(){
second := new{Logger}
first := Logger{NextChain: second}
command := &TimePassed{start: time.Now()}
first.Next(command)
*/
