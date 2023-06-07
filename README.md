
# Hospital-API


## Installation

Install this project

```bash
  cp .env-example .env
  {configure your .env}
  go mod tidy
  go run cmd/server/main.go
```

Run the database migration Up or Down

```bash
  make migrateup
  make migratedown
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`APP_PORT`

`JWT_SECRET`

`DB_DRIVER` `DB_HOST` `DB_PORT` `DB_USER` `DB_PASS` `DB_NAME`

`CLOUDINARY_CLOUD_NAME` `CLOUDINARY_API_KEY` `CLOUDINARY_API_SECRET` `CLOUDINARY_UPLOAD_FOLDER`
