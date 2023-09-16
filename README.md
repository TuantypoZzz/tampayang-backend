API Framework
==============

Migration
------------

1. install migration package : go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
2. Create migration : migrate create -ext sql -dir migrations {filename}
3. make sure to be specified when delacring {filename} and not too complicated
4. run migration : migrate -source file://./migrations -database "mysql://user:password@tcp(host:3306)/database" up
5. run migration : migrate -source file://./migrations -database "mysql://user:password@tcp(host:3306)/database" down

Setup
------------

1. Install migration package from
2. cp .env.example .env
3. cp config/config.example.json config/config.json
4. update file according to your local configuration
5. install node js (optional to enable live reload/hot reload)
6. npm -g install nodemon (optional to enable live reload/hot reload)
7. if follow step 4 and 5, run app with : nodemon --exec go run main.go --signal SIGTERM 
8. if not follow step 4 and 5, run app with : go run main.go
9. hit http://localhost:3000

