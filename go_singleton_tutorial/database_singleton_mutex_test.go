package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutexDatabase(t *testing.T) {
	t.Run("Mutex Singleton returns same database Object", func(t *testing.T) {
		database1 := GetMutexDatabase()
		database2 := GetMutexDatabase()
		assert.Equal(t, database1, database2)
	})

	t.Run("Modifying the mutex singleton object, also modifies the other object", func(t *testing.T) {
		database1 := GetMutexDatabase()
		database2 := GetMutexDatabase()
		database1.connection = "test"
		assert.Equal(t, database1, database2)
		assert.Equal(t, database1.connection, database2.connection)
	})

	t.Run("Modifying the 2nd mutex singleton object, also modifies the first object", func(t *testing.T) {
		db1 := GetMutexDatabase()
		db2 := GetMutexDatabase()
		db2.connection = "test"
		assert.Equal(t, db1, db2)
		assert.Equal(t, db1.connection, db2.connection)
	})
}
