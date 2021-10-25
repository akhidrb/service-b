# Service B

### Description
This service has two major responsibilities:
- Consuming data from kafka and storing them in MongoDB
    - Orders are separate by country and stored in different Collections in MongoDB for faster querying. It seemed the right approach in this exercise because we already know the countries that the data being processed maps too where if it was something dynamic and dependent on the user then this approach would not necessarily be the best.
    - Orders are grouped together based on the weight limit after they are consumed and before they are added to the database. This approach was done to avoid excessive processing when API calls are made on the endpoint to retrieve the daily manifest. It came in handy that the weight limit was known beforehand and not passed as a parameter when querying for the daily manifest.
- API that returns a daily manifest file of the orders based on the timestamp of their addition to the database. 
    - So an index of the `created_at` field is created for faster querying.
    - So my only query filter is the `created_at` because I want to group all the orders together from all countries but I want all the ones in the same day.
#### How To Run:
The docker-compose file in this directory contains the images for MongoDB and Service B.
Apply the following commands to run:
- `make all`
- `docker-compose up -d`

#### Run Tests
`go test -v ./service/*`

#### API Documentation
Copy the contents of the [Swagger file](api/v1/v1.yml) in this project located in the following path `api/v1/v1.yml` into [swagger online viewer](https://editor.swagger.io/)