# File Management Site
This project is intended to teach me about Go, Go-Gin, a little bit of Docker and Kubernetes, as well as working more with API Gateway. It'll also give me more experience working with various DB types (so far I'm planning on using Postgres and MongoDB; maybe ElasticSearch for searching but I'll need to look into that more). The end goal is to have a file management implementing a microservice architecture.

### Done so Far
Just barely got a Docker container running Apache to house the static HTML and JS files. Got the backend for the first service partially Dockerized and setup to CORs headers to allow remote API calls.
### To-Do
Still need to set up the DB server and hook that up to the backend. Will probably need to refactor the CORs part when I migrate this service up to EC2 and setup API gateway to work with the API calls. Once that is taken care of, the first service will be up and running. 

#### Side Note
The front end of this site might look ugly for a while cause I'm more of a backend focused guy and don't know any cool and fancy frontend frameworks like React. Just jQuery.