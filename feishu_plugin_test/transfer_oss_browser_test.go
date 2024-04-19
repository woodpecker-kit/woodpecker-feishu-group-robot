package feishu_plugin_test

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
	"path/filepath"
	"testing"
)

func TestNoticeTypeFileBrowser(t *testing.T) {
	doRenderTypeFileBrowserByI18n(t, "")
}

func doRenderTypeFileBrowserByI18n(t *testing.T, lang string) {
	// mock NoticeTypeFileBrowser

	// sendWithPasswdAndStatusSuccess
	sendWithPasswdAndStatusSuccessWorkRoot := filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "NoticeTypeFileBrowser", "sendWithPasswdAndStatusSuccess")
	sendWithPasswdAndStatusSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(sendWithPasswdAndStatusSuccessWorkRoot),
	)
	sendWithPasswdAndStatusSuccessSettings := mockPluginSettings()
	sendWithPasswdAndStatusSuccessSettings.Webhook = "some webhook"
	sendWithPasswdAndStatusSuccessSettings.Secret = "some secret"
	sendWithPasswdAndStatusSuccessSettings.NoticeTypes = []string{feishu_plugin.NoticeTypeBuildStatus, feishu_plugin.NoticeTypeFileBrowser}
	sendWithPasswdAndStatusSuccessStepKey := wd_share_file_browser_upload.WdShareKeyFileBrowserUpload
	sendWithPasswdAndStatusSuccessStepOut := wd_share_file_browser_upload.WdShareFileBrowserUpload{
		HostUrl:             mockFileBrowserHostUrl,
		IsSendSuccess:       true,
		FileBrowserUserName: mockFileBrowserUserName,
		ResourceUrl:         mockFileBrowserResourceUrl,
		DownloadPage:        mockFileBrowserDownloadPageUrl,
		DownloadPasswd:      mockFileBrowserDownloadPasswd,
	}

	// sendSuccess
	sendSuccessWorkRoot := filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "NoticeTypeFileBrowser", "sendSuccess")
	sendSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(sendSuccessWorkRoot),
	)
	sendSuccessSettings := mockPluginSettings()
	sendSuccessSettings.Webhook = "some webhook"
	sendSuccessSettings.NoticeTypes = []string{feishu_plugin.NoticeTypeFileBrowser}
	sendSuccessStepKey := wd_share_file_browser_upload.WdShareKeyFileBrowserUpload
	sendSuccessStepOut := wd_share_file_browser_upload.WdShareFileBrowserUpload{
		HostUrl:             mockFileBrowserHostUrl,
		IsSendSuccess:       true,
		FileBrowserUserName: mockFileBrowserUserName,
		ResourceUrl:         mockFileBrowserResourceUrl,
		DownloadPage:        mockFileBrowserDownloadPageUrl,
	}
	// sendSuccessWithPasswd
	sendSuccessWithPasswdWorkRoot := filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "NoticeTypeFileBrowser", "sendSuccessWithPasswd")
	sendSuccessWithPasswdWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(sendSuccessWithPasswdWorkRoot),
	)
	sendSuccessWithPasswdSettings := mockPluginSettings()
	sendSuccessWithPasswdSettings.Webhook = "some webhook"
	sendSuccessWithPasswdSettings.NoticeTypes = []string{feishu_plugin.NoticeTypeFileBrowser}
	sendSuccessWithPasswdStepKey := wd_share_file_browser_upload.WdShareKeyFileBrowserUpload
	sendSuccessWithPasswdStepOut := wd_share_file_browser_upload.WdShareFileBrowserUpload{
		HostUrl:             mockFileBrowserHostUrl,
		IsSendSuccess:       true,
		FileBrowserUserName: mockFileBrowserUserName,
		ResourceUrl:         mockFileBrowserResourceUrl,
		DownloadPage:        mockFileBrowserDownloadPageUrl,
		DownloadPasswd:      mockFileBrowserDownloadPasswd,
	}
	// failSendStatusSuccess
	failSendStatusSuccessRoot := filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "NoticeTypeFileBrowser", "failSendStatusSuccess")
	failSendStatusSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(failSendStatusSuccessRoot),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	failSendStatusSuccessSettings := mockPluginSettings()
	failSendStatusSuccessSettings.Webhook = "some webhook"
	failSendStatusSuccessSettings.NoticeTypes = []string{feishu_plugin.NoticeTypeFileBrowser}
	failSendStatusSuccessStepKey := wd_share_file_browser_upload.WdShareKeyFileBrowserUpload
	failSendStatusSuccessStepOut := wd_share_file_browser_upload.WdShareFileBrowserUpload{
		HostUrl:             mockFileBrowserHostUrl,
		IsSendSuccess:       false,
		FileBrowserUserName: mockFileBrowserUserName,
	}

	// pullRequestSendWithPasswd
	pullRequestSendWithPasswdWorkRoot := filepath.Join(testGoldenKit.GetTestDataFolderFullPath(), "NoticeTypeFileBrowser", "pullRequestSendWithPasswd")
	pullRequestSendWithPasswdWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(pullRequestSendWithPasswdWorkRoot),
		wd_mock.FastPullRequest("5", "new feature", "main", "main", "main"),
	)
	pullRequestSendWithPasswdSettings := mockPluginSettings()
	pullRequestSendWithPasswdSettings.Webhook = "some webhook"
	pullRequestSendWithPasswdSettings.NoticeTypes = []string{feishu_plugin.NoticeTypeFileBrowser}
	pullRequestSendWithPasswdStepKey := wd_share_file_browser_upload.WdShareKeyFileBrowserUpload
	pullRequestSendWithPasswdStepOut := wd_share_file_browser_upload.WdShareFileBrowserUpload{
		HostUrl:             mockFileBrowserHostUrl,
		IsSendSuccess:       true,
		FileBrowserUserName: mockFileBrowserUserName,
		ResourceUrl:         mockFileBrowserResourceUrl,
		DownloadPage:        mockFileBrowserDownloadPageUrl,
		DownloadPasswd:      mockFileBrowserDownloadPasswd,
	}

	appendTestcase := ""
	if lang != "" {
		appendTestcase = fmt.Sprintf("_%s", lang)
	}

	type args struct {
		stepOut wd_share_file_browser_upload.WdShareFileBrowserUpload
		stepKey string
	}
	tests := []struct {
		name           string
		workRoot       string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       feishu_plugin.Settings

		args        args
		isNotDryRun bool

		wantErr bool
	}{
		{
			name:           fmt.Sprintf("sendWithPasswdAndStatusSuccess%s", appendTestcase),
			workRoot:       sendWithPasswdAndStatusSuccessWorkRoot,
			woodpeckerInfo: sendWithPasswdAndStatusSuccessWoodpeckerInfo,
			settings:       sendWithPasswdAndStatusSuccessSettings,
			args: args{
				stepKey: sendWithPasswdAndStatusSuccessStepKey,
				stepOut: sendWithPasswdAndStatusSuccessStepOut,
			},
		},
		{
			name:           fmt.Sprintf("sendSuccess%s", appendTestcase),
			workRoot:       sendSuccessWorkRoot,
			woodpeckerInfo: sendSuccessWoodpeckerInfo,
			settings:       sendSuccessSettings,
			args: args{
				stepKey: sendSuccessStepKey,
				stepOut: sendSuccessStepOut,
			},
		},
		{
			name:           fmt.Sprintf("sendSuccessWithPasswd%s", appendTestcase),
			workRoot:       sendSuccessWithPasswdWorkRoot,
			woodpeckerInfo: sendSuccessWithPasswdWoodpeckerInfo,
			settings:       sendSuccessWithPasswdSettings,
			args: args{
				stepKey: sendSuccessWithPasswdStepKey,
				stepOut: sendSuccessWithPasswdStepOut,
			},
		},
		{
			name:           fmt.Sprintf("failSendStatusSuccess%s", appendTestcase),
			workRoot:       failSendStatusSuccessRoot,
			woodpeckerInfo: failSendStatusSuccessWoodpeckerInfo,
			settings:       failSendStatusSuccessSettings,
			args: args{
				stepKey: failSendStatusSuccessStepKey,
				stepOut: failSendStatusSuccessStepOut,
			},
		},
		{
			name:           fmt.Sprintf("pullRequestSendWithPasswd%s", appendTestcase),
			workRoot:       pullRequestSendWithPasswdWorkRoot,
			woodpeckerInfo: pullRequestSendWithPasswdWoodpeckerInfo,
			settings:       pullRequestSendWithPasswdSettings,
			args: args{
				stepKey: pullRequestSendWithPasswdStepKey,
				stepOut: pullRequestSendWithPasswdStepOut,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)
			if tc.workRoot == "" {
				t.Fatal("workRoot is empty")
			}

			// at some path
			tc.settings.DryRun = !tc.isNotDryRun
			tc.settings.RootPath = tc.woodpeckerInfo.BasicInfo.CIWorkspace
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			// generate transfer data
			errGenTransferData := generateTransferStepsOut(
				p,
				tc.args.stepKey,
				tc.args.stepOut,
			)
			if errGenTransferData != nil {
				t.Fatal(errGenTransferData)
			}

			// do NoticeTypeFileBrowser
			gotErr := p.Exec()
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("~> RenderFeishuCard NoticeTypeFileBrowse error: %s", gotErr.Error())
				return
			}
			renderFeishuCard := p.GetInnerCopyRenderFeishuCardContext()
			// verify RenderFeishuCard
			g.Assert(t, t.Name(), []byte(renderFeishuCard))
		})
	}
}
