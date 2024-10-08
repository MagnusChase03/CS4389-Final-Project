module github.com/MagnusChase03/CS4389-Project

replace github.com/MagnusChase03/CS4389-Project/models => ./models

replace github.com/MagnusChase03/CS4389-Project/routes => ./routes
replace github.com/MagnusChase03/CS4389-Project/routes/authRoutes => ./routes/authRoutes
replace github.com/MagnusChase03/CS4389-Project/routes/userRoutes => ./routes/userRoutes


replace github.com/MagnusChase03/CS4389-Project/controllers => ./controllers
replace github.com/MagnusChase03/CS4389-Project/controllers/authControllers => ./controllers/authControllers
replace github.com/MagnusChase03/CS4389-Project/controllers/userControllers => ./controllers/userControllers

replace github.com/MagnusChase03/CS4389-Project/middleware => ./middleware
replace github.com/MagnusChase03/CS4389-Project/utils => ./utils
replace github.com/MagnusChase03/CS4389-Project/db => ./db

go 1.23.1

require (
    filippo.io/edwards25519 v1.1.0 // indirect
    github.com/cespare/xxhash/v2 v2.1.2 // indirect
    github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
    github.com/go-redis/redis/v8 v8.11.5 // indirect
    github.com/go-sql-driver/mysql v1.8.1 // indirect
)
