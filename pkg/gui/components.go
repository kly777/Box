package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UIComponents struct {
	BoxList      *widget.List
	FileList     *widget.List
	CreateBoxBtn *widget.Button
}

func NewUIComponents(window fyne.Window, state *UIState) *UIComponents {
	// Box列表组件
	boxList := widget.NewList(
		func() int { return len(state.CurrentBoxes) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(state.CurrentBoxes[id].Name)
		},
	)

	// 文件列表组件
	fileList := widget.NewList(
		func() int { return len(state.GetCurrentFiles()) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(state.GetCurrentFiles()[id].Name)
		},
	)

	// 添加Box按钮
	createBoxBtn := widget.NewButton("添加Box", nil)

	return &UIComponents{
		BoxList:      boxList,
		FileList:     fileList,
		CreateBoxBtn: createBoxBtn,
	}
}

func (c *UIComponents) BuildLayout() fyne.CanvasObject {
	return container.NewBorder(
		container.NewHBox(c.CreateBoxBtn), // 顶部按钮
		nil, nil, nil,
		container.NewHSplit(
			c.BoxList,  // 左侧Box列表
			c.FileList, // 右侧文件列表
		),
	)
}
