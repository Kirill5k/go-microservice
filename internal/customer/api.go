package customer

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/common/errors"
	"kirill5k/go/microservice/internal/server"
	"net/http"
)

type Api struct {
	service *Service
}

func (hc *Api) RegisterRoutes(server server.Server) {
	parseId := func(ctx echo.Context) (uuid.UUID, error) {
		idString := ctx.Param("id")
		id, err := uuid.Parse(idString)
		if err != nil {
			return id, &errors.InvalidIdError{ID: idString}
		}
		return id, nil
	}

	server.PrefixRoute("/customers")

	server.AddRoute("GET", "", func(ctx echo.Context) error {
		email := ctx.QueryParam("email")
		customers, err := hc.service.FindBy(ctx.Request().Context(), email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, customers)
	})

	server.AddRoute("GET", "/:id", func(ctx echo.Context) error {
		id, err := parseId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		customer, err := hc.service.Get(ctx.Request().Context(), id)
		if err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				return ctx.JSON(http.StatusNotFound, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ctx.JSON(http.StatusOK, customer)
	})

	server.AddRoute("POST", "", func(ctx echo.Context) error {
		newCust := new(NewCustomer)
		if err := ctx.Bind(newCust); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		cust, err := hc.service.Create(ctx.Request().Context(), newCust)
		if err != nil {
			switch err.(type) {
			case *errors.ConflictError:
				return ctx.JSON(http.StatusConflict, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ctx.JSON(http.StatusCreated, cust)
	})

	server.AddRoute("PUT", "/:id", func(ctx echo.Context) error {
		id, err := parseId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		cust := new(Customer)
		if err := ctx.Bind(cust); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		if cust.ID != id {
			return ctx.JSON(http.StatusBadRequest, errors.IdMissmatchError{BodyID: cust.ID, PathID: id})
		}

		res, err := hc.service.Update(ctx.Request().Context(), cust)
		if err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				return ctx.JSON(http.StatusNotFound, err)
			case *errors.ConflictError:
				return ctx.JSON(http.StatusConflict, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}

		return ctx.JSON(http.StatusOK, res)
	})

	deleteById := func(ctx echo.Context) error {
		id, err := parseId(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		if err := hc.service.Delete(ctx.Request().Context(), id); err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				return ctx.JSON(http.StatusNotFound, err)
			default:
				return ctx.JSON(http.StatusInternalServerError, err)
			}
		}
		return ctx.NoContent(http.StatusResetContent)
	}
	server.AddRoute("DELETE", "/:id", deleteById)
}

func NewApi(service *Service) *Api {
	return &Api{service}
}
