package payments

// Payment resource
type Payment struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Version        int64  `json:"version"`
	OrganisationID string `json:"organisation_id"`
}
