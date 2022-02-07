package shapes

import "std/go_workspace/Design_Patterns/Behavioral_1/Strategy/strategy_2"

type TextSquare struct {
	strategy_2.PrintOutput
}

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
