#!/bin/bash

docker build -t todowheel-be .
docker run -p 8080:8080 -t todowheel-be
