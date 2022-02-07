package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface { //the interface for all the decorating types
	AddIngredient() (string, error)
}

type PizzaDecorator struct { //Type imlementing IngredientsAdd
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct { //Meat struct & Implementation
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the meat")
	}
	s, err := m.Ingredient.AddIngredient() //Check for pointer to inner AddIngredient() first and makes sure there are no errors
	if err != nil {
		return " ", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type Onion struct { //Onion struct and Implementation
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the onion")
	}
	s, err := o.Ingredient.AddIngredient() //Check for pointer to inner AddIngredient() first and makes sure there are no errors
	if err != nil {
		return " ", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
