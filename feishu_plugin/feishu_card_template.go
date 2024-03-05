package feishu_plugin

import (
	"encoding/json"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	_ "github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
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

// DefaultCardTemplate
// use FeishuPlugin
// - feishu_message.FeishuRobotMsgTemplate
// - feishu_plugin.Config
// - wd_info.WoodpeckerInfo
const DefaultCardTemplate string = `{
  "timestamp": {{ FeishuRobotMsgTemplate.Timestamp }},
  "sign": "{{ FeishuRobotMsgTemplate.Sign }}",
  "msg_type": "interactive",
  "card": {
    "config": {
      "enable_forward": {{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.EnableForward }}
    },
    "header": {
      "template": "{{ Config.CardOss.HeadTemplateStyle }}",
      "title": {
        "tag": "plain_text",
        "content": "{{#failure WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus}}[Failure]{{/failure}}{{ WoodpeckerInfo.RepositoryInfo.CIRepo }}{{#success Config.CardOss.InfoTagResult }} Tag: {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag }}{{/success}}{{#success Config.CardOss.InfoPullRequestResult }} PullRequest: #{{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest }}{{/success}}"
      }
    },
    "elements": [
      {
        "tag": "markdown",
        "content": "**{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.CardTitle }}**"
      },
      {
        "tag": "hr"
      },
{{#success Config.CardOss.InfoTagResult }}
      {
        "tag": "markdown",
        "content": "📦 **Tag:** {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag }}\nCommitCode: {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSha }}"
      },
{{/success}}
{{#failure Config.CardOss.InfoTagResult }}
      {
        "tag": "markdown",
        "content": "{{#success Config.CardOss.InfoPullRequestResult }}🏗️ Pull Request: {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch }} -> {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch }} [#{{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest }}]({{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl }}){{/success}}{{#failure Config.CardOss.InfoPullRequestResult }}📝 Commit by {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitAuthor }} on **{{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitBranch }}**\nCommitCode: {{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitSha }}{{/failure}}"
      },
{{/failure}}
      {
        "tag": "markdown",
        "content": "{{#success WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus }}✅{{/success}}{{#failure WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus}}❌{{/failure}} Build [#{{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineNumber }}]({{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl }}) {{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitMessage }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl }}) | [See Build Details]({{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl }})"
      },
      {
        "tag": "hr"
      },
{{#success Config.RenderOssCard }}
{{#success Config.CardOss.InfoSendResult }}
      {
        "tag": "markdown",
        "content": "[OSS {{ Config.CardOss.InfoUser }} ]({{ Config.CardOss.Host }})\nPath: {{ Config.CardOss.InfoPath }}\nPage: [{{ Config.CardOss.PageUrl }}]({{ Config.CardOss.PageUrl }}){{#failure Config.CardOss.RenderResourceUrl }}\nPassword: {{ Config.CardOss.PagePasswd }}\n{{/failure}}{{#success Config.CardOss.RenderResourceUrl }}\nDownload: [click me]({{ Config.CardOss.ResourceUrl }})\n{{/success}}"
      },
{{/success}}
{{#failure Config.CardOss.InfoSendResult }}
      {
        "tag": "markdown",
        "content": "[OSS {{ Config.CardOss.InfoUser }} ]({{ Config.CardOss.Host }}) send error, please check at [build Details]({{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl }})"
      },
{{/failure}}
      {
        "tag": "hr"
      },
{{/success}}
      {
        "tag": "markdown",
        "content": "**Build Created At:** {{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineCreatedT }}\n**Build Time:** {{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineDurationHuman }}\nMachine: {{ WoodpeckerInfo.CiSystemInfo.CiMachine }}\nPlatform: {{ WoodpeckerInfo.CiSystemInfo.CiSystemPlatform }}\nEvent: {{ WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent }}"
      },
      {
        "tag": "hr"
      },
      {
        "tag": "note",
        "elements": [
          {
            "tag": "plain_text",
            "content": "From {{ WoodpeckerInfo.CiSystemInfo.CiSystemVersion }}@{{WoodpeckerInfo.CiSystemInfo.CiSystemHost}}. Powered By"
          },
          {
            "tag": "img",
            "img_key": "{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoImgKey }}",
            "alt": {
              "tag": "plain_text",
              "content": "{{ FeishuRobotMsgTemplate.CtxTemp.CardTemp.LogoAltStr }}"
            }
          }
        ]
      }
    ]
  }
}`

func formatRenderData(f *FeishuPlugin) error {
	errFeishuPluginMember2LineRaw := string_tools.StructMemberString2LineRaw(f)
	if errFeishuPluginMember2LineRaw != nil {
		return errFeishuPluginMember2LineRaw
	}

	// set default CardOss.HeadTemplateStyle
	f.Config.CardOss.HeadTemplateStyle = headTemplateStyleDefault
	if f.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus == wd_info.BuildStatusSuccess {
		f.Config.CardOss.HeadTemplateStyle = headTemplateStyleGreen
	}

	// check out f.Config.CardOss.InfoTagResult
	if f.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitTag == "" {
		f.Config.CardOss.InfoTagResult = RenderStatusHide
	} else {
		f.Config.CardOss.InfoTagResult = RenderStatusShow
		// fix WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl compare not support, when tags Link get error
		f.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl = strings.Replace(f.WoodpeckerInfo.CurrentInfo.CurrentCommitInfo.CiCommitUrl, "compare/0000000000000000000000000000000000000000...", "commit/", -1)
		f.Config.CardOss.HeadTemplateStyle = headTemplateStylIndigo
	}

	// set default InfoPullRequestResult
	f.Config.CardOss.InfoPullRequestResult = RenderStatusHide
	if f.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent == wd_info.EventPipelinePullRequest ||
		f.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent == wd_info.EventPipelinePullRequestClose {
		f.Config.CardOss.InfoPullRequestResult = RenderStatusShow
		f.Config.CardOss.HeadTemplateStyle = headTemplateStyleOrange
	}

	// set HeadTemplateStyle at failure
	if f.WoodpeckerInfo.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus == wd_info.BuildStatusFailure {
		f.Config.CardOss.HeadTemplateStyle = headTemplateStyleRed
	}

	return nil
}

func RenderFeishuCardByPlugin(tpl string, p *FeishuPlugin) (string, error) {
	var renderPlugin FeishuPlugin
	err := deepCopyByPlugin(p, &renderPlugin)
	if err != nil {
		return "", err
	}
	errFormat := formatRenderData(&renderPlugin)
	if errFormat != nil {
		return "", errFormat
	}

	message, err := wd_template.RenderTrim(tpl, &renderPlugin)
	if err != nil {
		return "", err
	}
	return message, nil
}

func deepCopyByPlugin(src, dst *FeishuPlugin) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}
