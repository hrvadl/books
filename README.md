# Book service

This is a small project to play with the GCP & Terraform.

## Provisioning needed resources

1. `cd infra/terraform`
2. `gcloud auth login` (or use service account creds)
3. `terraform init`
4. `terraform plan`
5. `terraform apply`

## Running in docker

1. Populate `.env` file
2. Run `docker compose up -d`
3. Run a sample request

## Sample request

```sh
curl http://0.0.0.0:3000/v1/users -X POST -d '{"name":"vadym","email":"test@test.com","preferredGenres":["tech"]}' -H "Content-Type: application/json"
```

After user is created, message is populated into the pub/sub topic where it's consumed and logged to the console.

## TODO:

1. Add more meaning to this project :)
