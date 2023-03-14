package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoLevelDB_Database(t *testing.T) {

	dir, err := ioutil.TempDir("", "goleveldb")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	testDB, _ := openDatabase(GoLevelDBBackend, "test", dir)
	defer testDB.Close()

	key := []byte("hello")
	value := []byte("world")

	bucket, _ := testDB.GetBucket("hello")
	bucket.Set(key, value)
	result, _ := bucket.Get(key)
	assert.Equal(t, value, result, "equal")
	assert.True(t, bucket.Has(key), "True")

	bucket.Delete(key)
	result, _ = bucket.Get(key)
	assert.Nil(t, result, "empty")
}
