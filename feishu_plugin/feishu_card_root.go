package feishu_plugin

import feishuMessageApi "github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_open_api"

type FeishuCardRenderRoot struct {
	// FeishuRobotMsgTemplate
	// use FeishuPlugin root template
	FeishuRobotMsgTemplate feishuMessageApi.FeishuRobotMsgTemplate

	BuildStatus string

	// FeishuCardHeader
	FeishuCardHeader FeishuCardHeader

	// CardElementsHead
	CardElementsHead CardElementsHead

	// CardElementsNoticeList
	CardElementsNoticeList []string

	// CardElementsTail
	CardElementsTail CardElementsTail

	// CardNoteTail
	CardNoteTail CardNoteTail
}

type FeishuCardHeader struct {
	// HeadTemplateStyle
	// headTemplateStyleDefault
	// headTemplateStyleGreen
	// headTemplateStyleRed
	// headTemplateStyleOrange
	// headTemplateStylIndigo
	// @doc https://open.feishu.cn/document/common-capabilities/message-card/message-cards-content/card-header
	HeadTemplateStyle string

	// HeaderTitle
	HeaderTitle string
}

type CardElementsHead struct {
	// CardTitle
	CardTitle string
}

type CardElementsTail struct {
	// BuildTimeCreated
	BuildTimeCreated string

	// BuildTimeTotal
	BuildTimeTotal string

	// RunnerMachine
	RunnerMachine string

	// RunnerPlatform
	RunnerPlatform string

	// BuildEvent
	BuildEvent string
}

type CardNoteTail struct {

	// CiSystemVersion
	CiSystemVersion string

	// CiSystemHost
	CiSystemHost string

	// LogoAltStr
	LogoImgKey string

	// LogoAltStr
	LogoAltStr string
}
