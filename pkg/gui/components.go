package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

type UIComponents struct {
	BoxList             *widget.List
	FileList            *widget.List
	CreateBoxBtn        *widget.Button
	CurrentBoxNameLabel *widget.Label
}

func NewUIComponents(window fyne.Window, state *UIState) *UIComponents {
	// Box列表组件
	boxList := widget.NewList(
		func() int { return len(state.CurrentBoxes) },
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			lbl.TextStyle = fyne.TextStyle{Bold: true} // 加粗显示
			return lbl
		},
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

	currentBoxNameLabel := widget.NewLabel("Box")

	return &UIComponents{
		BoxList:             boxList,
		FileList:            fileList,
		CreateBoxBtn:        createBoxBtn,
		CurrentBoxNameLabel: currentBoxNameLabel,
	}
}

func (c *UIComponents) BuildLayout() fyne.CanvasObject {
	// 创建HSplit并保存引用
	hsplit := container.NewHSplit(
		container.NewScroll(c.BoxList),
		container.NewScroll(c.FileList),
	)

	// 创建顶部带分割线的布局
	topContent := container.NewVBox(
		container.NewHBox(
			c.CreateBoxBtn,
			c.CurrentBoxNameLabel,
		),
		widget.NewSeparator(), // 添加水平分割线
	)

	// 使用HSplit构建border布局
	border := container.NewBorder(
		container.NewPadded(topContent),
		nil, nil, nil,
		hsplit, // 使用预先配置好的HSplit
	)
	return border
}
