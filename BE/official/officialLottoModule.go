package official

import (
	"context"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/official/handler"
	"github.com/hellicopthecat/catchlot/official/repo"
	"github.com/hellicopthecat/catchlot/official/service"
)

type OfficialLottoModule struct {
	officialHandler *handler.OfficialHandler
}

func InitOfficialLottoModules(db *sql.DB, ctx context.Context) *OfficialLottoModule {
	officialLottoRepo := repo.InitOfficialLottoRepo(db)

	officialLottoService := service.InitOfficialLottoService(officialLottoRepo)
	err := officialLottoService.SInitAllRoundsLotto(ctx)
	if err != nil {
		log.Printf("❌ SInitAllRoundsLotto :: %s", err)
	}
	officialLottoService.SCronLottoRound()
	officialHandler := handler.InitOfficialHandler(officialLottoService)

	return &OfficialLottoModule{
		officialHandler: officialHandler,
	}
}

func (*OfficialLottoModule) OfficialLottoApi(r fiber.Router) {
	// g := r.Group("/official")
	// g.Post("/init")
	// g.Post("/init-manual",)
	// g.Post("/init-cron",)
}
