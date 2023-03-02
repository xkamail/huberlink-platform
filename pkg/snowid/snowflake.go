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

func (i ID) Int() int64 {
	return int64(i)
}

const Zero = ID(0)

var _ json.Marshaler = ID(0)

var node, _ = snowflake.NewNode(1)

func Gen() ID {
	return ID(node.Generate().Int64())
}
