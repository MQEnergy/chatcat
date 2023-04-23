package prompt

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/cpaginator"
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

// GetPromptList
// @Description: get list of prompts
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) GetPromptList(page int) *cresp.Response {
	var promptList = make([]model.Prompt, 0)
	pagination, err := cpaginator.NewBuilder(s.App.DB, s.App.Cfg).
		WithModel(&model.Prompt{}).
		WithOrderBy("id DESC").
		Pagination(&promptList, page, s.App.Cfg.App.DefaultPageSize)
	if err != nil {
		return cresp.Fail("GetChatCateList error:" + err.Error())
	}
	return cresp.Success(pagination)
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

// DeletePrompt
// @Description: delete prompt
// @receiver s
// @param id
// @return *cresp.Response
// @author cx
func (s *Service) DeletePrompt(id int) *cresp.Response {
	var promptInfo model.Prompt
	if err := s.App.DB.First(&promptInfo, "id = ?", id).Error; err != nil {
		return cresp.Fail("prompt is not existed")
	}
	if err := s.App.DB.Delete(&promptInfo).Error; err != nil {
		return cresp.Fail("prompt delete failed")
	}
	return cresp.Success("")
}
