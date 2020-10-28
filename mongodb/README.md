# mongoDB Serivce Documentation

## Service Author
Ron Arbo

## Description
This service will create an instance of a mongoDB database to be accessed by our Flask services. They are created using the default mongoDB image from DockerHub. The only other thing we do in our Dockerfile is expose port 27017, the default port for mongoDB. This (along with docker-compose networking), allow our Flask applications to access the instance of mongoDB. 