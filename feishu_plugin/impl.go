package feishu_plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/constant"
	feishuMessageApi "github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_open_api"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	_ "github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func (p *FeishuPlugin) ShortInfo() wd_short_info.WoodpeckerInfoShort {
	if p.wdShortInfo == nil {
		info2Short := wd_short_info.ParseWoodpeckerInfo2Short(*p.woodpeckerInfo)
		p.wdShortInfo = &info2Short
	}
	return *p.wdShortInfo
}

// SetWoodpeckerInfo
// also change ShortInfo() return
func (p *FeishuPlugin) SetWoodpeckerInfo(info wd_info.WoodpeckerInfo) {
	var newInfo wd_info.WoodpeckerInfo
	_ = struct_kit.DeepCopyByGob(&info, &newInfo)
	p.woodpeckerInfo = &newInfo
	info2Short := wd_short_info.ParseWoodpeckerInfo2Short(newInfo)
	p.wdShortInfo = &info2Short
}

func (p *FeishuPlugin) GetWoodPeckerInfo() wd_info.WoodpeckerInfo {
	return *p.woodpeckerInfo
}

func (p *FeishuPlugin) OnlyArgsCheck() {
	p.onlyArgsCheck = true
}

func (p *FeishuPlugin) Exec() error {
	errLoadStepsTransfer := p.loadStepsTransfer()
	if errLoadStepsTransfer != nil {
		return errLoadStepsTransfer
	}

	errCheckArgs := p.checkArgs()
	if errCheckArgs != nil {
		return fmt.Errorf("check args err %v", errCheckArgs)
	}

	if p.onlyArgsCheck {
		wd_log.Info("only check args, skip do doBiz")
		return nil
	}

	err := p.doBiz()
	if err != nil {
		return err
	}
	errSaveStepsTransfer := p.saveStepsTransfer()
	if errSaveStepsTransfer != nil {
		return errSaveStepsTransfer
	}

	return nil
}

func (p *FeishuPlugin) loadStepsTransfer() error {
	// load steps transfer

	if len(p.Settings.NoticeTypes) > 0 {
		for _, noticeType := range p.Settings.NoticeTypes {
			wd_log.Debugf("just load steps transfer by notice type [ %s ]", noticeType)
			switch noticeType {
			case NoticeTypeFileBrowser:
				errLoadStepsFileBrowser := p.loadStepsFileBrowser()
				if errLoadStepsFileBrowser != nil {
					return errLoadStepsFileBrowser
				}
			}
		}
	}

	return nil
}

func (p *FeishuPlugin) checkArgs() error {
	if p.Settings.Webhook == "" {
		return fmt.Errorf("missing feishu webhook, please set feishu webhook")
	}

	// set default MsgType
	if p.Settings.MsgType == "" {
		p.Settings.MsgType = MsgTypeInteractive
	}

	// set default I18nLangSet or check
	if p.Settings.I18nLangSet == "" {
		p.Settings.I18nLangSet = constant.LangEnUS
	} else {
		if !(string_tools.StringInArr(p.Settings.I18nLangSet, constant.SupportLanguage())) {
			return fmt.Errorf("settings [ feishu-msg-i18n-lang ], now is %s , now support: %s",
				p.Settings.I18nLangSet, strings.Join(constant.SupportLanguage(), ", "),
			)
		}
	}

	if !(string_tools.StringInArr(p.Settings.MsgType, supportMsgType)) {
		return fmt.Errorf("settings [ feishu-msg-type ] only support %v", supportMsgType)
	}

	return nil
}

// doBiz
//
//	replace this code with your feishu_plugin implementation
func (p *FeishuPlugin) doBiz() error {

	if p.Settings.StatusSuccessIgnore {
		if p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus == wd_info.BuildStatusSuccess {
			if p.Settings.StatusChangeSuccess {
				if p.woodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStatus == wd_info.BuildStatusSuccess {
					wd_log.Verbosef("ignore notification previous pipeline just %s, url: %s\n",
						wd_info.BuildStatusSuccess,
						p.woodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineUrl,
					)
					return nil
				}
			} else {
				wd_log.Verbosef("ignore notification this build status [ %s ], pipeline url: %s\n",
					wd_info.BuildStatusSuccess,
					p.woodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl,
				)
				return nil
			}
		}
	}

	// try use ntpd to sync time
	if p.Settings.NtpTarget != "" {
		errNtpSync := p.tryNtpSync()
		if errNtpSync != nil {
			wd_log.Warnf("tryNtpSync err: %v\n", errNtpSync)
		}
	}

	err := p.fetchInfoAndSend()
	if err != nil {
		return err
	}

	return nil
}

func (p *FeishuPlugin) tryNtpSync() error {
	if p.Settings.Debug {
		wd_log.Debugf("NtpTarget sync before by [%v] Unix time: %v\n", p.Settings.NtpTarget, time.Now().Unix())
	}

	wd_log.Verbosef("try to sync ntp by taget [%v]\n", p.Settings.NtpTarget)
	command := exec.Command("ntpd", "-d", "-q", "-n", "-p", p.Settings.NtpTarget)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr

	err := command.Run()
	if err != nil {
		return fmt.Errorf("run ntpd target %v stderr %v\nerr: %v", p.Settings.NtpTarget, stdErr.String(), err)
	}

	if p.Settings.Debug {
		wd_log.Debugf("NtpTarget sync after by [%v] Unix time: %v\n", p.Settings.NtpTarget, time.Now().Unix())
	}

	return nil
}

func (p *FeishuPlugin) fetchInfoAndSend() error {
	sendTarget, errFetch := p.fetchSendTarget()
	if errFetch != nil {
		return errFetch
	}

	if p.Settings.DryRun {
		wd_log.Info("dry run mode not send, please open debug to see more info")
		return nil
	}
	if p.Settings.Debug && !p.Settings.NoticeWhenDebug {
		wd_log.Infof("now debug mode is open, and noticeWhenDebug is close, will not send message")
	}

	errSendMessage := p.sendMessage(sendTarget)
	if errSendMessage != nil {
		return errSendMessage
	}

	return nil
}

func (p *FeishuPlugin) fetchSendTarget() (SendTarget, error) {
	nowTimestamp := time.Now().Unix()
	if p.Settings.Debug {
		wd_log.Debugf("fetchSendTarget nowTimestamp: %v\n", nowTimestamp)
	}
	sendTarget := SendTarget{
		Webhook: p.Settings.Webhook,
		Secret:  p.Settings.Secret,
	}

	ctxTemp := feishuMessageApi.CtxTemp{}

	robotMsgTemplate := feishuMessageApi.FeishuRobotMsgTemplate{
		Timestamp: nowTimestamp,
		MsgType:   p.Settings.MsgType,
	}
	if sendTarget.Secret != "" {
		sign, err := feishuMessageApi.GenSign(sendTarget.Secret, nowTimestamp)
		if err != nil {
			return sendTarget, err
		}
		robotMsgTemplate.Sign = sign
	}

	switch p.Settings.MsgType {
	default:
		wd_log.Debugf("fetchSend msg type now not support %v", p.Settings.MsgType)
		return sendTarget, fmt.Errorf("fetchSend msg type now not support %v", p.Settings.MsgType)
	case MsgTypeInteractive:
		cardTemp := (feishuMessageApi.CardTemp{}).Build(
			p.Settings.Title,
			p.Settings.PoweredByImageKey,
			p.Settings.PoweredByImageAlt,
		)

		cardTemp.EnableForward = p.Settings.FeishuEnableForward
		ctxTemp.CardTemp = cardTemp
		robotMsgTemplate.CtxTemp = ctxTemp
		p.FeishuRobotMsgTemplate = robotMsgTemplate

		renderFeishuCard, errRender := renderFeishuCardFromPlugin(p)
		if errRender != nil {
			return sendTarget, errRender
		}
		if p.Settings.Debug {
			wd_log.Debugf("fetchSendTarget copyRenderFeishuCard: %v\n", renderFeishuCard)
		}
		if renderFeishuCard == "" {
			return sendTarget, fmt.Errorf("fetchSendTarget copyRenderFeishuCard is empty")
		}
		sendTarget.FeishuRobotMeg = []byte(renderFeishuCard)

		errSave := p.saveDeepCopyInnerData()
		if errSave != nil {
			return sendTarget, errSave
		}
	}

	return sendTarget, nil
}

func (p *FeishuPlugin) sendMessage(sendTarget SendTarget) error {
	var feishuUrl = fmt.Sprintf("%s/%s", feishuMessageApi.ApiFeishuBotV2(), sendTarget.Webhook)
	if p.Settings.Debug {
		wd_log.Debugf("sendMessage url: %v", feishuUrl)
	}
	req, err := http.NewRequest("POST", feishuUrl, bytes.NewBuffer(sendTarget.FeishuRobotMeg))
	if err != nil {
		return fmt.Errorf("sendMessage http NewRequest err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * time.Duration(p.Settings.TimeoutSecond),
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sendMessage http Do err: %v", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalf("sendMessage panic err: %v", err)
		}
	}()
	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		if errBody != nil {
			return fmt.Errorf("sendMessage http status code: %v , body: %v", statusCode, string(errBody))
		}
		return fmt.Errorf("sendMessage http status code: %v", statusCode)
	}
	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return fmt.Errorf("sendMessage http read err: %v", errRead)
	}
	if p.Settings.Debug {
		wd_log.Debugf("response code [%v], response: %s", statusCode, string(body))
	}
	var respApi feishuMessageApi.ApiRespRotV2
	errUnmarshal := json.Unmarshal(body, &respApi)

	if errUnmarshal != nil {
		return fmt.Errorf("sendMessage http Unmarshal err: %v", errUnmarshal)
	}
	if respApi.Code != 0 {
		return fmt.Errorf("feishu message can not send by code [ %v ] err: %v", respApi.Code, respApi.Msg)
	}
	return nil
}

func (p *FeishuPlugin) saveStepsTransfer() error {
	if p.Settings.StepsOutDisable {
		wd_log.Debugf("steps out disable by flag [ %v ], skip save steps transfer", p.Settings.StepsOutDisable)
		return nil
	}
	//if p.Settings.StepsTransferDemo {
	//	transferAppendObj, errSave := wd_steps_transfer.Out(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.WoodpeckerInfo, "config", p.Settings)
	//	if errSave != nil {
	//		return errSave
	//	}
	//	wd_log.VerboseJsonf(transferAppendObj, "save steps transfer config mark [ %s ]", "config")
	//}
	return nil
}
