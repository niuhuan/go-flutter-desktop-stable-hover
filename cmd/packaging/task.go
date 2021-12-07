package packaging

import "github.com/niuhuan/stable-hover/internal/build"

// Task contains all configuration options for a given packaging method.
// TODO: Rename to something that suits it more? Mabe Executor?
type Task interface {
	Name() string
	Init()
	IsInitialized() bool
	AssertInitialized()
	Pack(buildVersion string, mode build.Mode)
	IsSupported() bool
	AssertSupported()
}
