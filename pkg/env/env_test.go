package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBoolVal(t *testing.T) {
	t.Run("Get 'true' value from existing env var", func(t *testing.T) {
		_ = os.Setenv("TEST_BOOL_VAL", "true")
		defer os.Setenv("TEST_BOOL_VAL", "")
		assert.True(t, GetBoolVal("TEST_BOOL_VAL", false))
	})
	t.Run("Get 'false' value from existing env var", func(t *testing.T) {
		_ = os.Setenv("TEST_BOOL_VAL", "false")
		defer os.Setenv("TEST_BOOL_VAL", "")
		assert.False(t, GetBoolVal("TEST_BOOL_VAL", true))
	})
	t.Run("Get default value from non-existing env var", func(t *testing.T) {
		_ = os.Setenv("TEST_BOOL_VAL", "")
		assert.True(t, GetBoolVal("TEST_BOOL_VAL", true))
	})
}

func Test_GetStringVal(t *testing.T) {
	t.Run("Get string value from existing env var", func(t *testing.T) {
		_ = os.Setenv("TEST_STRING_VAL", "test")
		defer os.Setenv("TEST_STRING_VAL", "")
		assert.Equal(t, "test", GetStringVal("TEST_STRING_VAL", "invalid"))
	})
	t.Run("Get default value from non-existing env var", func(t *testing.T) {
		_ = os.Setenv("TEST_STRING_VAL", "")
		defer os.Setenv("TEST_STRING_VAL", "")
		assert.Equal(t, "invalid", GetStringVal("TEST_STRING_VAL", "invalid"))
	})
}

func Test_StringsFromEnv(t *testing.T) {
	envKey := "SOMEKEY"
	def := []string{"one", "two"}

	testCases := []struct {
		name     string
		env      string
		expected []string
		def      []string
		sep      string
	}{
		{"List of strings", "one,two,three", []string{"one", "two", "three"}, def, ","},
		{"Comma separated with other delimeter", "one,two,three", []string{"one,two,three"}, def, ";"},
		{"With trimmed white space", "one, two   ,    three", []string{"one", "two", "three"}, def, ","},
		{"Env not set", "", def, def, ","},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envKey, tt.env)
			ss := StringsFromEnv(envKey, tt.def, tt.sep)
			assert.Equal(t, tt.expected, ss)
		})
	}
}
