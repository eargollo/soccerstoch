package futebol

type HTTPClient interface {
	Get(endpoint ...string) ([]byte, error)
}
