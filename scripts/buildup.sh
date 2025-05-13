#!/bin/bash

docker-compose down -v && \
docker-compose up --build -d && \
echo "Neo4j initializing" && \
docker-compose logs -f neo4j | grep -m 1 "Started." && \
docker-compose logs -f app monitor