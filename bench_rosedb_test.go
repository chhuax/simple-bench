package benchmark

import (
	"github.com/flower-corp/rosedb"
	"path/filepath"
	"testing"
)

var roseDB *rosedb.RoseDB

func init() {
	opts := rosedb.DefaultOptions(filepath.Join("bench", "rosedb"))
	opts.IndexMode = rosedb.KeyOnlyMemMode
	var err error
	roseDB, err = rosedb.Open(opts)
	if err != nil {
		panic(err)
	}
	initRoseDBData()
}

func initRoseDBData() {
	for i := 0; i < 500000; i++ {
		err := roseDB.Set(GetKey(i), GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkPutValue_RoseDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := roseDB.Set(GetKey(i), GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetValue_RoseDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := roseDB.Get(GetKey(i))
		if err != nil && err != rosedb.ErrKeyNotFound {
			panic(err)
		}
	}
}
