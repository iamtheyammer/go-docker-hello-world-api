# go-docker-hello-world-api

A teeny tiny API on Docker. Returns a simple OK page. Good for testing this with containers.

## Usage

### Building

1. `docker build .`
2. `docker run -p 8002:8002 {id from last command}`
  (if you set the port, make sure you change those!)
  
### From docker hub

`docker run -p 8002:8002 iamtheyammer/go-docker-hello-world-api`

## Set the port

If you want to set the port the server runs on, you can customize that with the `PORT` environment variable.

Make sure you change the port in the docker commands-- it's `-p hostPort:containerPort`.
