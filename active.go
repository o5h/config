package config

import (
	"io"
	"os"
)

var activeConfig *config = &config{}

func Get[T any](name string, defaultValue T) T {
	node := activeConfig.getNode(name)
	if node == nil {
		return defaultValue
	}
	var value T
	if err := node.Decode(&value); err != nil {
		return defaultValue
	}
	return value
}

func Load(path string) error {
	cfg := &config{}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := cfg.read(file); err != nil {
		return err
	}
	activeConfig = cfg
	return nil
}

func Read(r io.Reader) error {
	cfg := &config{}
	if err := cfg.read(r); err != nil {
		return err
	}
	activeConfig = cfg
	return nil
}
