package h3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistanceBetweenH3Indices(t *testing.T) {
	req := require.New(t)

	test := func(firstH3, secondH3 int64, want string) func(t *testing.T) {
		return func(t *testing.T) {
			actualDistance := DistanceBetweenH3Indices(firstH3, secondH3)
			req.Equal(want, fmt.Sprintf("%0.3f", actualDistance))
		}
	}

	t.Run("with real h3", test(0x871106028ffffff, 0x87110600dffffff, "4227.901"))
	t.Run("with invalid h3 args", test(0, 0, "0.000"))
}
