protoc --go_out=. --go-grpc_out=. proto/*.proto
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB"
docker exec -i tbdb81bc1a1e5 /bin/bash