package main

import (
    "database/sql"
    "fmt"
    mysql "gorm.io/driver/mysql"
    //gorm "github.com/jinzhu/gorm"
    gorm "gorm.io/gorm"
    "migrdoctor/pkg/api_echo"
    model "migrdoctor/pkg/model"
    arfcreatedoctor "migrdoctor/pkg/rule/ARF_CreateDoctor"
    trfmain "migrdoctor/pkg/rule/TRF_Main"
    xilogger "migrdoctor/pkg/xiLogger"
    "os"
)

import (
    "migrdoctor/pkg/env"
)
import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "migrdoctor/docs"
)
import (
    "flag"
    "github.com/golang/glog"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "migrdoctor/pkg/controller_echo"
)
import (
    "google.golang.org/grpc"
    "log"
    "migrdoctor/pkg/controller"
    "migrdoctor/pub"
    "net"
)

func readEnvs() {

	if val := os.Getenv(env.REST_PORT); val != "" {
		env.RestPort = val
	}
	if val := os.Getenv(env.GRPC_PORT); val != "" {
		env.GrpcPort = val
	}
	if val := os.Getenv(env.E3_URL); val != "" {
		env.E3url = val
	}
	if val := os.Getenv(env.E3_DIALET); val != "" {
		env.E3DIALET = val
	}
}

func main() {

	readEnvs()
	env.LoadEnvs()

    dsn := env.E3user + ":" + env.E3pwd + "@tcp(" + env.E3host + ":" + env.E3port + ")/" + env.E3db + "?charset=utf8mb4&parseTime=True&loc=Local"
    sqlDB, err := sql.Open("mysql", dsn)
    database0, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})

    if err != nil {
		xilogger.Log.Error(err)
	}

	env.RDB = database0
    sampleDoctor:= model.Doctor{
        Name:        "Gobbaya",
        Age:         10,
        Nic:         "99008833",
        Address:     "893, hui, ijju",
        Email:       "kowkq@jiew.com",
        Contact:     "099-88836689",
        Nationality: "LK",
    }
    fmt.Println(sampleDoctor)

	model.InitModels(database0)
    arfcreatedoctor.ARF_CreateDoctor(&sampleDoctor)
    arfcreatedoctor.ARF_CreateDoctor(&sampleDoctor)
    api_echo.CreateDoctor2(&sampleDoctor)


    trfmain.TRF_Main()


	go RunServer()
	RunProxy()

}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

var (
	endpoint = flag.String("endpoint", "localhost:50051", "Your Description")
)

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	r := e.Group("")
	controller_echo.APIControllerEcho(r)
	e.Logger.Fatal(e.Start(":" + env.RestPort))
}

func RunServer() {
	lis, err := net.Listen("tcp", ":"+env.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	xilogger.Log.Info("\nServer listening on port %v \n", ":"+env.GrpcPort)
	pub.RegisterPubServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
