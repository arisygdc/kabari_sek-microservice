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
	assert.NotEmpty(t, conf.Database.User)
	assert.NotEmpty(t, conf.Database.Host)
	assert.NotEmpty(t, conf.Database.Pass)
	assert.NotEmpty(t, conf.Database.Port)

	assert.NotEmpty(t, conf.Token.Dur)
	assert.NotEmpty(t, conf.Token.Secret)

	assert.NotEmpty(t, conf.Server.SvcName)
}
