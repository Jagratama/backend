# 🛡️ Jagratama Backend

The **Jagratama Backend** is a powerful and scalable document request and approval API service built with **Golang**, designed for organizations that require structured multi-level digital approvals, file uploads, and email notifications.

This backend is the core engine powering the **Jagratama Platform**, providing a secure and efficient workflow for document signing and review.

---

## 🚀 Features

- 🔐 Authentication with JWT
- 🧾 Document request, approval, and rejection system
- 🗂️ Multi-role user management (requester, approver, reviewer)
- ☁️ File upload to AWS S3 using AWS SDK
- 📬 Email notifications via Resend
- 📊 PostgreSQL with GORM ORM
- ✨ Clean, RESTful API design with Echo

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **Framework:** Echo
- **ORM:** GORM
- **Database:** PostgreSQL
- **Auth:** JWT
- **File Storage:** AWS S3 (via AWS SDK)
- **Email:** Resend
- **Others:** Docker (optional)


---

## 🔧 Prerequisites

Before running this project, make sure you have:

- Go 1.20+
- PostgreSQL running locally or remotely
- AWS account and S3 bucket configured
- Resend API key
- Git
- Docker

---

## 📄 Environment Variables (`.env`)

You need to set up an `.env` file, the template is provided in `.env.example`. Copy it to `.env` and fill in the required values:

---

## 🚀 Installation


### 1. Clone the repository
```
git clone https://github.com/your-username/jagratama-backend.git
```

### # 2. Install dependencies
```
go mod tidy
```

### 3. Run the server
```
go run main.go
```


### You also can run inside Docker (recommended)
```
docker build . -t jagratama-backend:1.0
docker run -d -p 80:8000 -e AWS_ACCESS_KEY_ID= -e AWS_SECRET_ACCESS_KEY= -e AWS_DEFAULT_REGION= jagratama-backend:1.0
```