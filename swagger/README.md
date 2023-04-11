## Running Swagger UI
`
docker run -p 9000:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/foo/swagger.json -v ./../swagger:/foo swaggerapi/swagger-ui:v4.18.2
`

