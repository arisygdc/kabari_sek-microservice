package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	var conf = &Config{}
	err := LoadConfig("../.", "config", conf)
	assert.NoError(t, err)
	assert.NotEmpty(t, conf.Database.Name)

	assert.Equal(t, "svc-authorization", conf.Service.Name)
}
