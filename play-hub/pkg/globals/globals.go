package globals

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ActiveUser string
var Client *mongo.Client
var IstLocation *time.Location
