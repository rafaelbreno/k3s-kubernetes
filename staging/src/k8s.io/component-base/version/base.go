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
	gitVersion   = "v1.32.4-k3s1"
	gitCommit    = "4742bcdcfb5dc69bb67f8071651df51f3c248d65"
	gitTreeState = "clean"
	buildDate = "2025-04-23T15:54:32Z"
)
