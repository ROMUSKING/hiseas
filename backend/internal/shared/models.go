package shared

import "time"

// Common data models used across modules

type User struct {
	ID             int64     `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	Reputation     int       `json:"reputation" db:"reputation"`
	PrivacyLevel   string    `json:"privacy_level" db:"privacy_level"`
	EncryptedEmail []byte    `json:"-" db:"encrypted_email"` // Omit from JSON responses
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type Vessel struct {
	ID        int64     `json:"id" db:"id"`
	OwnerID   int64     `json:"owner_id" db:"owner_id"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Voyage struct {
	ID            int64      `json:"id" db:"id"`
	SkipperID     int64      `json:"skipper_id" db:"skipper_id"`
	VesselID      int64      `json:"vessel_id" db:"vessel_id"`
	Title         string     `json:"title" db:"title"`
	Description   string     `json:"description" db:"description"`
	StartTime     *time.Time `json:"start_time" db:"start_time"` // Use pointer for nullability
	EndTime       *time.Time `json:"end_time" db:"end_time"`
	StartLocation string     `json:"start_location" db:"start_location"` // WKT or GeoJSON
	PlannedRoute  string     `json:"planned_route" db:"planned_route"`   // WKT or GeoJSON
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}
