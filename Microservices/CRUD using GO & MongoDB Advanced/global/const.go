package global

const (
	dburi       = "mongodb+srv:/<username>:<password>@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbname      = "Employee"
	performance = 100
)

var (
	jwtSecret = []byte("crudSecret")
)
