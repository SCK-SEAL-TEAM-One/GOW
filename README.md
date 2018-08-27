[![Build Status](https://travis-ci.org/SCK-SEAL-TEAM-One/GOW.svg?branch=master)](https://travis-ci.org/SCK-SEAL-TEAM-One/GOW) 
[![Coverage Status](https://coveralls.io/repos/github/SCK-SEAL-TEAM-One/GOW/badge.svg?branch=master)](https://coveralls.io/github/SCK-SEAL-TEAM-One/GOW?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/SCK-SEAL-TEAM-One/GOW)](https://goreportcard.com/report/github.com/SCK-SEAL-TEAM-One/GOW)

## Instruction to run
* Update go to go 1.11
* Change dir to gow-backend/gow
* Run $go test ./...
* Run $go build .
* API should be work corectly


## ติดตั้ง Library สำหรับโปรเจคนี้
```
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
docker pull mysql:5.7
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=sckshuhari -d mysql:5.7
```


