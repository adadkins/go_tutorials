package database

import "sync"

type Database struct {
	connection string
}

var (
	dbInstance *Database
	lock       sync.Once
)

func GetDatabase() *Database {
	lock.Do(func() {
		dbInstance = &Database{
			connection: "test_connection",
		}
	})
	return dbInstance
}
