package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.5-k3s1"
	gitCommit    = "fc306bc167a9bbddc53edd6d782b058d2b582a18"
	gitTreeState = "clean"
	buildDate = "2026-05-12T20:07:20Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
