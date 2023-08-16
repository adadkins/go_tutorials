package database

import "sync"

type MutexDatabase struct {
	connection string
}

var (
	mutexDb *MutexDatabase
	dbMutex = sync.Mutex{}
)

func GetMutexDatabase() *MutexDatabase {
	dbMutex.Lock()

	defer dbMutex.Unlock()
	if mutexDb == nil {
		mutexDb = &MutexDatabase{
			connection: "test_connection",
		}
	}

	return mutexDb
}
