package shared

// Common data models used across modules

type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Reputation     int    `json:"reputation"`
	PrivacyLevel   string `json:"privacy_level"`
	EncryptedEmail []byte `json:"encrypted_email"`
}

type Vessel struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Voyage struct {
	ID            int64  `json:"id"`
	SkipperID     int64  `json:"skipper_id"`
	VesselID      int64  `json:"vessel_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	StartLocation string `json:"start_location"` // WKT or GeoJSON
	PlannedRoute  string `json:"planned_route"`  // WKT or GeoJSON
}
