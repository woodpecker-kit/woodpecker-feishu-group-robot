package feishu_plugin_test

import (
	"encoding/json"
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	keyEnvDebug        = "CI_DEBUG"
	keyEnvCiNum        = "CI_NUMBER"
	keyEnvCiKey        = "CI_KEY"
	keyEnvCiKeys       = "CI_KEYS"
	mockVersion        = "1.0.0"
	mockName           = "woodpecker-feishu-group-robot"
	mockTitle          = "CI Notification"
	mockOssHost        = "https://docs.aws.amazon.com/s3/index.html"
	mockOssUser        = "ossAdmin"
	mockOssPath        = "dist/demo/pass.tar.gz"
	mockOssResourceUrl = "https://docs.aws.amazon.com/s/dist/demo/pass.tar.gz"
	mockOssPageUrl     = "https://docs.aws.amazon.com/p/dist/demo/page-xyz.html"
	mockOssPagePasswd  = "abc-zxy"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	envTimeoutSecond    uint
	envPaddingLeftMax   = 0
	envPrinterPrintKeys []string

	// mustSetInCiEnvList
	//  for check set in CI env not empty
	mustSetInCiEnvList = []string{
		wd_flag.EnvKeyCiSystemPlatform,
		wd_flag.EnvKeyCiSystemVersion,
	}

	// mustSetArgsAsEnvList
	mustSetArgsAsEnvList = []string{
		feishu_plugin.EnvPluginWebhook,
	}

	valEnvPluginDebug   = false
	valEnvPluginWebHook = ""
	valEnvPluginSecret  = ""
	valEnvPluginTitle   = ""
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	wd_log.SetLogLineDeep(2)

	valEnvPluginDebug = env_kit.FetchOsEnvBool(wd_flag.EnvKeyPluginDebug, false)
	envTimeoutSecond = uint(env_kit.FetchOsEnvInt(wd_flag.EnvKeyPluginTimeoutSecond, 10))
	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	wd_template.RegisterSettings(wd_template.DefaultHelpers)

	valEnvPluginWebHook = env_kit.FetchOsEnvStr(feishu_plugin.EnvPluginWebhook, "")
	valEnvPluginSecret = env_kit.FetchOsEnvStr(feishu_plugin.EnvPluginSecret, "")
	valEnvPluginTitle = env_kit.FetchOsEnvStr(feishu_plugin.EnvPluginTitle, mockTitle)
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end

func envCheck(t *testing.T) bool {

	if valEnvPluginDebug {
		wd_log.OpenDebug()
	}

	// most CI system will set env CI to true
	envCI := env_kit.FetchOsEnvStr("CI", "")
	if envCI == "" {
		t.Logf("not in CI system, skip envCheck")
		return false
	}
	t.Logf("check env for CI system")
	for _, item := range mustSetInCiEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetInCiEnvList)
			return true
		}
	}
	return false
}

func envMustArgsCheck(t *testing.T) bool {
	for _, item := range mustSetArgsAsEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetArgsAsEnvList)
			return true
		}
	}
	return false
}

func mockPlugin(t *testing.T) feishu_plugin.FeishuPlugin {

	p := feishu_plugin.FeishuPlugin{
		Name:    mockName,
		Version: mockVersion,
	}

	// use env:ENV_DEBUG
	p.Config.Debug = valEnvPluginDebug
	p.Config.TimeoutSecond = envTimeoutSecond
	p.Config.RootPath = testGoldenKit.GetTestDataFolderFullPath()
	p.Config.StepsTransferPath = wd_steps_transfer.DefaultKitStepsFileName

	p.Config.Webhook = valEnvPluginWebHook
	p.Config.Title = valEnvPluginTitle
	p.Config.Secret = valEnvPluginSecret

	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)
	p.WoodpeckerInfo = woodpeckerInfo
	return p
}

func deepCopyByPlugin(src, dst *feishu_plugin.FeishuPlugin) {
	if tmp, err := json.Marshal(&src); err != nil {
		return
	} else {
		err = json.Unmarshal(tmp, dst)
		return
	}
}

func checkCardOssRenderByPlugin(p *feishu_plugin.FeishuPlugin, pagePasswd string, sendOssSucc bool) {
	p.Config.CardOss.PagePasswd = pagePasswd
	if p.Config.CardOss.PagePasswd == "" {
		p.Config.CardOss.RenderResourceUrl = feishu_plugin.RenderStatusShow
	} else {
		p.Config.CardOss.RenderResourceUrl = feishu_plugin.RenderStatusHide
	}
	if sendOssSucc {
		p.Config.CardOss.InfoSendResult = feishu_plugin.RenderStatusShow
	} else {
		p.Config.CardOss.InfoSendResult = feishu_plugin.RenderStatusHide
	}
	p.Config.CardOss.Host = mockOssHost
	p.Config.CardOss.InfoUser = mockOssUser
	p.Config.CardOss.InfoPath = mockOssPath
	p.Config.CardOss.ResourceUrl = mockOssResourceUrl
	p.Config.CardOss.PageUrl = mockOssPageUrl
}
