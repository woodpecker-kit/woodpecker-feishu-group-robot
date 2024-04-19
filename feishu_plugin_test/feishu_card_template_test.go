package feishu_plugin_test

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/constant"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestRenderFeishuCard(t *testing.T) {
	doTestRenderFeishuCardByi18n(t, "")
}

func TestRenderFeishuCardEnUS(t *testing.T) {
	doTestRenderFeishuCardByi18n(t, constant.LangEnUS)
}

func TestRenderFeishuCardZhCn(t *testing.T) {
	doTestRenderFeishuCardByi18n(t, constant.LangZhCN)
}

func doTestRenderFeishuCardByi18n(t *testing.T, lang string) {
	t.Log("mock FeishuPlugin")

	// sampleRender
	sampleRenderWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo()
	sampleRenderSettings := mockPluginSettings()

	// sampleFailRender
	sampleFailRenderWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
	)
	sampleFailRenderSettings := mockPluginSettings()

	// sampleRenderWithMessage
	sampleRenderWithMessageWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithCurrentCommitInfo(
			wd_mock.WithCiCommitMessage("test: test message\r\n\r\n this is a test message"),
		),
	)
	sampleRenderWithMessageSettings := mockPluginSettings()

	// tagMessageRender
	tagMessageRenderWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastTag("v1.2.3", "new tag"),
	)
	tagMessageRenderSettings := mockPluginSettings()

	// prMessageRender
	prMessageRenderWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastPullRequest("13", "feature-new", "feature-new", "feature-new", "main"),
	)
	prMessageRenderSettings := mockPluginSettings()

	// prCloseMessageRender
	prCloseMessageRenderWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastPullRequestClose("13", "feature-new", "feature-new", "feature-new", "main"),
	)
	prCloseMessageRenderSettings := mockPluginSettings()

	appendTestcase := ""
	if lang != "" {
		appendTestcase = fmt.Sprintf("_%s", lang)
	}

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings
		wantErr        bool
	}{
		{
			name:           fmt.Sprintf("sample_render%s", appendTestcase),
			woodpeckerInfo: sampleRenderWoodpeckerInfo,
			settings:       sampleRenderSettings,
		},
		{
			name:           fmt.Sprintf("sample_fail%s", appendTestcase),
			woodpeckerInfo: sampleFailRenderWoodpeckerInfo,
			settings:       sampleFailRenderSettings,
		},
		{
			name:           fmt.Sprintf("sample_with_message%s", appendTestcase),
			woodpeckerInfo: sampleRenderWithMessageWoodpeckerInfo,
			settings:       sampleRenderWithMessageSettings,
		},
		{
			name:           fmt.Sprintf("tag%s", appendTestcase),
			woodpeckerInfo: tagMessageRenderWoodpeckerInfo,
			settings:       tagMessageRenderSettings,
		},
		{
			name:           fmt.Sprintf("pull_request%s", appendTestcase),
			woodpeckerInfo: prMessageRenderWoodpeckerInfo,
			settings:       prMessageRenderSettings,
		},
		{
			name:           fmt.Sprintf("pull_request_close%s", appendTestcase),
			woodpeckerInfo: prCloseMessageRenderWoodpeckerInfo,
			settings:       prCloseMessageRenderSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// each test case settings start
			tc.settings.Webhook = "some webhook"
			tc.settings.DryRun = true
			tc.settings.I18nLangSet = lang
			// each test case settings end

			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)

			// do RenderFeishuCard
			gotErr := p.Exec()
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("~> RenderFeishuCard error: %s", gotErr.Error())
				return
			}
			// verify RenderFeishuCard
			renderFeishuCard := p.GetInnerCopyRenderFeishuCardContext()
			g.Assert(t, t.Name(), []byte(renderFeishuCard))
		})
	}
}
