package template

func init() {
	_templates["/.gitignore"] = genGitIgnore()
}

func genGitIgnore() string {
	return `
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
__debug_bin

.vscode/*
`
}