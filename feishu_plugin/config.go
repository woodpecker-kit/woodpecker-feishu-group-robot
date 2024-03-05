package feishu_plugin

type (

	// SendTarget send feishu target
	SendTarget struct {
		Webhook        string
		Secret         string
		FeishuRobotMeg []byte
	}

	// Config feishu_plugin private config
	Config struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		RootPath          string
		DryRun            bool

		NtpTarget string

		Webhook             string
		Secret              string
		FeishuEnableForward bool
		MsgType             string
		Title               string
		PoweredByImageKey   string
		PoweredByImageAlt   string

		StatusSuccessIgnore bool
		StatusChangeSuccess bool

		RenderOssCard string
		CardOss       CardOss
	}

	CardOss struct {
		Host string

		// HeadTemplateStyle
		// @doc https://open.feishu.cn/document/common-capabilities/message-card/message-cards-content/card-header
		HeadTemplateStyle string

		// InfoTagResult
		// tag result [ success or failure]
		InfoTagResult string
		// InfoSendResult
		// send result [ success or failure]
		InfoSendResult string

		// pull request [ success or failure]
		InfoPullRequestResult string

		InfoUser string
		InfoPath string

		RenderResourceUrl string
		ResourceUrl       string
		PageUrl           string
		PagePasswd        string
	}
)
