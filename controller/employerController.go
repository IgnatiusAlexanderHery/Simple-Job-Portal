package controller

import (
	"Simple-Job-Portal/database"
	"Simple-Job-Portal/model"
	"Simple-Job-Portal/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	var job model.Job
	if err := c.BindJSON(&job); err != nil {
		if err.Error() == "EOF" {
			utils.BadRequestError(c, "All fields (title, description, requirements) are required")
			return
		}
		utils.BadRequestError(c, err.Error())
		return
	}
	if job.Title == "" || job.Description == "" || job.Requirements == "" {
		utils.BadRequestError(c, "All fields (title, description, requirements) are required")
		return
	}

	user, status := c.Get("user")
	if !status {
		utils.UnauthorizedError(c, "Unauthorized 6")
		return
	}
	employer, err := database.GetUserByUUID(user.(*model.Claims).UUID)
	if err != nil {
		utils.NotFoundError(c, "Employer not found")
		return
	}

	job.EmployerID = employer.UUID
	job.Status = "open"

	err = database.AddJob(&job)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	utils.SuccessfulResponse(c, job)
}

func UpdateJob(c *gin.Context) {
	jobID := c.Param("jobID")
	var job model.Job
	if err := c.BindJSON(&job); err != nil {
		if err.Error() == "EOF" {
			utils.BadRequestError(c, "Status fields are required")
			return
		}
		utils.BadRequestError(c, err.Error())
		return
	}
	if job.Status != "open" && job.Status != "closed" {
		utils.BadRequestError(c, "Invalid status value [open, closed]")
		return
	}

	user, status := c.Get("user")
	if !status {
		fmt.Println("Unauthorized")
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}

	employer, err := database.GetUserByUUID(user.(*model.Claims).UUID)
	if err != nil {
		utils.NotFoundError(c, "Employer not found")
		return
	}
	if employer.UUID != user.(*model.Claims).UUID {
		fmt.Println(employer.UUID, user.(*model.Claims).UUID)
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}

	updatedJob := job

	jobIDUint, err := strconv.ParseUint(jobID, 10, 64)
	if err != nil {
		utils.BadRequestError(c, "Invalid job ID")
		return
	}

	dbJob, err := database.GetJobByID(uint(jobIDUint))
	if err != nil {
		utils.NotFoundError(c, "Job not found")
		return
	}

	updatedJob.ID = dbJob.ID
	updatedJob.EmployerID = user.(*model.Claims).UUID
	if updatedJob.Title == "" {
		updatedJob.Title = dbJob.Title
	}
	if updatedJob.Description == "" {
		updatedJob.Description = dbJob.Description
	}
	if updatedJob.Requirements == "" {
		updatedJob.Requirements = dbJob.Requirements
	}

	err = database.UpdateJob(&updatedJob)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	utils.SuccessfulResponse(c, updatedJob)
}

func ProcessApplication(c *gin.Context) {
	applicationID := c.Param("applicationID")

	var requestBody struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		if err.Error() == "EOF" {
			utils.BadRequestError(c, "Invalid status value [interview, accept, reject]")
			return
		}
		utils.BadRequestError(c, err.Error())
		return
	}

	if requestBody.Status == "" && requestBody.Status != "interview" && requestBody.Status != "accept" && requestBody.Status != "reject" {
		utils.BadRequestError(c, "Invalid status value [interview, accept, reject]")
		return
	}
	applicationIDUint, err := strconv.ParseUint(applicationID, 10, 64)
	if err != nil {
		utils.BadRequestError(c, "Invalid application ID")
		return
	}

	application, err := database.GetApplicationsById(uint(applicationIDUint))
	if err != nil {
		utils.NotFoundError(c, "Application not found")
		return
	}

	application.Status = requestBody.Status

	err = database.UpdateApplication(application)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	utils.SuccessfulResponse(c, application)
}

func ViewJobsByEmployer(c *gin.Context) {
	user, _ := c.Get("user")
	jobs, err := database.GetJobsByEmployerID(user.(*model.Claims).UUID)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	utils.SuccessfulResponse(c, jobs)
}
