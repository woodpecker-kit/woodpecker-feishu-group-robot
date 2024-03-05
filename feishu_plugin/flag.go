package feishu_plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
)

const (
	CliPluginNtpTarget = "settings.ntp_target"
	EnvPluginNtpTarget = "PLUGIN_NTP_TARGET"

	CliPluginWebhook = "settings.webhook"
	EnvPluginWebhook = "PLUGIN_WEBHOOK"

	CliPluginSecret = "settings.secret"
	EnvPluginSecret = "PLUGIN_SECRET"

	CliPluginFeishuEnableForward = "settings.feishu_enable_forward"
	EnvPluginFeishuEnableForward = "PLUGIN_FEISHU_ENABLE_FORWARD"

	CliPluginMsgType = "settings.msg_type"
	EnvPluginMsgType = "PLUGIN_MSG_TYPE"

	CliPluginTitle = "settings.msg_title"
	EnvPluginTitle = "PLUGIN_MSG_TITLE"

	CliPluginPoweredByImageKey = "settings.msg_powered_by_image_key"
	EnvPluginPoweredByImageKey = "PLUGIN_MSG_POWERED_BY_IMAGE_KEY"

	CliPluginPoweredByImageAlt = "settings.msg_powered_by_image_alt"
	EnvPluginPoweredByImageAlt = "PLUGIN_MSG_POWERED_BY_IMAGE_ALT"

	CliPluginStatusSuccessIgnore = "settings.status_success_ignore"
	EnvPluginStatusSuccessIgnore = "PLUGIN_STATUS_SUCCESS_IGNORE"

	CliPluginStatusChangeSuccess = "settings.status_change_success"
	EnvPluginStatusChangeSuccess = "PLUGIN_STATUS_CHANGE_SUCCESS"

	OssUserNameUnknown = "unknown"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		// new flag string template if no use, please replace this
		&cli.StringFlag{
			Name:    CliPluginNtpTarget,
			Usage:   "ntp target for sync time like: pool.ntp.org, default not use ntpd to sync",
			EnvVars: []string{EnvPluginNtpTarget},
		},
		&cli.StringFlag{
			Name:    CliPluginWebhook,
			Usage:   "feishu robot webhook",
			EnvVars: []string{EnvPluginWebhook},
		},
		&cli.StringFlag{
			Name:    CliPluginSecret,
			Usage:   "feishu robot secret",
			EnvVars: []string{EnvPluginSecret},
		},
		&cli.BoolFlag{
			Name:    CliPluginFeishuEnableForward,
			Usage:   "feishu robot forward message enable",
			EnvVars: []string{EnvPluginFeishuEnableForward},
		},
		&cli.StringFlag{
			Name:    CliPluginMsgType,
			Usage:   "feishu message type",
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
			Usage:   fmt.Sprintf("must open [ ignore this build success status ], when status change to success, compare with %s", wd_flag.EnvKeyPreviousCiPipelineStatus),
			EnvVars: []string{EnvPluginStatusChangeSuccess},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{}
}

func BindCliFlags(c *cli.Context, cliName, cliVersion string, wdInfo *wd_info.WoodpeckerInfo, rootPath, stepsTransferPath string) (*FeishuPlugin, error) {
	debug := isBuildDebugOpen(c)

	config := Config{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		RootPath:          rootPath,

		NtpTarget:           c.String(CliPluginNtpTarget),
		Webhook:             c.String(CliPluginWebhook),
		Secret:              c.String(CliPluginSecret),
		FeishuEnableForward: c.Bool(CliPluginFeishuEnableForward),
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

	p := FeishuPlugin{
		Name:           cliName,
		Version:        cliVersion,
		WoodpeckerInfo: wdInfo,
		Config:         config,
	}

	return &p, nil
}

// isBuildDebugOpen
// when config or build open debug will open debug
func isBuildDebugOpen(c *cli.Context) bool {
	return c.Bool(wd_flag.NameCliPluginDebug)
}
