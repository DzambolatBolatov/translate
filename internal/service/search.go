package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (s *Service) SearchRu(c echo.Context) error {
	title, err := strconv.Atoi(c.Param("title"))
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.JSON(http.StatusOK, search)

}
