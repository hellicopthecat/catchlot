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

func InitTickeModule(db *sql.DB) *TicketModule {
	ticketRepo := repo.InitTicketRepo(db)
	ticketService := service.InitTicketService(ticketRepo)
	ticketHander := handler.InitTicketHandler(ticketService)
	return &TicketModule{
		TicketHandler: *ticketHander,
	}
}

func (m *TicketModule) TicketGroupApi(r fiber.Router, auth fiber.Handler, cache *gk.GakSooCache) {
	t := r.Group("/ticket")
	t.Post("/new", auth, m.TicketHandler.HCreateUserTicket)
}
