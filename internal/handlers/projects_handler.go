package handlers

import (
	"net/http"
	"portfoliosite_v4_admin_auth_service/internal/models"
	"portfoliosite_v4_admin_auth_service/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler creates a new user
func CreateProjectHandler(repo *repository.ProjectRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input struct {
            Name        string    `json:"name"`
            StartDate   time.Time `json:"start_date"`
            EndDate     time.Time `json:"end_date"`
            ShortDesc   string    `json:"short_desc"`
            LongDesc    string    `json:"long_desc"`
            ImageURL    string    `json:"image_url"`
            ProjectURL  string    `json:"project_url"`
            Status      string    `json:"status"`
            Affiliation string    `json:"affiliation"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }

        newProject := models.NewProject(input.Name, input.StartDate, input.EndDate, input.ShortDesc, input.LongDesc, input.ImageURL, input.ProjectURL, input.Status, input.Affiliation)
        if err := repo.CreateProject(newProject); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create project"})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"message": "Project created", "project": newProject})
    }
}


// ListAllProjectsHandler returns all Projects
func ListAllProjectsHandler(repo *repository.ProjectRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        projects, err := repo.ListAllProjects()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch projects"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"projects": projects})
    }
}