# Golang-Rate-Limiter

A demo application with rate-limiter middleware: 
- Only accept maximum 60 request within 60 seconds by IP
- Shows current request count from IP on the web if request count is less than limit.
- Shows "Error" if request count is greater than limit.

## How to run

- Prerequisites:
  - docker [installed](https://www.docker.com)
  - docker-compose [installed](https://docs.docker.com/compose/install/)
- docker-compose up -d
- endpoint: GET localhost:8080