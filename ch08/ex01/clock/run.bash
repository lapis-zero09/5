#!/bin/bash

TZ=US/Eastern go run main.go -port 8010 & TZ=Asia/Tokyo go run main.go -port 8020 & TZ=Europe/London go run main.go -port 8030
