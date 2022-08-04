package log_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/lthibault/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	logger := log.New(
		log.WithWriter(&buf),
		log.WithFormatter(&logrus.JSONFormatter{}))

	logger.With(log.F{
		"string":  "test",
		"int":     42,
		"pointer": nil,
	}).Info("test message")

	var out output
	err := json.NewDecoder(&buf).Decode(&out)
	require.NoError(t, err, "should decode json")

	assert.Equal(t, "test", out.String)
	assert.Equal(t, 42, out.Int)
	assert.Nil(t, out.Pointer)
	assert.Equal(t, "test message", out.Message)
}

type output struct {
	String  string    `json:"string"`
	Int     int       `json:"int"`
	Pointer *struct{} `json:"pointer"`
	Message string    `json:"msg"`
}
