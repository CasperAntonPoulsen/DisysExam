# disysminiproject3

## Starting the service

Requirements:
- docker
- docker-compose

To start all of the services, use: 

docker-compose up

if you're developing and want to test new code remember to add --build to rebuild the images.

docker-compose up --build

If you want to define an additional service, you can add it in the docker-compose.yml file

## Making requests

curl -X POST localhost:8000/Increment
