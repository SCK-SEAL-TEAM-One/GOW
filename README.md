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


