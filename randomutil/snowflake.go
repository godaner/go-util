package randomutil

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

//diff randomId will produce never same snow flake id
func GetSnowFlakeId() int64 {
	once.Do(func() {
		getSnowFlake()
	})
	id := node.Generate()
	return id.Int64()
}

func GetSnowFlakeIdStr() string {

	return fmt.Sprintf("%d", GetSnowFlakeId())
}

func getSnowFlake() {
	node1, err := snowflake.NewNode(NewUUID().Int64())
	if err != nil {
		fmt.Println(err)
		return
	} else {
		node = node1
	}
}