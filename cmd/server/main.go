package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/hiepdangnguyendai/example-go/config/database/pg"
	"github.com/hiepdangnguyendai/example-go/endpoints"
	serviceHttp "github.com/hiepdangnguyendai/example-go/http"
	"github.com/hiepdangnguyendai/example-go/service"
	bookSvc "github.com/hiepdangnguyendai/example-go/service/book"
	categorySvc "github.com/hiepdangnguyendai/example-go/service/category"
	lendSvc "github.com/hiepdangnguyendai/example-go/service/lend"
)

func main() {
	// setup env on local
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by errors: %v", err))
		}
	}

	// setup addrr
	httpAddr := ":" + os.Getenv("PORT")

	// setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// setup service
	var (
		pgDB, closeDB = pg.New(os.Getenv("PG_DATASOURCE"))
		s             = service.Service{
			CategoryService: service.Compose(
				categorySvc.NewPGService(pgDB),
			//	userSvc.ValidationMiddleware(),
			).(categorySvc.Service),
			BookService: service.Compose(
				bookSvc.NewPGService(pgDB),
			).(bookSvc.Service),
			LendService: service.Compose(
				lendSvc.NewPGService(pgDB),
			).(lendSvc.Service),
		}
	)
	defer closeDB()

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			endpoints.MakeServerEndpoints(s),
			logger,
			os.Getenv("ENV") == "local",
		)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
