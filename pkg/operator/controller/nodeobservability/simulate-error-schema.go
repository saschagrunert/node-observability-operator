package nodeobservabilitycontroller

const (
	saObj     = "serviceaccount"
	sccObj    = "securitycontextconstraint"
	crObj     = "clusterrole"
	crbObj    = "clusterrolebinding"
	secretObj = "secret"
	dsObj     = "daemonset"
)

// ErrTestObject - simple struct used to inject errors for testing
type ErrTestObject struct {
	Set      map[string]bool
	NotFound map[string]bool
}
