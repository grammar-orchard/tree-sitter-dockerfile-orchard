package tree_sitter_dockerfile_orchard_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_dockerfile_orchard "codeberg.org/grammar-orchard/tree-sitter-dockerfile-orchard/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dockerfile_orchard.Language())
	if language == nil {
		t.Errorf("Error loading Dockerfile grammar")
	}
}
