package storage

type config struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	SSLMode  string
}

// NewConfig returns a new config object
// Todo: Read from yaml file or env variables
func NewConfig() *config {
	return &config{
		host:     "localhost",
		port:     "5432",
		user:     "postgres",
		password: "postgres",
		dbName:   "basic-api",
		SSLMode:  "disable",
	}
}
