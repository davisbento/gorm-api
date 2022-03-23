# GO API with GORM

### Migrations (https://www.enterprisedb.com/postgres-tutorials/connecting-postgresql-using-psql-and-pgadmin)

```
migrate create -ext sql -dir db/migrations -seq create_users_table
```

```
POSTGRESQL_URL=postgres://davisbento:davisbento_pass@localhost:5432/articles?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
