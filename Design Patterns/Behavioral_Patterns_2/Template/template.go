package template

import "strings"

//We need two interfaces, one for the user data and other to send pre-existing(template)

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string //Executes the algorthim by adding the user data to the template
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}
func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ") //joins the array of strings with the second item between them.
}
