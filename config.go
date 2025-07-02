package config

import (
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type config struct {
	root *yaml.Node
}

func (cfg *config) read(r io.Reader) error {
	var node yaml.Node
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&node); err != nil {
		return err
	}
	cfg.root = &node
	return nil
}

func (cfg *config) getNode(path string) *yaml.Node {
	if cfg.root == nil {
		return nil
	}
	return findChildByPath(cfg.root.Content[0], path)
}

func findChildByPath(node *yaml.Node, path string) *yaml.Node {
	dotIndex := strings.Index(path, ".")
	if dotIndex == -1 {
		return findChild(node, path)
	}
	childName := path[:dotIndex]
	childNode := findChild(node, childName)
	if childNode == nil {
		return nil
	}
	remainingPath := path[dotIndex+1:]
	if remainingPath == "" {
		return childNode
	}
	return findChildByPath(childNode, remainingPath)
}

func findChild(node *yaml.Node, name string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i < len(node.Content); i += 2 {
		if node.Content[i].Value == name {
			return node.Content[i+1]
		}
	}
	return nil

}

func (cfg *config) is(name string, value bool) bool {
	node := cfg.getNode(name)
	if node == nil {
		return value
	}

	if node.Kind != yaml.ScalarNode {
		return value
	}
	return node.Value == "true"
}
