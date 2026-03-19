package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "33"
	gitVersion   = "v1.33.10-k3s1"
	gitCommit    = "0995fb211e0ec59a98e6871aca751be0bf08ff19"
	gitTreeState = "clean"
	buildDate = "2026-03-19T14:14:53Z"
)
