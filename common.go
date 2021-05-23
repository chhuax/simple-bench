package benchmark

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetKey(n int) []byte {
	return []byte("test_key_" + fmt.Sprintf("%09d", n))
}

func GetValue() []byte {
	return []byte("test_val-val-val-val-val-val-val-val-val-val-val-val-" + strconv.FormatInt(rand.Int63(), 10))
}
