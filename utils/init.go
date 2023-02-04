package utils

import "github.com/bwmarrin/snowflake"

func init() {
	var err error
	snowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}
