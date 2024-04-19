package feishu_plugin

type (

	// SendTarget send feishu target
	SendTarget struct {
		Webhook        string
		Secret         string
		FeishuRobotMeg []byte
	}

	// Settings feishu_plugin private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string
		DryRun            bool

		NtpTarget string

		Webhook             string
		Secret              string
		FeishuEnableForward bool
		NoticeWhenDebug     bool
		NoticeTypes         []string

		// set of i18n
		I18nLangSet string

		MsgType           string
		Title             string
		PoweredByImageKey string
		PoweredByImageAlt string

		// StatusSuccessIgnore
		// for NoticeTypeBuildStatus
		StatusSuccessIgnore bool

		// StatusChangeSuccess
		// for NoticeTypeBuildStatus
		// must open StatusSuccessIgnore, when status change to success, compare with wd_flag.EnvKeyPreviousCiPipelineStatus
		StatusChangeSuccess bool
	}
)
