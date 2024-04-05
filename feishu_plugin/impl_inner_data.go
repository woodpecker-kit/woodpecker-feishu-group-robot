package feishu_plugin

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
)

// GetInnerCopyRenderFeishuCardContext
// this data will remove timestamp and sign, must be json string
func (p *FeishuPlugin) GetInnerCopyRenderFeishuCardContext() string {
	return p.innerData.copyRenderFeishuCard
}

// saveDeepCopyInnerData
// do not forget to call save more item
func (p *FeishuPlugin) saveDeepCopyInnerData() error {
	var cpPlugin FeishuPlugin
	errCopy := struct_kit.DeepCopyByGob(*p, &cpPlugin)
	if errCopy != nil {
		return fmt.Errorf("copy FeishuPlugin err: %v", errCopy)
	}

	// remove timestamp and sign
	cpPlugin.SetWoodpeckerInfo(*p.woodpeckerInfo)
	cpPlugin.FeishuRobotMsgTemplate.Timestamp = 0
	cpPlugin.FeishuRobotMsgTemplate.Sign = ""

	if p.innerData.renderOssCardFileBrowser != nil {
		cpPlugin.innerData.renderOssCardFileBrowser = p.innerData.renderOssCardFileBrowser
	}

	renderCopyFeishuCard, errCopyRender := renderFeishuCardFromPlugin(&cpPlugin)
	if errCopyRender != nil {
		return errCopyRender
	}
	p.innerData.copyRenderFeishuCard = renderCopyFeishuCard
	return nil
}
