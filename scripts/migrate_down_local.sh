export CATSSOCIAL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" && migrate -database ${CATSSOCIAL} -path ./db/migrations down
