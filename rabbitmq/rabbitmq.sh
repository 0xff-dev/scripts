#!/bin/bash
# local test
docker pull rabbitmq:management
docker run -d --hostname localhost --name rabbitmq -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin -p 15672:15672 -p 5672:5672 rabbitmq:management

# feature we will use rabbitmq cluster to solve problems
