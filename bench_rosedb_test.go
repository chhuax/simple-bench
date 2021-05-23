package benchmark

import (
	"github.com/roseduan/rosedb"
	"log"
	"testing"
)

var (
	roseDBKeyValueRam *rosedb.RoseDB
	roseDBKeyOnlyRam  *rosedb.RoseDB
)

func init() {
	initKeyValueRam()
	initKeyOnlyRam()
}

func initKeyValueRam() {
	config := rosedb.DefaultConfig()
	config.DirPath = "bench/rosedb1"
	roseDBKeyValueRam, err = rosedb.Open(config)
	if err != nil {
		log.Fatal("open rosedb err.", err)
	}
}

func initKeyOnlyRam() {
	config := rosedb.DefaultConfig()
	config.IdxMode = rosedb.KeyOnlyRamMode
	config.DirPath = "bench/rosedb2"
	roseDBKeyOnlyRam, err = rosedb.Open(config)
	if err != nil {
		log.Fatal("open rosedb err.", err)
	}
}

func initRosedbData(mode int) {
	for i := 0; i < 10000; i++ {
		key := GetKey(i)
		val := GetValue()
		if mode == int(rosedb.KeyOnlyRamMode) {
			err = roseDBKeyOnlyRam.Set(key, val)
		} else {
			err = roseDBKeyValueRam.Set(key, val)
		}
		if err != nil {
			log.Fatal("rosedb write data err.", err)
		}
	}
}

func BenchmarkPutValue_RoseDB_KeyValRam(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue()
		err := roseDBKeyValueRam.Set(key, val)
		if err != nil {
			log.Fatal("rosedb write data err.", err)
		}
	}
}

func BenchmarkGetValue_RoseDB_KeyValRam(b *testing.B) {
	initRosedbData(0)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = roseDBKeyValueRam.Get(GetKey(i))
		if err != nil && err != rosedb.ErrKeyNotExist {
			log.Fatal("rosedb get data err.", err)
		}
	}
}

func BenchmarkPutValue_RoseDB_KeyOnlyRam(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue()
		err := roseDBKeyOnlyRam.Set(key, val)
		if err != nil {
			log.Fatal("rosedb write data err.", err)
		}
	}
}

func BenchmarkGetValue_RoseDB_KeyOnlyRam(b *testing.B) {
	initRosedbData(1)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = roseDBKeyOnlyRam.Get(GetKey(i))
		if err != nil && err != rosedb.ErrKeyNotExist {
			log.Fatal("rosedb get data err.", err)
		}
	}
}
