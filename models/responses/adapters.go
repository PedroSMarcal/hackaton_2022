package responses

import "time"

type APIKey struct {
	APIKey string `json:"apiKey"`
}

type ConnectorResponse struct {
	ID        string `json:"id"`
	Connector struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		PrimaryColor   string `json:"primaryColor"`
		InstitutionURL string `json:"institutionUrl"`
		Country        string `json:"country"`
		Type           string `json:"type"`
		Credentials    []struct {
			Label             string `json:"label"`
			Name              string `json:"name"`
			Type              string `json:"type"`
			Placeholder       string `json:"placeholder"`
			Validation        string `json:"validation"`
			ValidationMessage string `json:"validationMessage"`
			Optional          bool   `json:"optional"`
		} `json:"credentials"`
		ImageURL string `json:"imageUrl"`
		HasMFA   bool   `json:"hasMFA"`
		Health   struct {
			Status string      `json:"status"`
			Stage  interface{} `json:"stage"`
		} `json:"health"`
		Products  []string  `json:"products"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"connector"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Status          string      `json:"status"`
	ExecutionStatus string      `json:"executionStatus"`
	LastUpdatedAt   interface{} `json:"lastUpdatedAt"`
	WebhookURL      interface{} `json:"webhookUrl"`
	Error           interface{} `json:"error"`
	ClientUserID    interface{} `json:"clientUserId"`
	StatusDetail    interface{} `json:"statusDetail"`
	Parameter       interface{} `json:"parameter"`
}
type ItemsResponse struct {
	ID        string `json:"id"`
	Connector struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		PrimaryColor   string `json:"primaryColor"`
		InstitutionURL string `json:"institutionUrl"`
		Country        string `json:"country"`
		Type           string `json:"type"`
		Credentials    []struct {
			Label             string `json:"label"`
			Name              string `json:"name"`
			Type              string `json:"type"`
			Placeholder       string `json:"placeholder"`
			Validation        string `json:"validation"`
			ValidationMessage string `json:"validationMessage"`
			Optional          bool   `json:"optional"`
		} `json:"credentials"`
		ImageURL string `json:"imageUrl"`
		HasMFA   bool   `json:"hasMFA"`
		Health   struct {
			Status string      `json:"status"`
			Stage  interface{} `json:"stage"`
		} `json:"health"`
		Products  []string  `json:"products"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"connector"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Status          string      `json:"status"`
	ExecutionStatus string      `json:"executionStatus"`
	LastUpdatedAt   interface{} `json:"lastUpdatedAt"`
	WebhookURL      interface{} `json:"webhookUrl"`
	Error           interface{} `json:"error"`
	ClientUserID    interface{} `json:"clientUserId"`
	StatusDetail    interface{} `json:"statusDetail"`
	Parameter       interface{} `json:"parameter"`
}

type AccountResponse struct {
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
	Page       int `json:"page"`
	Results    []struct {
		ID            string      `json:"id"`
		Type          string      `json:"type"`
		Subtype       string      `json:"subtype"`
		Name          string      `json:"name"`
		Balance       float64     `json:"balance"`
		CurrencyCode  string      `json:"currencyCode"`
		ItemID        string      `json:"itemId"`
		Number        string      `json:"number"`
		MarketingName interface{} `json:"marketingName"`
		TaxNumber     string      `json:"taxNumber"`
		Owner         string      `json:"owner"`
		BankData      struct {
			TransferNumber string      `json:"transferNumber"`
			ClosingBalance interface{} `json:"closingBalance"`
		} `json:"bankData"`
		CreditData interface{} `json:"creditData"`
	} `json:"results"`
}

type TransactionResults struct {
	ID                 string      `json:"id"`
	Description        string      `json:"description"`
	DescriptionRaw     string      `json:"descriptionRaw"`
	CurrencyCode       string      `json:"currencyCode"`
	Amount             float64     `json:"amount"`
	Date               time.Time   `json:"date"`
	Category           interface{} `json:"category"`
	Balance            float64     `json:"balance"`
	AccountID          string      `json:"accountId"`
	ProviderCode       interface{} `json:"providerCode"`
	Status             string      `json:"status"`
	PaymentData        PaymentData `json:"paymentData"`
	Type               string      `json:"type"`
	CreditCardMetadata interface{} `json:"creditCardMetadata"`
	Merchant           interface{} `json:"merchant"`
}

type TransactionResponse struct {
	Total      int                  `json:"total"`
	TotalPages int                  `json:"totalPages"`
	Page       int                  `json:"page"`
	Results    []TransactionResults `json:"results"`
}

type PaymentData struct {
	Payer struct {
		Name           string `json:"name"`
		BranchNumber   string `json:"branchNumber"`
		AccountNumber  string `json:"accountNumber"`
		RoutingNumber  string `json:"routingNumber"`
		DocumentNumber struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"documentNumber"`
		RoutingNumberISPB string `json:"routingNumberISPB"`
	} `json:"payer"`
	Reason   string `json:"reason"`
	Receiver struct {
		Name           string `json:"name"`
		BranchNumber   string `json:"branchNumber"`
		AccountNumber  string `json:"accountNumber"`
		RoutingNumber  string `json:"routingNumber"`
		DocumentNumber struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"documentNumber"`
		RoutingNumberISPB string `json:"routingNumberISPB"`
	} `json:"receiver"`
	PaymentMethod   string `json:"paymentMethod"`
	ReferenceNumber string `json:"referenceNumber"`
}
