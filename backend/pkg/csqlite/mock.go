package csqlite

import (
	"chatcat/backend/model"
	"gorm.io/gorm"
	"time"
)

func MockSetting(db *gorm.DB) {
	db.Create(&model.Setting{
		ApiKey:           "",
		AskModel:         "gpt-3.5-turbo",
		ChatModel:        "gpt-3.5-turbo",
		Language:         "zh-CN",
		Theme:            1,
		ProxyUrl:         "",
		Account:          "",
		AccessToken:      "",
		IsSync:           2,
		SyncTime:         5,
		Temperature:      "0.3",
		MaxTokens:        0,
		PresencePenalty:  "0.6",
		FrequencyPenalty: "0.6",
		N:                1,
		TopP:             "0.1",
		BaseModel: model.BaseModel{
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		},
	})
}

// MockChatCate ...
func MockChatCate(db *gorm.DB) {

}

// MockTagList
// @Description: tag list
// @param db
// @author cx
func MockTagList(db *gorm.DB) {

}

// MockPromptList
// @Description: prompt list
// @param db
// @author cx
func MockPromptList(db *gorm.DB) {
	promptList := []model.Prompt{
		{
			Name:     "写作助理",
			Desc:     "👉 最常使用的 prompt，用于优化文本的语法、清晰度和简洁度，提高可读性。",
			Prompt:   "作为一名中文写作改进助理，你的任务是改进所提供文本的拼写、语法、清晰、简洁和整体可读性，同时分解长句，减少重复，并提供改进建议。请只提供文本的更正版本，避免包括解释。请从编辑以下文本开始：[文章内容］",
			Enprompt: "As a writing improvement assistant, your task is to improve the spelling, grammar, clarity, concision, and overall readability of the text provided, while breaking down long sentences, reducing repetition, and providing suggestions for improvement. Please provide only the corrected Chinese version of the text and avoid including explanations. Please begin by editing the following text: [文章内容]",
			TagName:  "写作,文案",
		}, {
			Name:     "保姆",
			Desc:     "👉 Babysitter",
			Prompt:   "我希望你能充当一个保姆。你将负责监督幼儿，准备饭菜和零食，协助做家庭作业和创意项目，参与游戏时间的活动，在需要时提供安慰和安全保障，注意家中的安全问题，并确保所有需求得到照顾。",
			Enprompt: "I want you to act as a babysitter. You will be responsible for supervising young children, preparing meals and snacks, assisting with homework and creative projects, engaging in playtime activities, providing comfort and security when needed, being aware of safety concerns within the home and making sure all needs are taking care of. My first suggestion request is [照顾对象]",
			TagName:  "生活",
		}, {
			Name:     "健身教练",
			Desc:     "👉 通过输入身高、体重、年龄等指标，来制定健身方案",
			Prompt:   "我希望你能充当私人教练。我将为你提供一个希望通过体能训练变得更健康、更强壮、更健康的人所需要的所有信息，而你的职责是根据这个人目前的体能水平、目标和生活习惯，为其制定最佳计划。你应该运用你的运动科学知识、营养建议和其他相关因素，以便制定出适合他们的计划。",
			Enprompt: "I want you to act as a personal trainer. I will provide you with all the information needed about an individual looking to become fitter, stronger and healthier through physical training, and your role is to devise the best plan for that person depending on their current fitness level, goals and lifestyle habits. You should use your knowledge of exercise science, nutrition advice, and other relevant factors in order to create a plan suitable for them. My first request is [健身目的]",
			TagName:  "生活",
		}, {
			Name:     "翻译成中文和润色",
			Desc:     "把任何语言翻译成中文",
			Prompt:   "将我输入的任何语言翻译成中文，如果我输入的是中文帮我润色一下",
			Enprompt: "Translate any language I input into Chinese, and if I input Chinese, please help me polish it.",
			TagName:  "生活",
		}, {
			Name:     "内容总结",
			Desc:     "👉 将文本内容总结为 100 字。",
			Prompt:   "将以下文字概括为 100 个字，使其易于阅读和理解。避免使用复杂的句子结构或技术术语。",
			Enprompt: "Summarize the following text into 100 words, making it easy to read and comprehend. The summary should be concise, clear, and capture the main points of the text. Avoid using complex sentence structures or technical jargon. Please begin by editing the following text: ",
			TagName:  "创作",
		}, {
			Name:     "写作标题生成器",
			Desc:     "👉 个人使用的提示词，可根据文章内容生成相应语言的标题。",
			Prompt:   "我想让你充当书面作品的标题生成器。我将向你提供一篇文章的主题和关键词，你将生成五个吸引人的标题。请保持标题简洁，不超过 20 个字，并确保保持其含义。答复时要利用题目的语言类型。我的第一个题目是 [文章内容]",
			Enprompt: "I want you to act as a title generator for written pieces. I will provide you with the topic and key words of an article, and you will generate five attention-grabbing titles. Please keep the title concise and under 20 words, and ensure that the meaning is maintained. Replies will utilize the language type of the topic. My first topic is [文章内容]",
			TagName:  "写作",
		}, {
			Name:     "开发：Vue3",
			Desc:     "👉 辅助 Vue3 开发",
			Prompt:   "作为 Vue3 开发人员，你的任务是使用 Yarn、Vite、Vue3、TS、Pinia 和 Vueuse 工具编写一个计数器。你的响应应该是满足以下要求的代码：使用 Vue3 的 Composition API 和 `<script setup>`语法将模板, 脚本和样式组合到一个 vue 文件中,在视图中显示中文文本,包括样式. 请注意, 您应该只提供满足这些要求所必需的代码; 不需要解释或描述.",
			Enprompt: "Create a Vue 3 component that displays a [开发项目] using Yarn, Vite, Vue 3, TypeScript, Pinia, and Vueuse tools. Use Vue 3\"s Composition API and <script setup> syntax to combine template, script, and style in a single .vue file. Display Chinese text in the view and include styles. Provide only the necessary code to meet these requirements without explanations or descriptions.",
			TagName:  "开发",
		}, {
			Name:     "正则生成器",
			Desc:     "👉 根据要求生成正则表达式",
			Prompt:   "我希望你充当一个正则表达式生成器。你的角色是生成匹配文本中特定模式的正则表达式。你应该提供正则表达式的格式，以便于复制和粘贴到支持正则表达式的文本编辑器或编程语言中。不要写关于正则表达式如何工作的解释或例子；只需提供正则表达式本身。我的第一个提示是生成一个匹配 [正则要求] 的正则表达式。",
			Enprompt: "I want you to act as a regex generator. Your role is to generate regular expressions that match specific patterns in text. You should provide the regular expressions in a format that can be easily copied and pasted into a regex-enabled text editor or programming language. Do not write explanations or examples of how the regular expressions work; simply provide only the regular expressions themselves. My first prompt is to generate a regular expression that matches [正则要求]",
			TagName:  "开发",
		},
	}
	db.Create(&promptList)
}
