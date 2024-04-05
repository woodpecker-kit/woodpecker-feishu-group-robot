package feishu_plugin

const (
	// NoticeTypeBuildStatus
	// build status notice, default notice type
	NoticeTypeBuildStatus = "build_status"

	// NoticeTypeFileBrowser
	// file browser notice
	NoticeTypeFileBrowser = "file_browser"

	RenderStatusShow = "success"
	RenderStatusHide = "failure"

	MsgTypeText        = "text"
	MsgTypePost        = "post"
	MsgTypeInteractive = "interactive"
)

var (
	noticeTypeSupport = []string{
		NoticeTypeBuildStatus,
		NoticeTypeFileBrowser,
	}

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
