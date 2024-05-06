package database

import (
	"fmt"
	"log"

	"Simple-Job-Portal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbDriver   = "postgres"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "postgres"
	dbHost     = "localhost"
	dbPort     = 5432
)

var db *gorm.DB
var err error

func init() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Job{}, &model.Application{})

	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	fmt.Println("Database connected and migrated successfully")
}

func GetDatabase() *gorm.DB {
	return db
}

func AddUser(user *model.User) error {
	return db.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.First(&user, "username = ?", username).Error
	return &user, err
}

func GetUserByUUID(uuid string) (*model.User, error) {
	var user model.User
	err := db.First(&user, "uuid = ?", uuid).Error
	return &user, err
}

func AddJob(job *model.Job) error {
	return db.Create(job).Error
}

func GetJobByID(jobID uint) (*model.Job, error) {
	var job model.Job
	err := db.First(&job, jobID).Error
	return &job, err
}

func GetJobsByStatus(status string) ([]model.Job, error) {
	var jobs []model.Job
	err := db.Where("status = ?", status).Find(&jobs).Error
	return jobs, err
}

func GetJobsByEmployerID(employerID string) ([]model.Job, error) {
	var jobs []model.Job
	err := db.Where("employer_id = ?", employerID).Find(&jobs).Error
	return jobs, err
}

func UpdateJob(job *model.Job) error {
	return db.Save(job).Error
}

func AddApplication(application *model.Application) error {
	return db.Create(application).Error
}

func GetApplicationsByJobID(jobID uint) ([]model.Application, error) {
	var applications []model.Application
	err := db.Where("job_id = ?", jobID).Find(&applications).Error
	return applications, err
}

func GetApplicationsByTalentID(talentID string) ([]model.Application, error) {

	var applications []model.Application
	err := db.Where("talent_id = ?", talentID).Find(&applications).Error
	return applications, err
}

func GetApplicationsByEmployerID(employerID string) ([]model.Application, error) {
	var applications []model.Application
	err := db.Where("employer_id = ?", employerID).Find(&applications).Error
	return applications, err
}

func UpdateApplication(application *model.Application) error {
	return db.Save(application).Error
}

func GetApplicationsById(applicationID uint) (*model.Application, error) {
	var application model.Application
	err := db.First(&application, applicationID).Error
	return &application, err
}
