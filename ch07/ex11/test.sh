#!/bin/bash

curl http://localhost:8080/items
curl http://localhost:8080/item?id=1
curl http://localhost:8080/update?id=2&name=test&price=2
curl http://localhost:8080/delete?id=1
curl http://localhost:8080/create?name=new&price=229734