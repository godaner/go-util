package randomutil

import (
	"math/rand"
	"fmt"
	"sync"
	"github.com/bwmarrin/snowflake"
)
var (
	node *snowflake.Node
	once sync.Once
)
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}


//获取db  id int64
func GetSnowFlakeId(randomId int64) int64 {
	once.Do(func() {
		getSnowFlake(randomId)
	})
	id := node.Generate()
	return id.Int64()
}


func getSnowFlake(randomId int64) {
	node1, err := snowflake.NewNode(int64(randomId))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		node = node1
	}
}