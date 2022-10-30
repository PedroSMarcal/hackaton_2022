package adapters

type Parameters struct {
	User     string
	Password string
}

type GetClient struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type BodyConnecorId struct {
	ConnectorID string
	Parameters  Parameters
}
