package models

import "time"

// Project represents a project with all its details.
type Project struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    ShortDesc   string    `json:"short_desc"`
    LongDesc    string    `json:"long_desc"`
    ImageURL    string    `json:"image_url"`
    ProjectURL  *string   `json:"project_url,omitempty"`
    Status      string    `json:"status"`
    Affiliation string    `json:"affiliation"`
}
