package interpreter

import "testing"

func TestCalculate(t *testing.T) {
	tempOperation := "3 4 sum 2 sub"
	res, err := Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}
	if res != 5 {
		t.Error("Expected result 5 was not returned")
	}

	tempOperation = "5 3 sub 8 mul 4 sum 5 div"
	res, err = Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}
	if res != 4 {
		t.Error("Expected result 4 was not returned")
	}

}
