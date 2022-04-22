package movies_test

import (
	"testing"

	"mainmodule/internal/apps/movies/services/movies"

	"github.com/stretchr/testify/assert"
)

func TestDoTest(t *testing.T) {
	m := &movies.MovieService{Name: "test", TableName: "test"}
	assert.Equal(t, m.DoTest(), []string{"hello", "movies"})
}
