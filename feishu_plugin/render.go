package feishu_plugin

const (
	RenderStatusShow = "success"
	RenderStatusHide = "failure"

	MsgTypeText        = "text"
	MsgTypePost        = "post"
	MsgTypeInteractive = "interactive"
)

var (
	// supportMsgType
	supportRenderStatus = []string{
		RenderStatusShow,
		RenderStatusHide,
	}

	// supportMsgType
	// @doc https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=zh-CN#8b0f2a1b
	supportMsgType = []string{
		MsgTypeInteractive,
	}
)
