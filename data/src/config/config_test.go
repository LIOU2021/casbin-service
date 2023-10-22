package config

import (
	"casbin-service/logger"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestMain(t *testing.M) {
	logger.Init()
	Init()
	code := t.Run()
	defer func() {
		os.Exit(code)
	}()
}

func Test_Config(t *testing.T) {
	out, err := yaml.Marshal(Config)
	assert.NoError(t, err)
	assert.NotEmpty(t, string(out))
	// t.Log(string(out))
}
