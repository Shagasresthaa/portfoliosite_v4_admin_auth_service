package models

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// User defines the user model for the database
type User struct {
    ID            string    `gorm:"primaryKey;size:16" json:"id"`
    Username      string    `gorm:"size:255;not null" json:"username"`
    Email         string    `gorm:"size:255;not null;unique" json:"email"`
    Password      string    `gorm:"size:255;not null" json:"password"`
    Role          string    `gorm:"size:50;not null" json:"role"`
    Sponsor       string    `gorm:"size:50;not null" json:"sponsor"`
    TOTPSecret    string    `gorm:"size:255" json:"totpSecret"`
    IsTOTPEnabled bool      `gorm:"default:false" json:"isTotpEnabled"`
    IsActive      bool      `gorm:"default:true" json:"isActive"`
    CreatedAt     time.Time `json:"createdAt"`
    LastAccessed  time.Time `json:"lastAccessed"`
}

// NewRand creates a new seeded random generator for use in generating IDs
func NewRand() *rand.Rand {
    src := rand.NewSource(time.Now().UnixNano())
    return rand.New(src)
}

// GenerateID generates a random 16-character alphanumeric ID
func GenerateID(length int) string {
    rnd := NewRand() 
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rnd.Intn(len(charset))]
    }
    return string(b)
}

// NewUser creates a new user instance and automatically generates a unique ID for the user
func NewUser(username, email, password, role string, sponsor string) *User {
    return &User{
        ID:            GenerateID(16),
        Username:      username,
        Email:         email,
        Password:      password, 
        Role:          role,
        Sponsor:       sponsor,
        TOTPSecret:    "", // To be set when TOTP is implemented (Probably until UI is ready wont be in use)
        IsTOTPEnabled: false,
        IsActive:      true,
        CreatedAt:     time.Now(),
        LastAccessed:  time.Now(),
    }
}
