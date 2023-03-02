package snowid

import (
	"encoding/json"
	"fmt"

	"github.com/xkamail/snowflake"
)

type ID int64

func (i ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(i))
}

const Zero = ID(0)

var _ json.Marshaler = ID(0)

var node, _ = snowflake.NewNode(1)

func Gen() int64 {
	return node.Generate().Int64()
}
