package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("Singleton returns same database Object", func(t *testing.T) {
		database1 := GetDatabase()
		database2 := GetDatabase()
		assert.Equal(t, database1, database2)
	})

	t.Run("Modifying the singleton object, also modifies the other object", func(t *testing.T) {
		database1 := GetDatabase()
		database2 := GetDatabase()
		database1.connection = "test"
		assert.Equal(t, database1, database2)
		assert.Equal(t, database1.connection, database2.connection)
	})

	t.Run("Modifying the 2nd singleton object, also modifies the first object", func(t *testing.T) {
		db1 := GetDatabase()
		db2 := GetDatabase()
		db2.connection = "test"
		assert.Equal(t, db1, db2)
		assert.Equal(t, db1.connection, db2.connection)
	})
}
