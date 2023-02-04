package utils

import "github.com/bwmarrin/snowflake"

var snowflakeNode *snowflake.Node

func GenUUID() snowflake.ID {
	return snowflakeNode.Generate()
}
