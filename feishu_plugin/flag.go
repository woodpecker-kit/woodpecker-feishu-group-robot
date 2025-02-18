package feishu_plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/constant"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"strings"
)

const (
	CliPluginForceStatus = "settings.force-status"
	EnvPluginForceStatus = "PLUGIN_FORCE_STATUS"

	// CliPluginFeishuNoticeTypes
	// feishu_notice_types
	CliPluginFeishuNoticeTypes = "settings.feishu-notice-types"
	EnvPluginFeishuNoticeTypes = "PLUGIN_FEISHU_NOTICE_TYPES"

	CliPluginWebhook = "settings.feishu-webhook"
	EnvPluginWebhook = "PLUGIN_FEISHU_WEBHOOK"

	CliPluginSecret = "settings.feishu-secret"
	EnvPluginSecret = "PLUGIN_FEISHU_SECRET"

	CliPluginFeishuEnableForward = "settings.feishu-enable-forward"
	EnvPluginFeishuEnableForward = "PLUGIN_FEISHU_ENABLE_FORWARD"

	CliPluginFeishuEnableDebugNotice = "settings.feishu-enable-debug-notice"
	EnvPluginFeishuEnableDebugNotice = "PLUGIN_FEISHU_ENABLE_DEBUG_NOTICE"

	CliPluginNtpTarget = "settings.feishu-ntp-target"
	EnvPluginNtpTarget = "PLUGIN_FEISHU_NTP_TARGET"

	CliPluginStatusSuccessIgnore = "settings.feishu-status-success-ignore"
	EnvPluginStatusSuccessIgnore = "PLUGIN_FEISHU_STATUS_SUCCESS_IGNORE"

	CliPluginStatusChangeSuccess = "settings.feishu-status-change-success"
	EnvPluginStatusChangeSuccess = "PLUGIN_FEISHU_STATUS_CHANGE_SUCCESS"

	CliPluginI18nLang = "settings.feishu-msg-i18n-lang"
	EnvPluginI18nLang = "PLUGIN_FEISHU_MSG_I18N_LANG"

	CliPluginTitle = "settings.feishu-msg-title"
	EnvPluginTitle = "PLUGIN_FEISHU_MSG_TITLE"

	CliPluginPoweredByImageKey = "settings.feishu-msg-powered-by-image-key"
	EnvPluginPoweredByImageKey = "PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY"

	CliPluginPoweredByImageAlt = "settings.feishu-msg-powered-by-image-alt"
	EnvPluginPoweredByImageAlt = "PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT"

	CliPluginMsgType = "settings.feishu-msg-type"
	EnvPluginMsgType = "PLUGIN_FEISHU_MSG_TYPE"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    CliPluginForceStatus,
			Usage:   fmt.Sprintf("force status (v1.8+). If empty will use woodpecker ci pipeline status, only support %v", supportRenderStatus),
			Value:   "",
			EnvVars: []string{EnvPluginForceStatus},
		},

		&cli.StringFlag{
			Name:    CliPluginNtpTarget,
			Usage:   "ntp target for sync time like: pool.ntp.org, default not use ntpd to sync",
			EnvVars: []string{EnvPluginNtpTarget},
		},
		&cli.StringFlag{
			Name:    CliPluginWebhook,
			Usage:   "feishu robot webhook, like `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}` end `{web_hook}`",
			EnvVars: []string{EnvPluginWebhook},
		},
		&cli.StringFlag{
			Name:    CliPluginSecret,
			Usage:   "feishu robot secret, just `signature verification`, empty will not open.",
			EnvVars: []string{EnvPluginSecret},
		},
		&cli.BoolFlag{
			Name:    CliPluginFeishuEnableForward,
			Usage:   "feishu robot forward message enable",
			EnvVars: []string{EnvPluginFeishuEnableForward},
		},
		&cli.BoolFlag{
			Name:    CliPluginFeishuEnableDebugNotice,
			Usage:   "when debug open, will not send message, must enable it to notice under debug open",
			EnvVars: []string{EnvPluginFeishuEnableDebugNotice},
		},

		&cli.StringSliceFlag{
			Name:    CliPluginFeishuNoticeTypes,
			Usage:   fmt.Sprintf("feishu notice types, if empty will use [ %s ], now support: %v", NoticeTypeBuildStatus, noticeTypeSupport),
			Value:   cli.NewStringSlice(NoticeTypeBuildStatus),
			EnvVars: []string{EnvPluginFeishuNoticeTypes},
		},

		&cli.StringFlag{
			Name:    CliPluginMsgType,
			Usage:   fmt.Sprintf("feishu message type, now support: %v", supportMsgType),
			Value:   MsgTypeInteractive,
			EnvVars: []string{EnvPluginMsgType},
		},
		&cli.StringFlag{
			Name:    CliPluginTitle,
			Usage:   "feishu message title",
			EnvVars: []string{EnvPluginTitle},
		},
		&cli.StringFlag{
			Name:    CliPluginPoweredByImageKey,
			Usage:   "feishu message powered by image key",
			EnvVars: []string{EnvPluginPoweredByImageKey},
		},
		&cli.StringFlag{
			Name:    CliPluginPoweredByImageAlt,
			Usage:   "feishu message powered by image alt",
			EnvVars: []string{EnvPluginPoweredByImageAlt},
		},

		&cli.BoolFlag{
			Name:    CliPluginStatusSuccessIgnore,
			Usage:   "ignore this build success status",
			EnvVars: []string{EnvPluginStatusSuccessIgnore},
		},
		&cli.BoolFlag{
			Name:    CliPluginStatusChangeSuccess,
			Usage:   fmt.Sprintf("must open [ %s ], when status change to success, compare with %s", CliPluginStatusSuccessIgnore, wd_flag.EnvKeyPreviousCiPipelineStatus),
			EnvVars: []string{EnvPluginStatusChangeSuccess},
		},

		&cli.StringFlag{
			Name: CliPluginI18nLang,
			Usage: fmt.Sprintf("feishu message i18n lang, default %s, support: %s",
				constant.LangEnUS, strings.Join(constant.SupportLanguage(), ", ")),
			Value:   constant.LangEnUS,
			EnvVars: []string{EnvPluginI18nLang},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*FeishuPlugin, error) {

	noticeTypes := c.StringSlice(CliPluginFeishuNoticeTypes)

	if len(noticeTypes) == 0 {
		wd_log.Warnf("notice types is empty, use default notice types %s", NoticeTypeBuildStatus)
		noticeTypes = []string{NoticeTypeBuildStatus}
	}

	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		ForceStatus: c.String(CliPluginForceStatus),

		NtpTarget:           c.String(CliPluginNtpTarget),
		Webhook:             c.String(CliPluginWebhook),
		Secret:              c.String(CliPluginSecret),
		FeishuEnableForward: c.Bool(CliPluginFeishuEnableForward),
		NoticeWhenDebug:     c.Bool(CliPluginFeishuEnableDebugNotice),
		NoticeTypes:         noticeTypes,

		I18nLangSet: c.String(CliPluginI18nLang),

		Title:               c.String(CliPluginTitle),
		StatusSuccessIgnore: c.Bool(CliPluginStatusSuccessIgnore),
		StatusChangeSuccess: c.Bool(CliPluginStatusChangeSuccess),

		MsgType:           c.String(CliPluginMsgType),
		PoweredByImageKey: c.String(CliPluginPoweredByImageKey),
		PoweredByImageAlt: c.String(CliPluginPoweredByImageAlt),
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := FeishuPlugin{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
