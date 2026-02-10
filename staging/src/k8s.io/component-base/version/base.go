package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.12-k3s1"
	gitCommit    = "845ed55ea258391ed71904477f2161504773029f"
	gitTreeState = "clean"
	buildDate = "2026-02-10T22:14:31Z"
)
