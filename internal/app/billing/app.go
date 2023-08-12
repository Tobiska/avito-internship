package billing

import (
	"fmt"
	"github.com/Tobiska/avito-internship/config"
	"github.com/Tobiska/avito-internship/internal/handlers"
	"github.com/Tobiska/avito-internship/internal/handlers/billing"
)

func Run(cfg *config.Config) {
	router := handlers.NewRouter()
	handler := billing.NewHandler()
	handler.Register(router)
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", cfg.ApiPort)))
}
