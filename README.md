# Service B

### Description
This service has two major responsibilities:
- Consuming data from kafka and storing them in MongoDB
    - Orders are separate by country and stored in different Collections in MongoDB for faster querying
- API that returns a daily manifest file of the orders based on the timestamp of their addition to the database. 
    - So an index of the `created_at` field is created for faster querying.

#### How To Run:
The docker-compose file in this directory contains the images for MongoDB and Service B.
Apply the following commands to run:
- `make all`
- `docker-compose up -d`

#### Run Tests
`go test -v ./service/*`

#### API Documentation
Copy the contents of the [Swagger file](api/v1/v1.yml) in this project located in the following path `api/v1/v1.yml` into [swagger online viewer](https://editor.swagger.io/)