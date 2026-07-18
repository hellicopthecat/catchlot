package handler

import "github.com/hellicopthecat/catchlot/official/service"

type OfficialHandler struct {
	handler *service.OfficialLottoService
}

func InitOfficialHandler(handler *service.OfficialLottoService) *OfficialHandler {
	return &OfficialHandler{
		handler: handler,
	}
}
