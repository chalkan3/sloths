curl -d '{"sloth": {"name": "maria", "family": "guicas"}}' -H "Content-Type: application/json" -X POST http://localhost:9000/v1/sloth
curl -d '{"sloth": {"name": "lady", "family": "guicas"}}' -H "Content-Type: application/json" -X PUT http://localhost:9000/v1/sloth/2
curl -X DELETE http://localhost:9000/v1/sloth/3
curl -X GET http://localhost:9000/v1/sloth/3
curl -X GET http://localhost:9000/v1/sloth/
