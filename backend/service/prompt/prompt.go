package prompt

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/service"
)

type Service struct {
	App *service.App
}

func New(app *service.App) *Service {
	return &Service{
		App: app,
	}
}

// GetList
// @Description: get list of prompts
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) GetList() *cresp.Response {
	var promptList = make([]model.Prompt, 0)
	if err := s.App.DB.Find(&promptList).Error; err != nil {
		return cresp.Fail("prompt list is empty")
	}
	return cresp.Success(promptList)
}

// SetPromptData
// @Description: set prompt data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetPromptData(data model.Prompt) *cresp.Response {
	var promptInfo model.Prompt
	if err := s.App.DB.First(&promptInfo, "name = ?", data.Name).Error; err == nil {
		return cresp.Fail("prompt name is existed")
	}
	if err := s.App.DB.Save(&data).Error; err != nil {
		return cresp.Fail("prompt save failed")
	}
	return cresp.Success("")
}

// UpdatePromptData
// @Description: update prompt data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) UpdatePromptData(data model.Prompt) *cresp.Response {
	var promptInfo model.Prompt
	if err := s.App.DB.First(&promptInfo, "id = ?", data.ID).Error; err != nil {
		return cresp.Fail("prompt is not existed")
	}
	if err := s.App.DB.Save(&data).Error; err != nil {
		return cresp.Fail("prompt save failed")
	}
	return cresp.Success("")
}
