package main

import (
	h "HexMicroservice/api"
	mr "HexMicroservice/repository/mongodb"
	my "HexMicroservice/repository/mysql"
	rr "HexMicroservice/repository/redis"
	"HexMicroservice/shortener"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	REDIS string = "redis"
	MONGO string = "mongo"
	MYSQL string = "mysql"
)

func httpPost() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case REDIS:
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case MONGO:
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case MYSQL:
		mysqURL := os.Getenv("MYSQL_URL")
		mysqlDBName := os.Getenv("MYSQL_DBNAME")
		repo, err := my.NewMysqlRepository(mysqURL, mysqlDBName)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port:8000")
		errs <- http.ListenAndServe(httpPost(), r)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminal %s", <-errs)
}
