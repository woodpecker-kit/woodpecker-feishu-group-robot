package feishu_plugin_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestRenderFeishuCard(t *testing.T) {

	// template config start
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

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings
		wantErr        bool
	}{
		{
			name:           "sample_render", // testdata/TestRenderFeishuCard/sample_render.golden
			woodpeckerInfo: sampleRenderWoodpeckerInfo,
			settings:       sampleRenderSettings,
		},
		{
			name:           "sample_fail", // testdata/TestRenderFeishuCard/sample_fail.golden
			woodpeckerInfo: sampleFailRenderWoodpeckerInfo,
			settings:       sampleFailRenderSettings,
		},
		{
			name:           "sample_with_message", // testdata/TestRenderFeishuCard/sample_with_message.golden
			woodpeckerInfo: sampleRenderWithMessageWoodpeckerInfo,
			settings:       sampleRenderWithMessageSettings,
		},
		{
			name:           "tag", // testdata/TestRenderFeishuCard/tag.golden
			woodpeckerInfo: tagMessageRenderWoodpeckerInfo,
			settings:       tagMessageRenderSettings,
		},
		{
			name:           "pull_request", // testdata/TestRenderFeishuCard/pull_request.golden
			woodpeckerInfo: prMessageRenderWoodpeckerInfo,
			settings:       prMessageRenderSettings,
		},
		{
			name:           "pull_request_close", // testdata/TestRenderFeishuCard/pull_request_close.golden
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
