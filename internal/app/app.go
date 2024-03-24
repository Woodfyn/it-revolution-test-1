package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/Woodfyn/it-revolution-test-1/internal/config"
	"github.com/Woodfyn/it-revolution-test-1/internal/repository/mongo"
	"github.com/Woodfyn/it-revolution-test-1/internal/service"
	"github.com/Woodfyn/it-revolution-test-1/internal/transport/rest"
	"github.com/Woodfyn/it-revolution-test-1/pkg/mdb"
	"github.com/Woodfyn/it-revolution-test-1/pkg/signaler"
	"github.com/Woodfyn/it-revolution-test-1/pkg/srv"
)

const (
	cfg_folder = "configs"
	cfg_file   = "prod.yml"
)

func Init() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(h))
}

func Run() {
	//init config
	cfg, err := config.InitConfig(cfg_folder, cfg_file)
	if err != nil {
		slog.Error("failed to init config", err)
		panic(err)
	}

	slog.Info("config loaded...", "config", cfg)

	//init mongo
	dbClient, err := mdb.NewMongoClient(context.Background(), mdb.ConnInfo{
		URI:      cfg.Mongo.URI,
		Username: cfg.Mongo.Username,
		Password: cfg.Mongo.Password,
	})
	if err != nil {
		slog.Error("failed to init mongo client", err)
		panic(err)
	}
	db := dbClient.Database(cfg.Mongo.Database)
	defer dbClient.Disconnect(context.Background())

	slog.Info("mongo connected...")

	//init dependencies
	repository := mongo.NewRepository(db)
	service := service.NewService(service.Deps{
		MongoRepo: repository,
	})
	handler := rest.NewHandler(service)

	//init server
	srv := new(srv.Server)

	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		slog.Error("failed to run server", err)
		panic(err)
	}

	slog.Info("server started...")

	//graceful shutdown
	signaler.Wait()

	slog.Info("server stopped...")

	//shutdown server
	srv.Shutdown(context.Background())
}
