package repository

import (
	"portfoliosite_v4_admin_auth_service/internal/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
    DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
    return &ProjectRepository{DB: db}
}

// CreateProject adds a new project to the database
func (repo *ProjectRepository) CreateProject(project *models.Project) error {
    return repo.DB.Create(project).Error
}

// ListAllUsers retrieves all projects from the database
func (repo *ProjectRepository) ListAllProjects() ([]models.Project, error) {
    var projects []models.Project
    result := repo.DB.Find(&projects)
    return projects, result.Error
}

// GetProjectByID retrieves a single project by ID
func (repo *ProjectRepository) GetProjectByID(id int) (*models.Project, error) {
    var project models.Project
    result := repo.DB.Where("id = ?", id).First(&project)
    return &project, result.Error
}

// UpdateUser updates an existing project
func (repo *ProjectRepository) UpdateProject(project *models.Project) error {
    return repo.DB.Save(project).Error
}

// DeleteUser deletes a project by ID
func (repo *ProjectRepository) DeleteProject(id int) error {
    result := repo.DB.Delete(&models.Project{}, "id = ?", id)
    return result.Error
}
