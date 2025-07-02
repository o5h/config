package config_test

import (
	"testing"

	"github.com/o5h/config"
	"github.com/o5h/testing/assert"
)

func TestIs(t *testing.T) {
	assert.Nil(t, config.Load("testdata/config.yaml"))
	assert.Eq(t, config.Get("app.name", "default"), "MyApp")
	assert.Eq(t, config.Get("app.database.host", "default"), "localhost")
	assert.Eq(t, config.Get("app.features.enable_feature_y", true), false)
	assert.Eq(t, config.Get("app.settings.max_connections", 1), 100)
	assert.EqSlice(t, config.Get("app.values", []string{""}), []string{"a", "b", "c", "d"})
	mapping := config.Get("app.mapping", map[string]string{})
	assert.Eq(t, mapping["key1"], "value1")
	assert.Eq(t, mapping["key2"], "value2")
}
