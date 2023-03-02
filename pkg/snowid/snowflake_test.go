package snowid

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	t.Parallel()
	t.Run("Gen", func(t *testing.T) {
		a := Gen()
		t.Log(a)
	})
	t.Run("json string unmarshal", func(t *testing.T) {
		ran := Gen()
		var a ID
		err := json.Unmarshal([]byte(fmt.Sprintf(`"%v"`, ran)), &a)
		assert.NoError(t, err)
		assert.Equal(t, ran, a)
	})
}
