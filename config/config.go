package config

import "time"

//ServerPort config
const ServerPort = "5000"

//MongodbURL mongodb connection string
const MongodbURL = "mongodb://localhost:27017"

//JwtSecret secret for JWT user
const JwtSecret = "secretstring"

//JwtExpireIn expire for JWT token
const JwtExpireIn = time.Minute * 30
