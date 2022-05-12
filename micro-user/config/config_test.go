package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	var conf = &Config{}
	err := LoadConfig("../.", "config", conf)
	assert.NoError(t, err)
	assert.Equal(t, "user", conf.Database.Name)
}
