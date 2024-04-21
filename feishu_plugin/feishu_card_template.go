package feishu_plugin

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/resource"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
	"strings"
)

const (
	headTemplateStyleDefault = "default"
	headTemplateStyleGreen   = "green"
	headTemplateStyleRed     = "red"
	headTemplateStyleOrange  = "orange"
	headTemplateStylIndigo   = "indigo"
)

func renderFeishuCardFromPlugin(p *FeishuPlugin) (string, error) {
	renderPlugin, errFormat := formatRenderPluginData(*p)
	if errFormat != nil {
		return "", errFormat
	}

	noticeList, filterBuildStatus, errNoticeList := parserCardElementNoticeList(&renderPlugin)
	if errNoticeList != nil {
		return "", errNoticeList
	}

	var renderRoot FeishuCardRenderRoot
	renderRoot.BuildStatus = filterBuildStatus
	renderRoot.FeishuRobotMsgTemplate = p.FeishuRobotMsgTemplate

	head, errParseHead := parserFeishuCardHeader(renderPlugin, filterBuildStatus)
	if errParseHead != nil {
		return "", errParseHead
	}
	renderRoot.FeishuCardHeader = head

	renderRoot.CardElementsHead = CardElementsHead{
		CardTitle: string_tools.String2LineRaw(renderPlugin.Settings.Title),
	}

	renderRoot.CardElementsNoticeList = noticeList

	cardElementTail := parserCardElementTail(renderPlugin.ShortInfo())
	renderRoot.CardElementsTail = cardElementTail

	cardNoteTail := parserCardNoteTail(renderPlugin)
	if cardNoteTail.LogoAltStr == "" {
		cardNoteTail.LogoAltStr = p.FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoAltStr
	}
	if cardNoteTail.LogoImgKey == "" {
		cardNoteTail.LogoImgKey = p.FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoImgKey
	}
	renderRoot.CardNoteTail = cardNoteTail

	tplFeishuCardHead, errTplFeishuCardHead := resource.FetchFeishuCardTemplateTplByName(resource.ItemFeishuCardTemplateHeadTpl, p.Settings.I18nLangSet)
	if errTplFeishuCardHead != nil {
		return "", errTplFeishuCardHead
	}

	headContent, errRenderRootHead := wd_template.Render(tplFeishuCardHead, &renderRoot)
	if errRenderRootHead != nil {
		return "", errRenderRootHead
	}

	tplFeishuCardTail, errTplFeishuCardTail := resource.FetchFeishuCardTemplateTplByName(resource.ItemFeishuCardTemplateTailTpl, p.Settings.I18nLangSet)
	if errTplFeishuCardTail != nil {
		return "", errTplFeishuCardTail
	}

	tailContent, errRenderRootTail := wd_template.Render(tplFeishuCardTail, &renderRoot)
	if errRenderRootTail != nil {
		return "", errRenderRootTail
	}

	// append to totalMsg
	totalMsg := fmt.Sprintf("%s%s%s",
		headContent,
		strings.Join(noticeList, ""),
		tailContent,
	)
	return totalMsg, nil
}

func parserCardNoteTail(p FeishuPlugin) CardNoteTail {
	info := p.ShortInfo()
	return CardNoteTail{
		CiSystemVersion: info.System.Version,
		CiSystemHost:    info.System.Hostname,
		LogoImgKey:      p.Settings.PoweredByImageKey,
		LogoAltStr:      p.Settings.PoweredByImageAlt,
	}
}

func parserCardElementTail(short wd_short_info.WoodpeckerInfoShort) CardElementsTail {
	var cardElementsTail CardElementsTail
	cardElementsTail.BuildTimeCreated = short.Build.CreatedAt
	cardElementsTail.BuildTimeTotal = short.Build.DurationHuman
	cardElementsTail.RunnerMachine = short.System.RunnerMachine
	cardElementsTail.RunnerPlatform = short.System.RunnerPlatform
	cardElementsTail.BuildEvent = short.Build.Event
	return cardElementsTail
}

func parserFeishuCardHeader(p FeishuPlugin, status string) (FeishuCardHeader, error) {
	var feishuCardHeader FeishuCardHeader
	if !string_tools.StringInArr(status, supportRenderStatus) {
		return feishuCardHeader, fmt.Errorf("status %s not in supportRenderStatus %v", status, supportRenderStatus)
	}

	shortInfo := p.ShortInfo()

	var headStyle string
	var headerTitle = shortInfo.Repo.FullName

	if status == wd_info.BuildStatusSuccess {
		switch shortInfo.Build.Event {
		default:
			headStyle = headTemplateStyleGreen
		case wd_info.EventPipelinePullRequest:
			headStyle = headTemplateStyleOrange
			headerTitle = fmt.Sprintf("%s PR: #%s", headerTitle, shortInfo.Commit.PR)
		case wd_info.EventPipelinePullRequestClose:
			headStyle = headTemplateStyleOrange
			headerTitle = fmt.Sprintf("%s ClosePR: #%s", headerTitle, shortInfo.Commit.PR)
		case wd_info.EventPipelineTag:
			headStyle = headTemplateStylIndigo
			headerTitle = fmt.Sprintf("%s Tag: %s", headerTitle, shortInfo.Commit.Tag)
		case wd_info.EventPipelineRelease:
			headStyle = headTemplateStylIndigo
			headerTitle = fmt.Sprintf("%s Relesae: %s", headerTitle, shortInfo.Commit.Sha)
		case wd_info.EventPipelineCron:
			headStyle = headTemplateStyleDefault
			headerTitle = fmt.Sprintf("%s Cron: %s", headerTitle, shortInfo.Build.Number)
		}
	} else {
		headStyle = headTemplateStyleRed
		switch shortInfo.Build.Event {
		default:
			headerTitle = fmt.Sprintf("[Failure] %s", headerTitle)
		case wd_info.EventPipelinePullRequest:
			headerTitle = fmt.Sprintf("[Failure] %s PR: #%s", headerTitle, shortInfo.Commit.PR)
		case wd_info.EventPipelinePullRequestClose:
			headerTitle = fmt.Sprintf("[Failure] %s ClosePR: #%s", headerTitle, shortInfo.Commit.PR)
		case wd_info.EventPipelineTag:
			headerTitle = fmt.Sprintf("[Failure] %s Tag: %s", headerTitle, shortInfo.Commit.Tag)
		case wd_info.EventPipelineRelease:
			headerTitle = fmt.Sprintf("[Failure] %s Relesae: %s", headerTitle, shortInfo.Commit.Sha)
		case wd_info.EventPipelineCron:
			headerTitle = fmt.Sprintf("[Failure] %s Cron: %s", headerTitle, shortInfo.Build.Number)
		}
	}

	if shortInfo.CurrentWorkflow.Name != "" {
		headerTitle = fmt.Sprintf("%s ➜ %s", headerTitle, shortInfo.CurrentWorkflow.Name)
	}

	feishuCardHeader.HeadTemplateStyle = headStyle
	feishuCardHeader.HeaderTitle = headerTitle

	return feishuCardHeader, nil
}

func parserCardElementNoticeList(f *FeishuPlugin) ([]string, string, error) {
	var noticeList []string
	buildStatus := f.ShortInfo().Build.Status

	// check notice type
	isHasBuildStatus := false

	if len(f.Settings.NoticeTypes) == 0 {
		buildStatus = wd_info.BuildStatusFailure
		return noticeList, buildStatus, fmt.Errorf("no notice type in FeishuPlugin.Settings.NoticeTypes")
	}

	var lessNoticeType []string
	for _, noticeType := range f.Settings.NoticeTypes {
		if !(string_tools.StringInArr(noticeType, noticeTypeSupport)) {
			buildStatus = wd_info.BuildStatusFailure
			return noticeList, buildStatus, fmt.Errorf("config [ feishu-notice-types ] only support %v", noticeTypeSupport)
		}

		// build_status special treatment required
		if noticeType != NoticeTypeBuildStatus {
			lessNoticeType = append(lessNoticeType, noticeType)
		} else {
			isHasBuildStatus = true
		}
	}

	// render notice type
	var lessNoticeContent []string
	for _, noticeType := range lessNoticeType {
		switch noticeType {
		case NoticeTypeFileBrowser: // NoticeTypeFileBrowser
			if f.innerData.renderOssCardFileBrowser != nil {
				if !f.innerData.renderOssCardFileBrowser.CardOssFB.IsSendSuccess {
					buildStatus = wd_info.BuildStatusFailure
				}
				cardFileBrowserContent, errRenderOssCardFileBrowser := renderOssCardFileBrowser(f.innerData.renderOssCardFileBrowser, f.Settings.I18nLangSet)
				if errRenderOssCardFileBrowser != nil {
					return noticeList, buildStatus, errRenderOssCardFileBrowser
				}
				lessNoticeContent = append(lessNoticeContent, cardFileBrowserContent)
			}
		}
	}

	if isHasBuildStatus {
		buildStatusContent, err := renderBuildStatus(*f, buildStatus)
		if err != nil {
			return nil, buildStatus, err
		}
		// let build_status be first
		noticeList = []string{buildStatusContent}
		noticeList = append(noticeList, lessNoticeContent...)
	} else {
		noticeList = lessNoticeContent
	}

	return noticeList, buildStatus, nil
}

func formatRenderPluginData(f FeishuPlugin) (FeishuPlugin, error) {
	var p FeishuPlugin
	woodPeckerInfo := f.GetWoodPeckerInfo()
	errFeishuPluginMember2LineRaw := string_tools.StructMemberString2LineRaw(&woodPeckerInfo)
	if errFeishuPluginMember2LineRaw != nil {
		return p, errFeishuPluginMember2LineRaw
	}

	if woodPeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl != "" {
		woodPeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl = strings.Replace(
			woodPeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl,
			"compare/0000000000000000000000000000000000000000...", "commit/", -1,
		)
	}
	if woodPeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitUrl != "" {
		woodPeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitUrl = strings.Replace(
			woodPeckerInfo.PreviousInfo.PreviousCommitInfo.CiPreviousCommitUrl,
			"compare/0000000000000000000000000000000000000000...", "commit/", -1,
		)
	}
	p = FeishuPlugin{
		Name:           f.Name,
		Version:        f.Version,
		woodpeckerInfo: &woodPeckerInfo,
		Settings:       f.Settings,

		innerData: f.innerData,
	}

	p.SetWoodpeckerInfo(woodPeckerInfo)
	return p, nil
}
