package ticket

import (
	"mime/multipart"
	"strconv"

	"github.com/CosmeticsShiraz/Backend/bootstrap"
	ticketdto "github.com/CosmeticsShiraz/Backend/internal/application/dto/ticket"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

type CustomerTicketController struct {
	constants     *bootstrap.Constants
	ticketService usecase.TicketService
	pagination    *bootstrap.Pagination
}

func NewCustomerTicketController(
	constants *bootstrap.Constants,
	ticketService usecase.TicketService,
	pagination *bootstrap.Pagination,
) *CustomerTicketController {
	return &CustomerTicketController{
		constants:     constants,
		ticketService: ticketService,
		pagination:    pagination,
	}
}

func (ticketController *CustomerTicketController) CreateTicket(ctx *gin.Context) {
	type createTicketParams struct {
		Subject     string                `form:"subject" validate:"required"`
		Description string                `form:"description" validate:"required"`
		Image       *multipart.FileHeader `form:"image"`
	}
	params := controller.Validated[createTicketParams](ctx)
	// TODO: what? why? :)
	subject, err := strconv.Atoi(params.Subject)
	if err != nil {
		subject = 2
	}
	userID, _ := ctx.Get(ticketController.constants.Context.ID)
	requestInfo := ticketdto.CreateTicketRequest{
		OwnerID:     userID.(uint),
		OwnerType:   ticketController.constants.TicketOwners.User,
		Subject:     enum.TicketSubject(subject),
		Description: params.Description,
		Image:       params.Image,
	}

	if err := ticketController.ticketService.CreateCustomerTicket(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, ticketController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.createTicket")
	controller.Response(ctx, 200, message, nil)
}

func (ticketController *CustomerTicketController) GetTickets(ctx *gin.Context) {
	type GetTicketsRequest struct {
		Status uint `form:"status" validate:"required"`
	}
	params := controller.Validated[GetTicketsRequest](ctx)
	ownerID, _ := ctx.Get(ticketController.constants.Context.ID)
	pagination := controller.GetPagination(ctx, ticketController.pagination.DefaultPage, ticketController.pagination.DefaultPageSize)
	offset, limit := pagination.GetOffsetLimit()
	listInfo := ticketdto.TicketListRequest{
		OwnerID: ownerID.(uint),
		Status:  params.Status,
		Offset:  offset,
		Limit:   limit,
	}

	tickets, err := ticketController.ticketService.GetCustomerTickets(listInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", tickets)
}

func (ticketController *CustomerTicketController) GetComments(ctx *gin.Context) {
	type getCommentsParams struct {
		TicketID uint `uri:"ticketID" binding:"required"`
	}
	params := controller.Validated[getCommentsParams](ctx)

	pagination := controller.GetPagination(ctx, ticketController.pagination.DefaultPage, ticketController.pagination.DefaultPageSize)
	offset, limit := pagination.GetOffsetLimit()

	ownerID, _ := ctx.Get(ticketController.constants.Context.ID)

	listInfo := ticketdto.TicketCommentListRequest{
		TicketID: params.TicketID,
		OwnerID:  ownerID.(uint),
		Offset:   offset,
		Limit:    limit,
	}
	comments, err := ticketController.ticketService.GetCustomerTicketComments(listInfo)
	if err != nil {
		panic(err)
	}

	controller.Response(ctx, 200, "", comments)
}

func (ticketController *CustomerTicketController) CreateComment(ctx *gin.Context) {
	type createCommentParams struct {
		TicketID uint   `uri:"ticketID" validate:"required"`
		Body     string `json:"body" validate:"required"`
	}
	params := controller.Validated[createCommentParams](ctx)
	userID, _ := ctx.Get(ticketController.constants.Context.ID)

	requestInfo := ticketdto.CreateTicketCommentRequest{
		TicketID:  params.TicketID,
		OwnerID:   userID.(uint),
		OwnerType: ticketController.constants.TicketCommentOwners.User,
		Body:      params.Body,
	}
	if err := ticketController.ticketService.CreateCustomerTicketComment(requestInfo); err != nil {
		panic(err)
	}

	trans := controller.GetTranslator(ctx, ticketController.constants.Context.Translator)
	message, _ := trans.Translate("successMessage.createTicketComment")
	controller.Response(ctx, 200, message, nil)
}
