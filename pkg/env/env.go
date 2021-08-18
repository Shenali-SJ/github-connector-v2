package env

import (
	model "github-connector/pkg/model"
	mongo "go.mongodb.org/mongo-driver/mongo"
	gorm "gorm.io/gorm"
)

var RDB *gorm.DB
var MDB *mongo.Database
var CMDCtrl *model.StorageCtrl
var QueryCtrl *model.StorageCtrl
var ESCtrl *model.MessageCtrl

const GRPC_PORT = "GRPC_PORT"
const REST_PORT = "REST_PORT"

var GRPCPORT string
var RESTPORT string
