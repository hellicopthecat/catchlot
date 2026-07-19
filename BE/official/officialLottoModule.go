package official

import (
	"context"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v3"
	initGakSooStatusRepo "github.com/hellicopthecat/catchlot/gakSooStatus/repo"
	"github.com/hellicopthecat/catchlot/official/handler"
	"github.com/hellicopthecat/catchlot/official/repo"
	"github.com/hellicopthecat/catchlot/official/service"
)

type OfficialLottoModule struct {
	officialHandler *handler.OfficialHandler
}

func InitOfficialLottoModules(db *sql.DB, ctx context.Context) *OfficialLottoModule {
	officialLottoRepo := repo.InitOfficialLottoRepo(db)
	gakSooStatusRepo := initGakSooStatusRepo.InitGakSooStatus(db)

	officialLottoService := service.InitOfficialLottoService(officialLottoRepo, gakSooStatusRepo)
	err := officialLottoService.SInitAllRoundsLotto(ctx)
	if err != nil {
		log.Printf("❌ SInitAllRoundsLotto :: %s", err)
	}
	officialLottoService.SCronLottoRound(ctx)
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
