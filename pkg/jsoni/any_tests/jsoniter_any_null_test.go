package any_tests

import (
	"testing"

	"github.com/bingoohuang/gg/pkg/jsoni"
	"github.com/stretchr/testify/require"
)

func Test_read_null_as_any(t *testing.T) {
	should := require.New(t)
	any := jsoni.Get([]byte(`null`))
	should.Equal(0, any.ToInt())
	should.Equal(float64(0), any.ToFloat64())
	should.Equal("", any.ToString())
	should.False(any.ToBool())
}
