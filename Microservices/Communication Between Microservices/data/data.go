package data

import "fmt"

// data func
func Data(id string) string {
	var cred string
	cred = "default"
	fmt.Println("Data file processing")
	if id == "1324" {
		cred = "Employee Details" + id + " | Excaliber | Developer | Python | Java | GO "
	} else {
		cred = "Invalid"
	}
	return cred
}
