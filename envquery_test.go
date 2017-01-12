package envquery

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setTestingEnvs() {
	datas := map[string]string{
		"TEST1": "test",
		"TEST2": "=test=",
	}

	for k, v := range datas {
		os.Setenv(k, v)
	}
}

func TestParseVars(t *testing.T) {
	setTestingEnvs()
	result := parseVars()

	assert.Equal(t, "test", (*result)["TEST1"], "Must extract and parse environment variables")
	assert.Contains(t, "=test=", (*result)["TEST2"], "Must extract and parse environment variables")
}

func TestNewEnQuery(t *testing.T) {
	result := NewEnvQuery()

	assert.Equal(t, "test", (*result.envs)["TEST1"], "Must extract and parse environment variables")
	assert.Contains(t, "=test=", (*result.envs)["TEST2"], "Must extract and parse environment variables")
}

func TestGetAllKeys(t *testing.T) {
	setTestingEnvs()

	q := NewEnvQuery()

	keys := q.GetAllKeys()

	results := []string{}

	for _, k := range keys {
		if k == "TEST1" || k == "TEST2" {
			results = append(results, k)
		}
	}

	assert.Len(t, results, 2, "Must contains 2 elements")
}

func TestFindEntries(t *testing.T) {
	setTestingEnvs()

	q := NewEnvQuery()

	keys, err := q.FindEntries(".*?1")

	assert.NoError(t, err, "Must return no errors")
	assert.Len(t, keys, 1, "Must contains 1 elements")
	assert.Equal(t, "test", keys["TEST1"], "Must have env key and value")

	_, err = q.FindEntries("?")

	assert.EqualError(t, err, "error parsing regexp: missing argument to repetition operator: `?`", "Must return an error when regexp is unvalid")
}