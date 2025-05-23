package feishu_plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock TestCheckArgsPlugin")
	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()
	successArgsSettings.Webhook = "some webhook"

	// forceStatusFailure
	forceStatusFailureWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	forceStatusFailureSettings := mockPluginSettings()
	forceStatusFailureSettings.Webhook = "some webhook"
	forceStatusFailureSettings.ForceStatus = wd_info.BuildStatusFailure

	// emptyWebhook
	emptyWebhookWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	emptyWebhookSettings := mockPluginSettings()
	emptyWebhookSettings.Webhook = ""

	// errorLang
	errorLangWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	errorLangSettings := mockPluginSettings()
	errorLangSettings.Webhook = "some webhook"
	errorLangSettings.I18nLangSet = "errorLang"

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:              "forceStatusFailure",
			woodpeckerInfo:    forceStatusFailureWoodpeckerInfo,
			settings:          forceStatusFailureSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "emptyWebhook",
			woodpeckerInfo: emptyWebhookWoodpeckerInfo,
			settings:       emptyWebhookSettings,
		},
		{
			name:           "errorLang",
			woodpeckerInfo: errorLangWoodpeckerInfo,
			settings:       errorLangSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}

	t.Log("mock woodpecker plugin")

	// statusSuccessIgnore
	statusSuccessIgnoreWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusSuccess),
		),
	)
	statusSuccessIgnoreSettings := mockPluginSettings()
	statusSuccessIgnoreSettings.StatusSuccessIgnore = true

	// statusChangeSuccess
	statusChangeSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)
	statusChangeSuccessSettings := mockPluginSettings()
	statusChangeSuccessSettings.StatusChangeSuccess = true

	// statusSuccess
	statusSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	statusSuccessSettings := mockPluginSettings()

	// statusFailure
	statusFailureWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
	)
	statusFailureSettings := mockPluginSettings()

	// tagSuccess
	tagSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastTag("v1.2.3", "new tag v1.2.3"),
	)
	tagSuccessSettings := mockPluginSettings()

	// prOpenSuccess
	prOpenSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastPullRequest("13", "new pr 13", "feature-new", "main", "main"),
	)
	prOpenSuccessSettings := mockPluginSettings()

	// prCloseSuccess
	prCloseSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithFastMockPullRequest("13", "new pr 13", "feature-new", "main", "main"),
	)
	prCloseSuccessSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings
		isDryRun       bool
		wantErr        bool
	}{
		{
			name:           "statusSuccessIgnore",
			woodpeckerInfo: statusSuccessIgnoreWoodpeckerInfo,
			settings:       statusSuccessIgnoreSettings,
			isDryRun:       false,
		},
		{
			name:           "statusChangeSuccess",
			woodpeckerInfo: statusChangeSuccessWoodpeckerInfo,
			settings:       statusChangeSuccessSettings,
			isDryRun:       true,
		},
		{
			name:           "statusSuccess",
			woodpeckerInfo: statusSuccessWoodpeckerInfo,
			settings:       statusSuccessSettings,
			isDryRun:       true,
		},
		{
			name:           "statusFailure",
			woodpeckerInfo: statusFailureWoodpeckerInfo,
			settings:       statusFailureSettings,
			isDryRun:       true,
		},
		{
			name:           "tagSuccess",
			woodpeckerInfo: tagSuccessWoodpeckerInfo,
			settings:       tagSuccessSettings,
			isDryRun:       true,
		},
		{
			name:           "prOpenSuccess",
			woodpeckerInfo: prOpenSuccessWoodpeckerInfo,
			settings:       prOpenSuccessSettings,
			isDryRun:       true,
		},
		{
			name:           "prCloseSuccess",
			woodpeckerInfo: prCloseSuccessWoodpeckerInfo,
			settings:       prCloseSuccessSettings,
			isDryRun:       true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.settings.DryRun = tc.isDryRun

			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
