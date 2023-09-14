package config

import (
	"path/filepath"
	"runtime"
)

var (
	// GET CURRENT FILE FULL PATH FROM RUNTIME
	_, b, _, _ = runtime.Caller(0)

	// ROOT FOLDER OF THIS PROJECT
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)