
# E-commerce Device Mart 

This is an E-commerce web API built using Go, Gin, PostgreSQL. This application built on top of Clean Architecture for decoupling and ease of maintainability.





## Project Overview

An E-commerce application meant to sell Laptops that relying to specific Brands.
The app is a blueprint of an E-commerce application which serves the basic and some advanced functionality of an E-commerce app.   



## Technologies

- Go Programming Language

- Gin Framework

- PostgreSQL

- GORM 

- AWS(S3)

- JWT

- Swagger

- Wire


## Run Locally

Clone the project

```bash
  git clone https://github.com/anazibinurasheed/ecommerce-device-mart.git
```

Get into project directory

```bash
  cd ecommerce-device-mart
```
Create .env file and configure with yours
(command is for linux)
```bash
touch .env
```


Install dependencies

```bash
  make deps
  
  go mod tidy
```

Setup Environment Variables 
```.env
DB_HOST=
DB_NAME=
DB_USER=
DB_PORT=
DB_PASSWORD=
ADMIN=
ADMINPASS=
JWT_SECRET=
TWILIO_ACCOUNT_SID=
TWILIO_AUTH_TOKEN=
VERIFY_SERVICE_SID=
RAZORPAY_KEY_ID=
RAZORPAY_KEY_SECRET=
AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
S3_BUCKET_NAME=
S3_BUCKET_MEDIA_PATH= (ex: folder/)
PORT=
```
Start the server

```bash
  make run
```

## Access API
`http://localhost:3000/swagger/index.html`
<!-- 
