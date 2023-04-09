package setting

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/service"
)

type Service struct {
	App *service.App
}

func New(a *service.App) *Service {
	return &Service{
		App: a,
	}
}

// GetGeneralInfo
// @Description: get general info
// @receiver s
// @return *cresp.Response
// @author cx
func (s *Service) GetGeneralInfo() *cresp.Response {
	var settingInfo model.Setting
	if err := s.App.DB.First(&settingInfo).Error; err != nil {
		return cresp.Success("")
	}
	return cresp.Success(settingInfo)
}

// SetGeneralData
// @Description: set general data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetGeneralData(data model.Setting) *cresp.Response {
	if err := s.App.DB.Save(&data).Error; err != nil {
		return cresp.Fail("save failed err:" + err.Error())
	}
	return cresp.Success("")
}
