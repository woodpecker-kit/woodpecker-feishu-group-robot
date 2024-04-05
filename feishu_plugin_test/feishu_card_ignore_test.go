package feishu_plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestIgnoreConfig(t *testing.T) {
	t.Log("mock TestIgnoreConfig")

	// notIgnore
	notIgnoreWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	notIgnoreSettings := mockPluginSettings()
	notIgnoreSettings.StatusSuccessIgnore = false
	notIgnoreSettings.StatusChangeSuccess = false

	// ignoreSuccessStatusSuccess
	ignoreSuccessStatusSuccessWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	ignoreSuccessStatusSuccessSettings := mockPluginSettings()
	ignoreSuccessStatusSuccessSettings.StatusSuccessIgnore = true
	ignoreSuccessStatusSuccessSettings.StatusChangeSuccess = false

	// ignoreSuccessStatusFailure
	ignoreSuccessStatusFailureWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
	)
	ignoreSuccessStatusFailureSettings := mockPluginSettings()
	ignoreSuccessStatusFailureSettings.StatusSuccessIgnore = true
	ignoreSuccessStatusFailureSettings.StatusChangeSuccess = false

	// ignoreChangeSuccessStatusAllSuccess
	ignoreChangeSuccessStatusAllSuccessWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusSuccess),
		),
	)
	ignoreChangeSuccessStatusAllSuccessSettings := mockPluginSettings()
	ignoreChangeSuccessStatusAllSuccessSettings.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusAllSuccessSettings.StatusChangeSuccess = true

	// ignoreChangeSuccessStatusAllFailure
	ignoreChangeSuccessStatusAllFailureWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)
	ignoreChangeSuccessStatusAllFailureSettings := mockPluginSettings()
	ignoreChangeSuccessStatusAllFailureSettings.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusAllFailureSettings.StatusChangeSuccess = true

	// ignoreChangeSuccessStatusNowFailure
	ignoreChangeSuccessStatusNowFailureWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusSuccess),
		),
	)
	ignoreChangeSuccessStatusNowFailureSettings := mockPluginSettings()
	ignoreChangeSuccessStatusNowFailureSettings.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusNowFailureSettings.StatusChangeSuccess = true

	// ignoreChangeSuccessStatusFailureChange2Success
	ignoreChangeSuccessStatusFailureChange2SuccessWdInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)
	ignoreChangeSuccessStatusFailureChange2SuccessSettings := mockPluginSettings()
	ignoreChangeSuccessStatusFailureChange2SuccessSettings.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusFailureChange2SuccessSettings.StatusChangeSuccess = true

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings
		wantErr        bool
	}{
		{
			name:           "not ignore",
			woodpeckerInfo: notIgnoreWdInfo,
			settings:       notIgnoreSettings,
		},
		{
			name:           "ignoreSuccessStatusSuccess",
			woodpeckerInfo: ignoreSuccessStatusSuccessWdInfo,
			settings:       ignoreSuccessStatusSuccessSettings,
		},
		{
			name:           "ignoreSuccessStatusFailure",
			woodpeckerInfo: ignoreSuccessStatusFailureWdInfo,
			settings:       ignoreSuccessStatusFailureSettings,
		},
		{
			name:           "ignoreChangeSuccessStatusAllSuccess",
			woodpeckerInfo: ignoreChangeSuccessStatusAllSuccessWdInfo,
			settings:       ignoreChangeSuccessStatusAllSuccessSettings,
		},
		{
			name:           "ignoreChangeSuccessStatusAllFailure",
			woodpeckerInfo: ignoreChangeSuccessStatusAllFailureWdInfo,
			settings:       ignoreChangeSuccessStatusAllFailureSettings,
		},
		{
			name:           "ignoreChangeSuccessStatusNowFailure",
			woodpeckerInfo: ignoreChangeSuccessStatusNowFailureWdInfo,
			settings:       ignoreChangeSuccessStatusNowFailureSettings,
		},
		{
			name:           "ignoreChangeSuccessStatusFailureChange2Success",
			woodpeckerInfo: ignoreChangeSuccessStatusFailureChange2SuccessWdInfo,
			settings:       ignoreChangeSuccessStatusFailureChange2SuccessSettings,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// each test case settings start
			tc.settings.Webhook = "some webhook"
			tc.settings.DryRun = true
			// each test case settings end

			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
