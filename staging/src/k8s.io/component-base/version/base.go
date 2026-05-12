package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.8-k3s1"
	gitCommit    = "ce2375645df005fdaaf166804bf5e2b2e2fec891"
	gitTreeState = "clean"
	buildDate = "2026-05-12T20:02:00Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)
