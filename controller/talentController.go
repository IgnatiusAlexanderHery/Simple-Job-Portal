package controller

import (
	"Simple-Job-Portal/database"
	"Simple-Job-Portal/model"
	"Simple-Job-Portal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ViewJobs(c *gin.Context) {

	jobs, err := database.GetJobsByStatus("open")
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}
	utils.SuccessfulResponse(c, jobs)
}

func ApplyForJob(c *gin.Context) {
	var application model.Application
	jobID := c.Param("jobID")
	jobIDUint, err := strconv.ParseUint(jobID, 10, 64)
	if err != nil {
		utils.BadRequestError(c, "Invalid job ID")
		return
	}

	user, status := c.Get("user")
	if !status {
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}
	application.TalentID = user.(*model.Claims).UUID
	application.JobID = uint(jobIDUint)

	job, err := database.GetJobByID(application.JobID)
	if err != nil {
		utils.NotFoundError(c, "Job not found")
		return
	}

	_, err = database.GetUserByUUID(application.TalentID)
	if err != nil {
		utils.NotFoundError(c, "Talent not found")
		return
	}

	application.Status = "applied"
	application.EmployerID = job.EmployerID

	err = database.AddApplication(&application)
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}

	utils.SuccessfulResponse(c, application)
}

func ViewApplications(c *gin.Context) {
	jobID := c.Param("jobID")
	user, status := c.Get("user")
	if !status {
		utils.UnauthorizedError(c, "Unauthorized")
		return
	}
	if jobID == "" {
		var applications []model.Application
		var err error
		if user.(*model.Claims).Role == "talent" {
			applications, err = database.GetApplicationsByTalentID(user.(*model.Claims).UUID)
			if err != nil {
				utils.BadRequestError(c, err.Error())
				return
			}

		} else if user.(*model.Claims).Role == "employer" {
			applications, err = database.GetApplicationsByEmployerID(user.(*model.Claims).UUID)
			if err != nil {
				utils.BadRequestError(c, err.Error())
				return
			}
		}
		utils.SuccessfulResponse(c, applications)
		return
	}
	jobIDUint, err := strconv.ParseUint(jobID, 10, 64)
	if err != nil {
		utils.BadRequestError(c, "Invalid job ID")
		return
	}

	applications, err := database.GetApplicationsByJobID(uint(jobIDUint))
	if err != nil {
		utils.BadRequestError(c, err.Error())
		return
	}
	utils.SuccessfulResponse(c, applications)
}

func ViewJobDetail(c *gin.Context) {
	jobID := c.Param("jobID")
	jobIDUint, err := strconv.ParseUint(jobID, 10, 64)
	if err != nil {
		utils.BadRequestError(c, "Invalid job ID")
		return
	}

	job, err := database.GetJobByID(uint(jobIDUint))
	if err != nil {
		utils.NotFoundError(c, "Job not found")
		return
	}
	utils.SuccessfulResponse(c, job)
}
