# Nginx Service Documentation

## Service Authors
Ron Arbo & TJ Goldblatt

## Description
This is one of our more intuitive services. We create the nginx instance using a Dockerfile and the default nginx image given on DockerHub. As a part of the Dockerfile, we remove the default Nginx configuration file within the container and replace it with our own local version that will apply specifically to our application. This configuration file does a couple things. It serves on port 80, and uses a proxy to serve our UI on the empty endpoint, '/'. It also routes other requests to the particular service that will be able to fulfill them, using the URL endpoint to do so. Our nginx service also uses uWSGI to assist in the connection to our Flask applications. Like all of our services, the nginx service is referenced in our docker-compose file so we can connect it to our other services over the common network shared between them.