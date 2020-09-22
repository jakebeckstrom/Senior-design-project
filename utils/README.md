# Utils

## Running web app with docker-compose

1. Make sure you have docker installed locally.
2. Run `docker-compose pull` to ensure all services are using their most recent
   images.
3. Run `docker-compose up`. This will use images hosted on Docker Hub (ex: our [API image](https://hub.docker.com/repository/docker/csci4950tgt/api)) that are
   automatically built from each service's github master branch

That should be it.

**Optional**

To run with locally built images, do the following:

1. Copy the docker-compose file into the root project directory, where `api`,
   `frontend`, and `honeyclient` are all subdirectories

2. Run `docker-compose -f docker-compose.development.yml up --build` (this will take a while to build the first time,
   runs using locally built images from subdirectories)
