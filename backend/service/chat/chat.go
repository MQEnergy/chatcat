package chat

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/cgpt"
	"chatcat/backend/pkg/cpaginator"
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/service"
	"chatcat/backend/service/setting"
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	App *service.App
}

func New(a *service.App) *Service {
	return &Service{
		App: a,
	}
}

// GetChatCateList
// @Description: get chat categories list
// @receiver s
// @param page
// @return *cresp.Response
// @author cx
func (s *Service) GetChatCateList(page int) *cresp.Response {
	var chatCateList = make([]model.ChatCate, 0)
	pagination, err := cpaginator.NewBuilder(s.App.DB, s.App.Cfg).
		WithModel(&model.ChatCate{}).
		WithOrderBy("id DESC").
		Pagination(&chatCateList, page, s.App.Cfg.App.DefaultPageSize)
	if err != nil {
		return cresp.Fail("GetChatCateList error:" + err.Error())
	}
	return cresp.Success(pagination)
}

// GetChatList
// @Description: get chat list
// @receiver s
// @param page
// @return *cresp.Response
// @author cx
func (s *Service) GetChatList(cateId uint, page int) *cresp.Response {
	var chatList = make([]model.Chat, 0)
	pagination, err := cpaginator.NewBuilder(s.App.DB, s.App.Cfg).
		WithModel(&model.Chat{}).
		WithCondition("cate_id = ?", cateId).
		WithOrderBy("id DESC").
		Pagination(&chatList, page, s.App.Cfg.App.DefaultPageSize)
	if err != nil {
		return cresp.Fail("GetChatList error:" + err.Error())
	}
	return cresp.Success(pagination)
}

// GetChatRecordList
// @Description: get chat record list
// @receiver s
// @param chatId
// @param page
// @return *cresp.Response
// @author cx
func (s *Service) GetChatRecordList(chatId uint, page int) *cresp.Response {
	var chatRecordList = make([]model.ChatRecord, 0)
	pagination, err := cpaginator.NewBuilder(s.App.DB, s.App.Cfg).
		WithModel(&model.ChatRecord{}).
		WithCondition("chat_id = ?", chatId).
		WithOrderBy("id DESC").
		Pagination(&chatRecordList, page, s.App.Cfg.App.DefaultPageSize)
	if err != nil {
		return cresp.Fail("GetChatRecordList error:" + err.Error())
	}
	return cresp.Success(pagination)
}

// SetChatCateData
// @Description: set chat category
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetChatCateData(data model.ChatCate) *cresp.Response {
	var chatCateInfo model.ChatCate
	if err := s.App.DB.First(&chatCateInfo, "name = ?", data.Name).Error; err == nil {
		return cresp.Fail("chat cate is already existed")
	}
	if err := s.App.DB.Save(data).Error; err != nil {
		return cresp.Fail("chat cate save failed")
	}
	return cresp.Success(chatCateInfo)
}

// SetChatData
// @Description: set chat data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetChatData(data model.Chat) *cresp.Response {
	var chatData model.Chat
	if err := s.App.DB.First(&chatData, "name = ?", data.Name).Error; err == nil {
		return cresp.Fail("chat is already existed")
	}
	if err := s.App.DB.Save(data).Error; err != nil {
		return cresp.Fail("chat save failed")
	}
	return cresp.Success(chatData)
}

// SetChatRecordData
// @Description: set chat record data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetChatRecordData(data model.ChatRecord) *cresp.Response {
	var chatRecordInfo model.ChatRecord
	if err := s.App.DB.First(&chatRecordInfo, "name = ?", data.Name).Error; err == nil {
		return cresp.Fail("chat record is already existed")
	}
	if err := s.App.DB.Save(data).Error; err != nil {
		return cresp.Fail("chat record save failed")
	}
	return cresp.Success(chatRecordInfo)
}

// CompletionStream
// @Description: 问答类型数据流
// @receiver s
// @param prompt
// @author cx
func (s *Service) CompletionStream(prompt string) {
	generalInfo := setting.New(s.App).GetGeneralInfo()
	data := generalInfo.Data.(model.Setting)
	cgpt.New(data.ApiKey, s.App).
		WithProxy(data.ProxyUrl).
		WithModel(data.AskModel).
		WithPrompt(prompt).
		WithMaxTokens(0).
		WithCompletionRequest().
		CompletionStream()
}

func (s *Service) ChatCompletionStream(prompt string) {
	generalInfo := setting.New(s.App).GetGeneralInfo()
	data := generalInfo.Data.(model.Setting)
	cgpt.New(data.ApiKey, s.App).
		WithProxy(data.ProxyUrl).
		WithModel(data.ChatModel).
		WithPrompt(prompt).
		WithMessages([]openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		}).
		WithMaxTokens(0).
		WithChatCompletionRequest().
		ChatCompletionStream()
}

func (s *Service) GetWsUrl() string {
	return s.App.Cfg.App.WsUrl
}
