package model

type BillingDetails struct {
	UserID               string `json:"user_id"`
	CreationDate         string `json:"creation_date"`
	PersonType           string `json:"person_type"`
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	CompanyName          string `json:"company_name"`
	CountryRegion        string `json:"country_region"`
	StreetAddress        string `json:"street_address"`
	City                 string `json:"city"`
	StateProvince        string `json:"state_province"`
	PostalCode           string `json:"postal_code"`
	Phone                string `json:"phone"`
	EmailAddress         string `json:"email_address"`
	OrderNotes           string `json:"order_notes"`
}
