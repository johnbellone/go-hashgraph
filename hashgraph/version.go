package hashgraph

const ProtocolVersion uint8 = 1

var (
	GitCommit   string
	GitDescribe string

	Version           = "unknown"
	VersionPrerelease = "unknown"
)
