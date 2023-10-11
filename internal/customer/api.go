package customer

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/common"
	"kirill5k/go/microservice/internal/server"
	"net/http"
)

type Api struct {
	service *Service
}

func (hc *Api) RegisterRoutes(server server.Server) {
	server.PrefixRoute("/customers")

	getAll := func(ctx echo.Context) error {
		email := ctx.QueryParam("email")
		customers, err := hc.service.FindBy(ctx.Request().Context(), email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, customers)
	}
	server.AddRoute("GET", "", getAll)

	getById := func(ctx echo.Context) error {
		idString := ctx.Param("id")
		id, err := uuid.Parse(ctx.Param(idString))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, common.InvalidIdError{ID: idString})
		}
		customer, err := hc.service.Get(ctx.Request().Context(), id)
		if err != nil {
			switch err.(type) {
			case *common.NotFoundError:
				return ctx.JSON(http.StatusNotFound, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ctx.JSON(http.StatusOK, customer)
	}
	server.AddRoute("GET", "/:id", getById)

	createNew := func(ctx echo.Context) error {
		newCust := new(NewCustomer)
		if err := ctx.Bind(newCust); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		cust, err := hc.service.Create(ctx.Request().Context(), newCust)
		if err != nil {
			switch err.(type) {
			case *common.ConflictError:
				return ctx.JSON(http.StatusConflict, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ctx.JSON(http.StatusCreated, cust)
	}
	server.AddRoute("POST", "", createNew)
}

func NewApi(service *Service) *Api {
	return &Api{service}
}
