
# Functionality and Important Design Decisions

he project architecture is based on Domain Driven Design. That helps us to decouple the domain logic and the application logic, in this case the application is a REST API.

## The cmd package

This package contains the main.go that is used to execute the program.

## The Internal package

This package contains all the code related to the service logic, here we going to find sibpackages like `api`, `domain`, `infrastructure`, and `shared`. More details of each one are provided in the next sections.

### api package

This is an interesting package because contains all the logic related to the web service, here we can find defined the functionality to create the REST API, functionality to define the needed endpoints, and we have defined two subpackages `dto` and `handler`.

**dto**:
* The use of the functionality defined here is important because help us to decouple the application and the domain "objects".
* Defines the structs to map the domain objects into response objects. 
* Defines the structs used to parse the request objects.
* Defines the validations needed for the request objects.

**handler**:
* This package defines the handler functions executed when a request is sent to our web service.

### domain package
* This package defines  the `contact` domain logic. (defined on its own package)
* Contains the `contact` entity
* Defines the `contact` repository
* Defines a `contact` service used to interact with the repository and with the REST API.

### infrastructure package
* Defines a subpackage related with the persistency. In this scenario we are using an in memory persistency approach.

* If we need to replace the in-memory persistency by a database, for example, as we are using dependency injection we can replace the logic easier.

### shared package
* This package defines two subpackages:

**apiconfig**
* Defines the logic to load the EnvVars used to configure our web service.

**errorhandler** 
* Is a little example to show how we can handle and/or create custom errors.

# How to execute the service
* In the root folder execute the following:

`go run cmd/main.go`


# Unit test

* For this example I'm going to test the domain service layer.
* I going to make use of mocking frameworks in order to mock and inject the dependencies without need to interact with external services, in this case the data storage.
* Will be used `gomock` https://github.com/golang/mock, to build the mock files.
* I'm going to follow the following approach for the unit test:

```
1. Define all the needed test functions to cover the functionality
2. For each one of the tests, define a test table that contains the posible test cases
3. Iterate the test cases in order to execute the test and ensure that all works as expected.
```

Example of how to define the mocks:
`mockgen -source=./repository.go -destination=./repository_mock.go -package=contact `

**Important Note:**
As we want to keep this example simple, I'm going to add just one unit test function, as a justification I can say for the other functions we can replicate and adapt the unit test. Let us think in this project as a template.

# Swagger
* swag init -d cmd,internal --parseDependency --parseInternal --parseDepth 8
* http://localhost:8080/swagger/index.html

# Request Examples

You can import the collection to `Postman` or `Insomnia` using the json on `insomnia_postman_collection` folder

```
curl --request GET \
  --url http://localhost:8080/ping
```

```
curl --request POST \
  --url http://localhost:8080/api/v1/contact \
  --header 'Content-Type: application/json' \
  --data '{
	    "email": "theemail@gmail.com",
		"lastName": "Last name",
		"name": "name",
		"phoneNumber": "3333333333"
}'
```

```
curl --request GET \
  --url http://localhost:8080/api/v1/contact/4e197e3f-8ac7-4774-a76d-a18b7e4c6f48
```

```
curl --request PUT \
  --url http://localhost:8080/api/v1/contact/4e197e3f-8ac7-4774-a76d-a18b7e4c6f48 \
  --header 'Content-Type: application/json' \
  --data '{
	    "email": "edited@gmail.com",
		"lastName": "Last name - edited",
		"name": "name - edited",
		"phoneNumber": "11111"
}'
```

```
curl --request DELETE \
  --url http://localhost:8080/api/v1/contact/4e166a6c-f843-485f-9430-f4d9d46bc0cf
```# challenge-accel-one
