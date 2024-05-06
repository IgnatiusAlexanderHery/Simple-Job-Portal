# Simple Job Portal API Documentation

This document outlines the API endpoints and usage for the Simple Job Portal.

## API Reference

- [Talent](#talent)
  - [View Jobs](#view-jobs)
  - [Apply for Job](#apply-for-job)
  - [View Applications](#view-applications)
  - [View All Applications](#view-all-applications)
  - [View Job Detail](#view-job-detail)
- [Employer](#employer)
  - [Create Job](#create-job)
  - [Process Application](#process-application)
  - [View Applications](#view-applications-employer)
  - [Update Job](#update-job)
  - [View Job Detail](#view-job-detail-employer)
  - [View All Applications](#view-all-applications-employer)
- [Authentication](#authentication)
  - [Register](#register)
  - [Login](#login)
  - [Logout](#logout)
  - [Profile](#profile)

## Talent

#### View Jobs

```http
  GET /talent/jobs
```

Retrieves a list of available jobs.

Response

```json
[
  {
    "id": 1,
    "title": "Software Engineering edited",
    "description": "Working as Software Engineering edited",
    "requirements": "1 year experience\nUsing Java edited",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "open"
  }
]
```

### Apply for Job

```http
  POST /talent/jobs/:jobID/apply
```

Allows a talent to apply for a job.

| Parameter | Type     | Description                               |
| :-------- | :------- | :---------------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to apply for. |

Request Headers

```json
X-Csrf-Token: your-auth-token
```

Response

```json
{
  "id": 1,
  "job_id": 1,
  "talent_id": "54c7353f-79bf-422e-b536-067465420751",
  "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
  "status": "applied"
}
```

### View Applications

```http
  GET /talent/jobs/:jobID/applications
```

Retrieves applications for a specific job.

| Parameter | Type     | Description                                           |
| :-------- | :------- | :---------------------------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to view applications for. |

Response

```json
[
  {
    "id": 1,
    "job_id": 1,
    "talent_id": "54c7353f-79bf-422e-b536-067465420751",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "applied"
  }
]
```

### View All Applications

```http
  GET /talent/applications
```

Retrieves all applications submitted by the talent.

Response

```json
[
  {
    "id": 1,
    "job_id": 1,
    "talent_id": "54c7353f-79bf-422e-b536-067465420751",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "applied"
  }
]
```

### View Job Detail

```http
  GET /employer/jobs/:jobID
```

Retrieves details of a specific job.

| Parameter | Type     | Description                          |
| :-------- | :------- | :----------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to view. |

Response

```json
{
  "id": 1,
  "title": "Software Engineering edited",
  "description": "Working as Software Engineering edited",
  "requirements": "1 year experience\nUsing Java edited",
  "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
  "status": "open"
}
```

## Employer

#### View Jobs

```http
  GET /employer/jobs
```

Retrieves a list of available jobs.

Response

```json
[
  {
    "id": 1,
    "title": "Software Engineering edited",
    "description": "Working as Software Engineering edited",
    "requirements": "1 year experience\nUsing Java edited",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "open"
  }
]
```

### Create Job

```http
  POST /employer/jobs
```

Allows an employer to create a new job listing.

Request Headers

```json
X-Csrf-Token: your-auth-token
```

Request Body

```json
{
  "title": "Software Engineering 3",
  "description": "Working as Software Engineering",
  "requirements": "1 year experience\nUsing Java"
}
```

Response

```json
{
  "id": 1,
  "title": "Software Engineering",
  "description": "Working as Software Engineering",
  "requirements": "1 year experience\nUsing Java",
  "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
  "status": "open"
}
```

### Process Application

```http
  POST /employer/applications/:applicationID
```

Allows an employer to process a job application.

| Parameter       | Type     | Description                                     |
| :-------------- | :------- | :---------------------------------------------- |
| `applicationID` | `string` | **Required**. ID of the application to process. |

Request Headers

```json
X-Csrf-Token: your-auth-token
```

Request Body

```json
{
  "status": "accept"
}
```

Response

```json
{
  "id": 1,
  "job_id": 1,
  "talent_id": "54c7353f-79bf-422e-b536-067465420751",
  "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
  "status": "accept"
}
```

### View Applications

```http
  GET /employer/jobs/:jobID/applications
```

Retrieves applications for a specific job posted by the employer.

| Parameter | Type     | Description                                           |
| :-------- | :------- | :---------------------------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to view applications for. |

Response

```json
[
  {
    "id": 1,
    "job_id": 1,
    "talent_id": "54c7353f-79bf-422e-b536-067465420751",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "accept"
  }
]
```

### Update Job

```http
  PUT /employer/jobs/:jobID
```

Allows an employer to update the status of a job listing.

| Parameter | Type     | Description                            |
| :-------- | :------- | :------------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to update. |

Request Body

```json
{
  "title": "Software Engineering edited",
  "description": "Working as Software Engineering edited",
  "requirements": "1 year experience\nUsing Java edited",
  "status": "open"
}
```

Response

```json
{
  "updatedJob": {
    "id": 1,
    "title": "Software Engineering edited",
    "description": "Working as Software Engineering edited",
    "requirements": "1 year experience\nUsing Java edited",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "open"
  }
}
```

### View Job Detail

```http
  GET /employer/jobs/:jobID
```

Retrieves details of a specific job posted by the employer.

| Parameter | Type     | Description                          |
| :-------- | :------- | :----------------------------------- |
| `jobID`   | `string` | **Required**. ID of the job to view. |

Response

```json
{
  "id": 1,
  "title": "Software Engineering edited",
  "description": "Working as Software Engineering edited",
  "requirements": "1 year experience\nUsing Java edited",
  "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
  "status": "open"
}
```

### View All Applications

```http
  GET /employer/applications
```

Retrieves all applications submitted to the employer.

Response

```json
[
  {
    "id": 1,
    "job_id": 1,
    "talent_id": "54c7353f-79bf-422e-b536-067465420751",
    "employer_id": "9dfbef79-ae43-4d88-85a0-269ee744bbf8",
    "status": "accept"
  }
]
```

## Authentication

### Register

```http
  POST /register
```

Registers a new user.

Request Body

```json
{
  "username": "john_doe",
  "password": "password123",
  "role": "talent"
}
```

Response

```json
{
  "id": 1,
  "uuid": "54c7353f-79bf-422e-b536-067465420751",
  "username": "john_doe",
  "password": "",
  "role": "talent"
}
```

#### Login

```http
  POST /login
```

Logs in a user.

Request Body

```json
{
  "username": "john_doe1",
  "password": "password123"
}
```

Response

```json
{
  "X-CSRF-Token": "mQi-DvvdjvXeY8HDWA7sCHbAClc=",
  "message": "Login successful"
}
```

### Logout

```http
  POST /logout
```

Logs out the current user.

Response

```json
{
  "message": "Logout successful"
}
```

### Profile

```http
  GET /profile
```

Retrieves the user's profile information.

Response

```json
{
  "User": {
    "username": "john_doe",
    "uuid": "54c7353f-79bf-422e-b536-067465420751",
    "role": "talent",
    "exp": 1714197420
  }
}
```
