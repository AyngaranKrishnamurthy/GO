package global

const (
	dburi       = "mongodb+srv://standard:example@cluster0.ut86q.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbname      = "crud_service"
	performance = 100
)

var (
	jwtSecret = []byte("crudSecret")
)
