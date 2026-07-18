package tickets

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	gk "github.com/hellicopthecat/catchlot/gakSoo/repo"
	"github.com/hellicopthecat/catchlot/tickets/handler"
	"github.com/hellicopthecat/catchlot/tickets/repo"
	"github.com/hellicopthecat/catchlot/tickets/service"
)

type TicketModule struct {
	TicketHandler handler.TicketHandler
}

func InitTickeModule(db *sql.DB, gkCache *gk.GakSooCache) *TicketModule {
	ticketRepo := repo.InitTicketRepo(db, gkCache)
	ticketService := service.InitTicketService(ticketRepo)
	ticketHandler := handler.InitTicketHandler(ticketService)
	return &TicketModule{
		TicketHandler: *ticketHandler,
	}
}

func (m *TicketModule) TicketGroupApi(r fiber.Router, auth fiber.Handler) {
	t := r.Group("/ticket")
	t.Post("/new", auth, m.TicketHandler.HCreateUserTickets)
}
