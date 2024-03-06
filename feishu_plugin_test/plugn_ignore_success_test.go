package feishu_plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestIgnoreConfig(t *testing.T) {
	t.Log("mock TestIgnoreConfig")
	p := mockPlugin(t)

	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)
	// mock must args
	p.Config.DryRun = true
	p.WoodpeckerInfo = woodpeckerInfo
	p.Config.Webhook = "some webhook"

	// notIgnore
	var notIgnore feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &notIgnore)
	notIgnore.Config.StatusSuccessIgnore = false
	notIgnore.Config.StatusChangeSuccess = false

	// ignoreSuccessStatusSuccess
	var ignoreSuccessStatusSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreSuccessStatusSuccess)
	ignoreSuccessStatusSuccess.Config.StatusSuccessIgnore = true
	ignoreSuccessStatusSuccess.Config.StatusChangeSuccess = false
	ignoreSuccessStatusSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)

	// ignoreSuccessStatusFailure
	var ignoreSuccessStatusFailure feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreSuccessStatusFailure)
	ignoreSuccessStatusFailure.Config.StatusSuccessIgnore = true
	ignoreSuccessStatusFailure.Config.StatusChangeSuccess = false
	ignoreSuccessStatusFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)

	// ignoreChangeSuccessStatusAllSuccess
	var ignoreChangeSuccessStatusAllSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreChangeSuccessStatusAllSuccess)
	ignoreChangeSuccessStatusAllSuccess.Config.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusAllSuccess.Config.StatusChangeSuccess = true
	ignoreChangeSuccessStatusAllSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusSuccess),
		),
	)

	// ignoreChangeSuccessStatusAllFailure
	var ignoreChangeSuccessStatusAllFailure feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreChangeSuccessStatusAllFailure)
	ignoreChangeSuccessStatusAllFailure.Config.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusAllFailure.Config.StatusChangeSuccess = true
	ignoreChangeSuccessStatusAllFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)

	// ignoreChangeSuccessStatusNowFailure
	var ignoreChangeSuccessStatusNowFailure feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreChangeSuccessStatusNowFailure)
	ignoreChangeSuccessStatusNowFailure.Config.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusNowFailure.Config.StatusChangeSuccess = true
	ignoreChangeSuccessStatusNowFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusSuccess),
		),
	)
	// ignoreChangeSuccessStatusFailureChange2Success
	var ignoreChangeSuccessStatusFailureChange2Success feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ignoreChangeSuccessStatusFailureChange2Success)
	ignoreChangeSuccessStatusFailureChange2Success.Config.StatusSuccessIgnore = true
	ignoreChangeSuccessStatusFailureChange2Success.Config.StatusChangeSuccess = true
	ignoreChangeSuccessStatusFailureChange2Success.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)

	tests := []struct {
		name    string
		p       feishu_plugin.FeishuPlugin
		wantErr bool
	}{
		{
			name: "not ignore",
			p:    notIgnore,
		},
		{
			name: "ignoreSuccessStatusSuccess",
			p:    ignoreSuccessStatusSuccess,
		},
		{
			name: "ignoreSuccessStatusFailure",
			p:    ignoreSuccessStatusFailure,
		},
		{
			name: "ignoreChangeSuccessStatusAllSuccess",
			p:    ignoreChangeSuccessStatusAllSuccess,
		},
		{
			name: "ignoreChangeSuccessStatusAllFailure",
			p:    ignoreChangeSuccessStatusAllFailure,
		},
		{
			name: "ignoreChangeSuccessStatusNowFailure",
			p:    ignoreChangeSuccessStatusNowFailure,
		},
		{
			name: "ignoreChangeSuccessStatusFailureChange2Success",
			p:    ignoreChangeSuccessStatusFailureChange2Success,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
