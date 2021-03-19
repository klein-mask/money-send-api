package infrastructure

type PostgresDSN struct {
	host string
	user string
	password string
	dbname string
	port string
	sslmode string
}
