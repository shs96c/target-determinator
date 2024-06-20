//go:build tools

package tools

import (
	_ "github.com/aristanetworks/goarista/cmd/importsort"
	_ "github.com/bazelbuild/bazel-gazelle/label"
	_ "golang.org/x/tools/imports"
)
