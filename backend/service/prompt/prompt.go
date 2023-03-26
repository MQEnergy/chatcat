package prompt

import "chatcat/backend/service"

type Service struct {
	App *service.App
}

func New(app *service.App) *Service {
	return &Service{
		App: app,
	}
}
func (s *Service) Test() {

}
