package env

import (
    gorm "gorm.io/gorm"
	model "migrdoctor/pkg/model"
	"os"
)

var RDB *gorm.DB
var CMDCtrl *model.StorageCtrl
var QueryCtrl *model.StorageCtrl
var ESCtrl *model.MessageCtrl

var REST_PORT = "REST_PORT"

var RestPort = "8080"

var GRPC_PORT = "GRPC_PORT"

var GrpcPort = "5051"

const E3_DIALET = "E3_DIALET"

var E3DIALET = "mysql"

const E3_DB = "E3_DB"

var E3db = "hospitaldb"

const E3_HOST = "E3_HOST"

var E3host = "127.0.0.1"

const E3_PORT = "E3_PORT"

var E3port = "3306"

const E3_USER = "E3_USER"

var E3user = "root"

const E3_PWD = "E3_PWD"

var E3pwd = "root"

const E3_URL = "E3_URL"

var E3url = "127.0.0.1"

var DEFAULTDIALET string
var DEFAULTDB string
var DEFAULTHOST string
var DEFAULTPORT string
var DEFAULTUSER string
var DEFAULTPWD string
var DEFAULTURL string
//var REST_PORT int64
//var GRPC_PORT int64

const DEFAULT_DIALET = "DEFAULT_DIALET"
const DEFAULT_DB = "DEFAULT_DB"
const DEFAULT_HOST = "DEFAULT_HOST"
const DEFAULT_PORT = "DEFAULT_PORT"
const DEFAULT_USER = "DEFAULT_USER"
const DEFAULT_PWD = "DEFAULT_PWD"
const DEFAULT_URL = "DEFAULT_URL"
//const REST_PORT = "REST_PORT"
//const GRPC_PORT = "GRPC_PORT"

var RDBHOST string
var RDBPORT string
var RDBDBNAME string
var RDBUSER string
var RDBPASSWORD string

func LoadEnvs() {

	if val := os.Getenv(E3_DIALET); val != "" {
		E3DIALET = val
	}
	if val := os.Getenv(E3_DB); val != "" {
		E3db = val
	}
	if val := os.Getenv(E3_HOST); val != "" {
		E3host = val
	}
	if val := os.Getenv(E3_PORT); val != "" {
		E3port = val
	}
	if val := os.Getenv(E3_USER); val != "" {
		E3user = val
	}
	if val := os.Getenv(E3_PWD); val != "" {
		E3pwd = val
	}
	if val := os.Getenv(E3_URL); val != "" {
		E3url = val
	}
}
