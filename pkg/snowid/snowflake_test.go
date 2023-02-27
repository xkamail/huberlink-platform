package snowid

import (
	"testing"
)

func TestGen(t *testing.T) {
	t.Run("Gen", func(t *testing.T) {
		a := Gen()
		t.Log(a)
	})
}
