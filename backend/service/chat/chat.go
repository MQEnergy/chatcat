package chat

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/cgpt"
	"chatcat/backend/pkg/cpaginator"
	"chatcat/backend/pkg/cresp"
	"chatcat/backend/service"
	"chatcat/backend/service/setting"
	"github.com/mozillazg/go-pinyin"
	"github.com/sashabaranov/go-openai"
	"strings"
)

type Service struct {
	App *service.App
}

func New(a *service.App) *Service {
	return &Service{
		App: a,
	}
}

var (
	GPTPkg *cgpt.GPT
)

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
	if data.Name == "" {
		return cresp.Fail("cate name is required")
	}
	// 获取letter
	lazyPinyin := pinyin.LazyPinyin(data.Name, pinyin.NewArgs())
	data.Letter = strings.ToUpper(lazyPinyin[0][0:1])
	data.Color = "#3370ff" // Todo
	if err := s.App.DB.Save(&data).Error; err != nil {
		return cresp.Fail("chat cate save failed")
	}
	return cresp.Success(data)
}

// SetChatData
// @Description: set chat data
// @receiver s
// @param data
// @return *cresp.Response
// @author cx
func (s *Service) SetChatData(data model.Chat) *cresp.Response {
	var chatData model.Chat
	if err := s.App.DB.First(&chatData, "name = ? and cate_id = ?", data.Name, data.CateId).Error; err == nil {
		return cresp.Fail("chat is already existed")
	}
	if data.Name == "" {
		return cresp.Fail("chat name is required")
	}
	if err := s.App.DB.Save(&data).Error; err != nil {
		return cresp.Fail("chat save failed")
	}
	return cresp.Success(data)
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

// DeleteChat
// @Description: delete chat record
// @receiver s
// @param id
// @return *cresp.Response
// @author cx
func (s *Service) DeleteChat(id uint) *cresp.Response {
	var chatData model.Chat
	if err := s.App.DB.First(&chatData, "id = ?", id).Error; err != nil {
		return cresp.Fail("record is already deleted")
	}
	if err := s.App.DB.Delete(&chatData, "id = ?", id).Error; err != nil {
		return cresp.Fail("delete record failed")
	}
	return cresp.Success(chatData)
}

// EditChat
// @Description: edit chat record
// @receiver s
// @param id
// @param name
// @return *cresp.Response
// @author cx
func (s *Service) EditChat(id uint, name string) *cresp.Response {
	var chatData model.Chat
	if err := s.App.DB.First(&chatData, "id = ?", id).Error; err != nil {
		return cresp.Fail("record is not existed")
	}
	if err := s.App.DB.Model(&model.Chat{}).
		Where("id = ?", id).Update("name", name).Error; err != nil {
		return cresp.Fail("update chat record failed")
	}
	return cresp.Success(chatData)
}

// CompletionStream
// @Description: 问答类型数据流 适用于：text-davinci-003, text-davinci-002, text-curie-001, text-babbage-001, text-ada-001, davinci, curie, babbage, ada
// @receiver s
// @param prompt
// @return error
// @author cx
func (s *Service) CompletionStream(prompt string) *cresp.Response {
	generalInfo := setting.New(s.App).GetGeneralInfo()
	data := generalInfo.Data.(model.Setting)
	if data.ApiKey == "" {
		return cresp.Fail("Chatcat Warm Reminder: You didn't provide an API key. You need to provide your API key in an Authorization header using Bearer auth (i.e. Authorization: Bearer YOUR_KEY), or as the password field (with blank username) if you're accessing the API from your browser and are prompted for a username and password. You can obtain an API key from https://platform.openai.com/account/api-keys.")
	}
	GPTPkg = cgpt.New(data.ApiKey, s.App)
	GPTPkg.WithProxy(data.ProxyUrl).
		WithModel(data.AskModel).
		WithPrompt(prompt).
		WithMaxTokens(0).
		WithCompletionRequest().
		CompletionStream()

	return cresp.Success("")
}

// ChatCompletionStream
// @Description: 对话类型数据流 主流 适用于：gpt-4, gpt-4-0314, gpt-4-32k, gpt-4-32k-0314, gpt-3.5-turbo, gpt-3.5-turbo-0301
// @receiver s
// @param prompt
// @return error
// @author cx
func (s *Service) ChatCompletionStream(messages []openai.ChatCompletionMessage, clientId string) *cresp.Response {
	generalInfo := setting.New(s.App).GetGeneralInfo()
	data := generalInfo.Data.(model.Setting)
	if data.ApiKey == "" {
		return cresp.Fail("Chatcat Warm Reminder: You didn't provide an API key. You need to provide your API key in an Authorization header using Bearer auth (i.e. Authorization: Bearer YOUR_KEY), or as the password field (with blank username) if you're accessing the API from your browser and are prompted for a username and password. You can obtain an API key from https://platform.openai.com/account/api-keys.")
	}
	GPTPkg = cgpt.New(data.ApiKey, s.App)
	gpt := GPTPkg.WithProxy(data.ProxyUrl).
		WithModel(data.ChatModel).
		WithMessages(messages).
		WithMaxTokens(0)
	if cgpt.MaxTokens <= 0 {
		return cresp.Fail("Chatcat Warm Reminder: Your token is running low, Please start a new conversation.")
	}
	gpt.WithChatCompletionRequest().
		ChatCompletionStream(clientId)

	return cresp.Success("")
}

// GetWsUrl ...
func (s *Service) GetWsUrl() string {
	return s.App.Cfg.App.WsUrl
}

// BreakOffChatStream
// @Description: break off chat stream
// @receiver s
// @author cx
func (s *Service) BreakOffChatStream() {
	if GPTPkg != nil {
		GPTPkg.ChatCompStream.Close()
	}
}

// GetTokensNumFromMessages
// @Description: get token number from messages
// @receiver s
// @param messages
// @return *cresp.Response
// @author cx
func (s *Service) GetTokensNumFromMessages(data model.Setting, messages []openai.ChatCompletionMessage) *cresp.Response {
	if data.ApiKey == "" {
		return cresp.Success(0)
	}
	gpt := cgpt.New(data.ApiKey, s.App).
		WithProxy(data.ProxyUrl).
		WithModel(data.ChatModel).
		WithMessages(messages).
		WithMaxTokens(0)
	if cgpt.MaxTokens <= 0 {
		return cresp.Success(0)
	}
	return cresp.Success(gpt.TikToken)
}
