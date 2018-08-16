#!/bin/bash

 export GITHUB_TOKEN="xxxxxxxxxxxxxxxxxxx"
go run main.go -o golang -r go -i 26998 >  issue.html
go run main.go -o golang -r go -i 26969 >  issue2.html
open issue.html issue2.html