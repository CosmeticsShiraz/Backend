package ticket

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type GeneralTicketController struct {
	constants     *bootstrap.Constants
	ticketService usecase.TicketService
	pagination    *bootstrap.Pagination
}

func NewGeneralTicketController(
	constants *bootstrap.Constants,
	ticketService usecase.TicketService,
	pagination *bootstrap.Pagination,
) *GeneralTicketController {
	return &GeneralTicketController{
		constants:     constants,
		ticketService: ticketService,
		pagination:    pagination,
	}
}

func (ticketController *GeneralTicketController) GetTicketStatuses(ctx *gin.Context) {
	ticketStatuses := ticketController.ticketService.GetTicketStatuses()
	controller.Response(ctx, 200, "", ticketStatuses)
}
