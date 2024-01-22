# Bank Application using Go

## PostGres installation
``` bash
docker pull postgres:16-alpine
docker run --name dev-postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=bank-app -d -p 5432:5432 postgres:16-alpine
docker exec -it <container-id> psql -U <dbname> <dbUser>
```

### DB migration
```bash
# uses golang-migrate
migrate create -ext sql -dir db/migration -seq init_schema
```