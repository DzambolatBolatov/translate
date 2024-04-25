package service

import (
	"net/http"
	"strconv"

	"dictionary/internal/words"
	"github.com/labstack/echo/v4"
)

// GetWordById идем слово по id
// localhost:8000/api/word/:id
func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.JSON(http.StatusOK, word)
}

// CreateWords добавляем в базу новые слова в базу
// localhost:8000/api/words
func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	for _, word := range wordSlice {
		err = repo.CreateNewWords(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// надо добавить изменение перевода из базы
// localhost:8000/api/words
func (s *Service) UpdateWordsById(c echo.Context) error {
	var word Word
	err := c.Bind(&word)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	err = repo.UpdateWordById(word.Id, word.Title, word.Translation)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// Удаление перевода из базы
func (s *Service) DeleteWordById(c echo.Context) error {
	var word Word
	err := c.Bind(&word)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	err = repo.DeleteWordById(word.Id)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)

}
