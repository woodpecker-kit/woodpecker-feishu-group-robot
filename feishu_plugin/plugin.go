package feishu_plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	feishu_message "github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_open_api"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin_transfer"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type (
	// FeishuPlugin feishu_plugin all config
	FeishuPlugin struct {
		Name           string
		Version        string
		WoodpeckerInfo *wd_info.WoodpeckerInfo
		Config         Config

		FeishuRobotMsgTemplate feishu_message.FeishuRobotMsgTemplate
		FuncPlugin             FuncPlugin `json:"-"`
	}
)

type FuncPlugin interface {
	Exec() error

	loadStepsTransfer() error
	checkArgs() error
	saveStepsTransfer() error
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
	// set default
	p.Config.RenderOssCard = RenderStatusHide
	p.Config.CardOss.RenderResourceUrl = RenderStatusHide

	var sendTransfer feishu_plugin_transfer.OssSendTransfer
	errReadSendTransfer := wd_steps_transfer.In(p.Config.RootPath, p.Config.StepsTransferPath, *p.WoodpeckerInfo, feishu_plugin_transfer.OssSendTransferKey, &sendTransfer)
	if errReadSendTransfer == nil {
		wd_log.DebugJsonf(sendTransfer, "load steps transfer config mark [ %s ]", feishu_plugin_transfer.OssSendTransferKey)
		if sendTransfer.OssHost != "" {
			p.Config.CardOss.InfoSendResult = sendTransfer.InfoSendResult
			p.Config.RenderOssCard = RenderStatusShow
			p.Config.CardOss.Host = sendTransfer.OssHost
			ossUserName := sendTransfer.OssUserProfile.UserName
			if ossUserName == "" {
				ossUserName = OssUserNameUnknown
			}
			p.Config.CardOss.InfoUser = ossUserName
			p.Config.CardOss.InfoPath = sendTransfer.OssPath
			if sendTransfer.PagePasswd == "" {
				p.Config.CardOss.RenderResourceUrl = RenderStatusShow
				p.Config.CardOss.ResourceUrl = sendTransfer.ResourceUrl
			} else {
				p.Config.CardOss.RenderResourceUrl = RenderStatusHide
				p.Config.CardOss.PagePasswd = sendTransfer.PagePasswd
				p.Config.CardOss.PageUrl = sendTransfer.PageUrl
			}
		}
	}

	return nil
}

func (p *FeishuPlugin) checkArgs() error {
	if p.Config.Webhook == "" {
		return fmt.Errorf("missing feishu webhook, please set feishu webhook")
	}

	// set default MsgType
	if p.Config.MsgType == "" {
		p.Config.MsgType = MsgTypeInteractive
	}

	if !(string_tools.StringInArr(p.Config.MsgType, supportMsgType)) {
		return fmt.Errorf("config [ msg_type ] only support %v", supportMsgType)
	}

	if p.Config.RenderOssCard == RenderStatusShow {
		if !(string_tools.StringInArr(p.Config.CardOss.InfoSendResult, supportRenderStatus)) {
			return fmt.Errorf("config [ p.Config.CardOss.InfoSendResult ] only support %v , now is %s", supportRenderStatus, p.Config.CardOss.InfoSendResult)
		}
		if !(string_tools.StringInArr(p.Config.CardOss.RenderResourceUrl, supportRenderStatus)) {
			return fmt.Errorf("config [ p.Config.CardOss.RenderResourceUrl ] only support %v , now is %s", supportRenderStatus, p.Config.CardOss.RenderResourceUrl)
		}
		if p.Config.CardOss.InfoSendResult != RenderStatusShow {
			if p.Config.Debug {
				wd_log.Debugf("in p.Config.RenderOssCard mode [ %s ] will set p.Config.CardOss.InfoSendResult to [ %s ] and change p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus to [ %s ]\n",
					RenderStatusShow, RenderStatusHide, RenderStatusHide,
				)
			}
			p.Config.CardOss.InfoSendResult = RenderStatusHide
			p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus = wd_info.BuildStatusFailure
		}
	}

	return nil
}

// doBiz
//
//	replace this code with your feishu_plugin implementation
func (p *FeishuPlugin) doBiz() error {

	if p.Config.StatusSuccessIgnore {
		if p.Config.StatusChangeSuccess {
			if p.WoodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStatus == wd_info.BuildStatusSuccess {
				wd_log.Verbosef("ignore notification previous pipeline just %s, url: %s\n",
					wd_info.BuildStatusSuccess,
					p.WoodpeckerInfo.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineUrl,
				)
				return nil
			}

		} else {
			if p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus == wd_info.BuildStatusSuccess {
				wd_log.Verbosef("ignore notification this build status [ %s ], pipeline url: %s\n",
					wd_info.BuildStatusSuccess,
					p.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl,
				)
				return nil
			}
		}
	}

	// try use ntpd to sync time
	if p.Config.NtpTarget != "" {
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
	if p.Config.Debug {
		wd_log.Debugf("NtpTarget sync before by [%v] Unix time: %v\n", p.Config.NtpTarget, time.Now().Unix())
	}

	wd_log.Verbosef("try to sync ntp by taget [%v]\n", p.Config.NtpTarget)
	command := exec.Command("ntpd", "-d", "-q", "-n", "-p", p.Config.NtpTarget)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr

	err := command.Run()
	if err != nil {
		return fmt.Errorf("run ntpd target %v stderr %v\nerr: %v", p.Config.NtpTarget, stdErr.String(), err)
	}

	if p.Config.Debug {
		wd_log.Debugf("NtpTarget sync after by [%v] Unix time: %v\n", p.Config.NtpTarget, time.Now().Unix())
	}

	return nil
}

func (p *FeishuPlugin) fetchInfoAndSend() error {
	sendTarget, errFetch := p.fetchSendTarget()
	if errFetch != nil {
		return errFetch
	}
	if p.Config.DryRun {
		wd_log.Info("dry run mode not send, please open debug to see more info")
		return nil
	}

	errSendMessage := p.sendMessage(sendTarget)
	if errSendMessage != nil {
		return errSendMessage
	}

	return nil
}

func (p *FeishuPlugin) fetchSendTarget() (SendTarget, error) {
	nowTimestamp := time.Now().Unix()
	if p.Config.Debug {
		wd_log.Debugf("fetchSendTarget nowTimestamp: %v\n", nowTimestamp)
	}
	sendTarget := SendTarget{
		Webhook: p.Config.Webhook,
		Secret:  p.Config.Secret,
	}

	ctxTemp := feishu_message.CtxTemp{}

	robotMsgTemplate := feishu_message.FeishuRobotMsgTemplate{
		Timestamp: nowTimestamp,
		MsgType:   p.Config.MsgType,
	}
	if sendTarget.Secret != "" {
		sign, err := feishu_message.GenSign(sendTarget.Secret, nowTimestamp)
		if err != nil {
			return sendTarget, err
		}
		robotMsgTemplate.Sign = sign
	}

	switch p.Config.MsgType {
	default:
		wd_log.Debugf("fetchSend msg type now not support %v", p.Config.MsgType)
		return sendTarget, fmt.Errorf("fetchSend msg type now not support %v", p.Config.MsgType)
	case MsgTypeInteractive:
		cardTemp := (feishu_message.CardTemp{}).Build(
			p.Config.Title,
			p.Config.PoweredByImageKey,
			p.Config.PoweredByImageAlt,
		)

		cardTemp.EnableForward = p.Config.FeishuEnableForward
		ctxTemp.CardTemp = cardTemp
		robotMsgTemplate.CtxTemp = ctxTemp
		p.FeishuRobotMsgTemplate = robotMsgTemplate

		renderFeishuCard, err := RenderFeishuCardByPlugin(DefaultCardTemplate, p)
		if err != nil {
			return sendTarget, err
		}
		if p.Config.Debug {
			wd_log.Debugf("fetchSendTarget renderFeishuCard: %v\n", renderFeishuCard)
		}
		if renderFeishuCard != "" {
			sendTarget.FeishuRobotMeg = []byte(renderFeishuCard)
		}
	}

	return sendTarget, nil
}

func (p *FeishuPlugin) sendMessage(sendTarget SendTarget) error {
	var feishuUrl = fmt.Sprintf("%s/%s", feishu_message.ApiFeishuBotV2(), sendTarget.Webhook)
	if p.Config.Debug {
		wd_log.Debugf("sendMessage url: %v", feishuUrl)
	}
	req, err := http.NewRequest("POST", feishuUrl, bytes.NewBuffer(sendTarget.FeishuRobotMeg))
	if err != nil {
		return fmt.Errorf("sendMessage http NewRequest err: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * time.Duration(p.Config.TimeoutSecond),
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
	if p.Config.Debug {
		wd_log.Debugf("response code [%v], response: %s", statusCode, string(body))
	}
	var respApi feishu_message.ApiRespRotV2
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
	//if p.Config.StepsTransferDemo {
	//	transferAppendObj, errSave := wd_steps_transfer.Out(p.Config.RootPath, p.Config.StepsTransferPath, *p.WoodpeckerInfo, "config", p.Config)
	//	if errSave != nil {
	//		return errSave
	//	}
	//	wd_log.VerboseJsonf(transferAppendObj, "save steps transfer config mark [ %s ]", "config")
	//}
	return nil
}
