package benchmark

import (
	"github.com/recoilme/pudge"
	"log"
	"testing"
)

var (
	pudgeDB *pudge.Db
)

func init() {
	dir := "bench/pudge"
	pudgeDB, err = pudge.Open(dir, pudge.DefaultConfig)
}

func initPudgeData() {
	for i := 0; i < 10000; i++ {
		key := GetKey(i)
		val := GetValue()
		err := pudgeDB.Set(key, val)
		if err != nil {
			log.Fatal("pudge write data err.", err)
		}
	}
}

func BenchmarkPutValue_Pudge(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := GetKey(i)
		val := GetValue()
		err := pudgeDB.Set(key, val)
		if err != nil {
			log.Fatal("pudge write data err.", err)
		}
	}
}

func BenchmarkGetValue_Pudge(b *testing.B) {
	initPudgeData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := GetKey(i)
		var b []byte
		err = pudgeDB.Get(key, &b)
		if err != nil && err != pudge.ErrKeyNotFound {
			log.Fatal("pudge read data err.", err)
		}
	}
}
