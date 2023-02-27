package snowid

import (
	"github.com/xkamail/snowflake"
)

var node, _ = snowflake.NewNode(1)

func Gen() int64 {
	return node.Generate().Int64()
}
