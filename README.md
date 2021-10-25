# Service B

### Description
This service has two major responsibilities:
- Consuming data from kafka and storing them in MongoDB
    - Orders are separate by country and stored in different Collections in MongoDB for faster querying. It seemed the right approach in this exercise because we already know the countries that the data being processed maps too where if it was something dynamic and dependent on the user then this approach would not necessarily be the best.
    - Orders are grouped together based on the weight limit after they are consumed from kafka and before they are added to the database. This approach was done to avoid excessive processing when API calls are made on the endpoint to retrieve the daily manifest. It came in handy that the weight limit was known beforehand and not passed as a parameter when querying for the daily manifest. 
    - The weight limit is passed as an argument when running the service so it can be changed whenever needed. Changing the weight limit at run time might cause a problem because it can be changed while data is being consumed so that would cause inconsistency in the data being stored in the database. Passing the weight limit can be useful if we want to run several instances of the service and storing them on different databases if we have couriers with Vans of different sizes and we want to run an instance for each weight limit.
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