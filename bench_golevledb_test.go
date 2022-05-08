package benchmark

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"testing"
)

var (
	levelDb *leveldb.DB
)

func init() {
	dir := "bench/leveldb"
	var err error
	levelDb, err = leveldb.OpenFile(dir, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initLevelDBValue() {
	for i := 0; i < 500000; i++ {
		key := GetKey(i)
		val := GetValue()
		err := levelDb.Put(key, val, nil)
		if err != nil {
			log.Fatal("leveldb write data err.", err)
		}
	}
}

func BenchmarkPutValue_GoLevelDB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := GetKey(i)
		val := GetValue()
		err := levelDb.Put(key, val, nil)
		if err != nil {
			log.Fatal("leveldb write data err.", err)
		}
	}
}

func BenchmarkGetValue_GoLevelDB(b *testing.B) {
	initLevelDBValue()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := levelDb.Get(GetKey(i), nil)
		if err != nil && err != leveldb.ErrNotFound {
			log.Fatal("leveldb read data err.", err)
		}
	}
}
