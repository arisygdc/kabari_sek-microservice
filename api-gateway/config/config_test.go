package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := Config{}
	assert.Empty(t, cfg)
	err := LoadConfig("../", "config", &cfg)
	assert.NoError(t, err)

	assert.NotEmpty(t, cfg.Endpoint.User)
}
