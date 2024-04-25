package service

import (
	"dictionary/internal/reports"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (s *Service) CreateReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := reports.NewRepo(s.db)
	err = repo.CreateNewReport(report.WordId, report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (s *Service) GetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}
	repo := reports.NewRepo(s.db)
	report, err := repo.GetReport(id)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.JSON(http.StatusOK, report)
}

func (s *Service) UpdateReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := reports.NewRepo(s.db)
	err = repo.UpdateReport(report.Id, report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (s *Service) DeleteReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}
	repo := reports.NewRepo(s.db)
	err = repo.DeleteReport(report.Id)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
