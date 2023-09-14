API Framework
==============

Setup
------------

1. cp .env.example .env
2. cp config/config.example.json config/config.json
3. update file according to your local configuration
4. install node js (optional to enable live reload/hot reload)
5. npm -g install nodemon (optional to enable live reload/hot reload)
6. if follow step 4 and 5, run app with : nodemon --exec go run main.go --signal SIGTERM 
11. if not follow step 4 and 5, run app with : go run main.go
12. hit http://localhost:3000
