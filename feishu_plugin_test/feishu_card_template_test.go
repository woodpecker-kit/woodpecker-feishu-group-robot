package feishu_plugin_test

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"testing"
)

func TestRenderFeishuCard(t *testing.T) {

	// template config start
	t.Log("mock FeishuPlugin")
	p := mockPlugin(t)

	// use env:ENV_DEBUG
	p.Config.Debug = valEnvPluginDebug
	p.Config.TimeoutSecond = envTimeoutSecond
	p.Config.RootPath = testGoldenKit.GetTestDataFolderFullPath()
	p.Config.StepsTransferPath = wd_steps_transfer.DefaultKitStepsFileName

	p.Config.FeishuEnableForward = false
	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)
	p.WoodpeckerInfo = woodpeckerInfo
	// template config end

	var sampleRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &sampleRender)

	var sampleFailRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &sampleFailRender)
	sampleFailRender.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)

	var sampleRenderWithMessage feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &sampleRenderWithMessage)
	sampleRenderWithMessage.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithCurrentCommitInfo(
			wd_mock.WithCiCommitMessage("test: test message\r\n\r\n this is a test message"),
		),
	)

	var tagMessageRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &tagMessageRender)
	tagMessageRender.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithCurrentPipelineInfo(
			wd_mock.WithCiPipelineEvent(wd_info.EventPipelineTag),
		),
		wd_mock.WithCurrentCommitInfo(
			wd_mock.WithCiCommitRef(fmt.Sprintf("refs/tags/%s", "v1.2.3")),
			wd_mock.WithCiCommitTag("v1.2.3"),
			wd_mock.WithCiCommitBranch(""),
		),
	)
	//tagMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl = fmt.Sprintf(
	//	"%s/tag/%s",
	//	tagMessageRender.WoodpeckerInfo.RepositoryInfo.CIRepoURL, tagMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag,
	//)

	var prMessageRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &prMessageRender)

	prMessageRender.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineInfo(
			wd_mock.WithCiPipelineEvent(wd_info.EventPipelinePullRequest),
		),
		wd_mock.WithCurrentCommitInfo(
			wd_mock.WithCiCommitBranch("main"),
			wd_mock.WithCiCommitPullRequest("13"),
			wd_mock.WithCiCommitRef(fmt.Sprintf("refs/pull/%s/head", "13")),
			wd_mock.WithCiCommitSourceBranch("feature-new"),
			wd_mock.WithCiCommitTargetBranch("main"),
		),
	)
	prMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl = fmt.Sprintf(
		"%s/pulls/%s",
		prMessageRender.WoodpeckerInfo.RepositoryInfo.CIRepoURL, prMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest,
	)

	var prCloseMessageRender feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &prCloseMessageRender)
	prCloseMessageRender.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineInfo(
			wd_mock.WithCiPipelineEvent(wd_info.EventPipelinePullRequestClose),
		),
		wd_mock.WithCurrentCommitInfo(
			wd_mock.WithCiCommitBranch("main"),
			wd_mock.WithCiCommitPullRequest("13"),
			wd_mock.WithCiCommitRef(fmt.Sprintf("refs/pull/%s/head", "13")),
			wd_mock.WithCiCommitSourceBranch("feature-new"),
			wd_mock.WithCiCommitTargetBranch("main"),
		),
	)
	prCloseMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl = fmt.Sprintf(
		"%s/pulls/%s",
		prCloseMessageRender.WoodpeckerInfo.RepositoryInfo.CIRepoURL, prCloseMessageRender.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest,
	)

	var ossRenderSendSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossRenderSendSuccess)
	ossRenderSendSuccess.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	checkCardOssRenderByPlugin(&ossRenderSendSuccess, "", true)

	var ossRenderSendSuccessWithPass feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossRenderSendSuccessWithPass)
	ossRenderSendSuccessWithPass.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	checkCardOssRenderByPlugin(&ossRenderSendSuccessWithPass, mockOssPagePasswd, true)

	var ossRenderSendFail feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossRenderSendFail)
	ossRenderSendFail.Config.RenderOssCard = feishu_plugin.RenderStatusShow
	checkCardOssRenderByPlugin(&ossRenderSendFail, "", false)

	tests := []struct {
		name    string
		p       feishu_plugin.FeishuPlugin
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestRenderFeishuCard/sample.golden
			p:    sampleRender,
		},
		{
			name: "sample_fail", // testdata/TestRenderFeishuCard/sample_fail.golden
			p:    sampleFailRender,
		},
		{
			name: "sample_with_message", // testdata/TestRenderFeishuCard/sample_with_message.golden
			p:    sampleRenderWithMessage,
		},
		{
			name: "tag", // testdata/TestRenderFeishuCard/tag.golden
			p:    tagMessageRender,
		},
		{
			name: "pull_request", // testdata/TestRenderFeishuCard/pull_request.golden
			p:    prMessageRender,
		},
		{
			name: "pull_request_close", // testdata/TestRenderFeishuCard/pull_request_close.golden
			p:    prCloseMessageRender,
		},
		{
			name: "oss_render_send_success", // testdata/TestRenderFeishuCard/oss_render_send_success.golden
			p:    ossRenderSendSuccess,
		},
		{
			name: "oss_render_send_success_with_pass", // testdata/TestRenderFeishuCard/oss_render_send_success_with_pass.golden
			p:    ossRenderSendSuccessWithPass,
		},
		{
			name: "oss_render_send_fail", // testdata/TestRenderFeishuCard/oss_render_send_fail.golden
			p:    ossRenderSendFail,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do RenderFeishuCard
			renderFeishuCard, gotErr := feishu_plugin.RenderFeishuCardByPlugin(feishu_plugin.DefaultCardTemplate, &tc.p)
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("~> RenderFeishuCard error: %s", gotErr.Error())
				return
			}
			// verify RenderFeishuCard
			g.Assert(t, t.Name(), []byte(renderFeishuCard))
		})
	}

}
