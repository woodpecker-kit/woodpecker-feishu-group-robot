package feishu_plugin_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin_transfer"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"path/filepath"
	"testing"
)

func TestPluginMustArgs(t *testing.T) {
	t.Log("mock FeishuPlugin")
	p := mockPlugin(t)

	t.Log("mock woodpecker info")
	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusCreated),
	)
	p.WoodpeckerInfo = woodpeckerInfo

	p.Config.Webhook = ""
	err := p.Exec()
	assert.Equal(t, fmt.Errorf("check args err missing feishu webhook, please set feishu webhook"), err)

	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}

	t.Log("do FeishuPlugin")
	err = p.Exec()
	t.Log("verify woodpecker FeishuPlugin")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPlugin(t *testing.T) {
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}

	t.Log("mock FeishuPlugin")
	p := mockPlugin(t)

	t.Log("mock woodpecker plugin")
	// statusSuccessIgnore
	var statusSuccessIgnore feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &statusSuccessIgnore)
	statusSuccessIgnore.Config.StatusSuccessIgnore = true

	// statusChangeSuccess
	var statusChangeSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &statusChangeSuccess)
	statusChangeSuccess.Config.StatusChangeSuccess = true
	statusChangeSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
		wd_mock.WithPreviousPipelineInfo(
			wd_mock.WithCiPreviousPipelineStatus(wd_info.BuildStatusFailure),
		),
	)

	// statusSuccess
	var statusSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &statusSuccess)

	// statusFailure
	var statusFailure feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &statusFailure)
	statusFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)

	// tagSuccess
	var tagSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &tagSuccess)
	tagSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
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

	// prOpenSuccess
	var prOpenSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &prOpenSuccess)
	prOpenSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
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
	// prCloseSuccess
	var prCloseSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &prCloseSuccess)
	prCloseSuccess.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
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
	// ossCardSendFailure
	var ossCardSendFailure feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossCardSendFailure)

	// ossCardOssSuccess
	var ossCardOssSuccess feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossCardOssSuccess)

	// ossCardSendSuccessWithPass
	var ossCardSendSuccessWithPass feishu_plugin.FeishuPlugin
	deepCopyByPlugin(&p, &ossCardSendSuccessWithPass)

	tests := []struct {
		name            string
		p               feishu_plugin.FeishuPlugin
		isDryRun        bool
		workRoot        string
		ossTransferKey  string
		ossTransferData feishu_plugin_transfer.OssSendTransfer
		wantErr         bool
	}{
		{
			name:     "statusSuccessIgnore",
			p:        statusSuccessIgnore,
			isDryRun: false,
		},
		{
			name:     "statusChangeSuccess",
			p:        statusChangeSuccess,
			isDryRun: true,
		},
		{
			name:     "statusSuccess",
			p:        statusSuccess,
			isDryRun: true,
		},
		{
			name:     "statusFailure",
			p:        statusFailure,
			isDryRun: true,
		},
		{
			name:     "tagSuccess",
			p:        tagSuccess,
			isDryRun: true,
		},
		{
			name:     "prOpenSuccess",
			p:        prOpenSuccess,
			isDryRun: true,
		},
		{
			name:     "prCloseSuccess",
			p:        prCloseSuccess,
			isDryRun: true,
		},
		{
			name:           "ossCardSendFailure",
			p:              ossCardSendFailure,
			workRoot:       filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "ossCardSendFailure"),
			ossTransferKey: feishu_plugin_transfer.OssSendTransferKey,
			ossTransferData: feishu_plugin_transfer.OssSendTransfer{
				InfoSendResult: wd_info.BuildStatusFailure,
				OssHost:        mockOssHost,
			},
			isDryRun: true,
		},
		{
			name:           "ossCardOssSuccess",
			p:              ossCardOssSuccess,
			workRoot:       filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "ossCardOssSuccess"),
			ossTransferKey: feishu_plugin_transfer.OssSendTransferKey,
			ossTransferData: feishu_plugin_transfer.OssSendTransfer{
				InfoSendResult: wd_info.BuildStatusSuccess,
				OssHost:        mockOssHost,
				OssPath:        mockOssPath,
				ResourceUrl:    mockOssResourceUrl,
			},
			isDryRun: true,
		},
		{
			name:           "ossCardSendSuccessWithPass",
			p:              ossCardSendSuccessWithPass,
			workRoot:       filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "ossCardSendSuccessWithPass"),
			ossTransferKey: feishu_plugin_transfer.OssSendTransferKey,
			ossTransferData: feishu_plugin_transfer.OssSendTransfer{
				InfoSendResult: wd_info.BuildStatusSuccess,
				OssHost:        mockOssHost,
				OssPath:        mockOssPath,
				PageUrl:        mockOssPageUrl,
				PagePasswd:     mockOssPagePasswd,
			},
			isDryRun: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Config.DryRun = tc.isDryRun
			if tc.workRoot != "" {
				tc.p.Config.RootPath = tc.workRoot
				errGenTransferData := generateTransferStepsOut(
					tc.p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := tc.p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func generateTransferStepsOut(plugin feishu_plugin.FeishuPlugin, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Config.RootPath, plugin.Config.StepsTransferPath, *plugin.WoodpeckerInfo, mark, data)
	return err
}
