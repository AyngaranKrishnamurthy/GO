package strategy_2

import "io"

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)    //Adds logger strategy to types, this provides feedback to users
	SetWriter(io.Writer) //Used yo set io.Writer
}

//Both TextSquare and ImageSquare strategiesss must satisfy the setLog and setWriter methods which will simply store some objects on their field.
//Instead of creating the same twice we create a struct that implements them and embed these struct to the strategies. Hint: Like Composite Pattern

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}
