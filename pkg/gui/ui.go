package gui

import (
	"box/internal/core/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func CreateMainUI(window fyne.Window, boxService service.BoxService) fyne.CanvasObject {
	// 加载真实数据
	boxes, err := boxService.GetRootBoxes()
	if err != nil {
		// 错误处理
		errorLabel := widget.NewLabel("加载数据失败: " + err.Error())
		return container.NewCenter(errorLabel)
	}

	// 先定义Box列表组件
	boxList := widget.NewList(
		func() int { return len(boxes) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Box")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(boxes[id].Name)
		},
	)

	// 刷新Box列表的函数
	refreshBoxes := func() {
		newBoxes, err := boxService.GetRootBoxes()
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			boxes = newBoxes
			boxList.Refresh()
		}
	}

	// 添加Box按钮
	createBoxBtn := widget.NewButton("添加Box", func() {
		inputDialog := dialog.NewEntryDialog("新建Box", "输入Box名称：", func(name string) {
			if name != "" {
				_, err := boxService.CreateBox(name)
				if err != nil {
					dialog.ShowError(err, window)
				} else {
					refreshBoxes()
				}
			}
		}, window)
		inputDialog.Show()
	})

	// 文件列表（暂时空实现）
	fileList := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject {
			return widget.NewLabel("File")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {},
	)

	// 添加文件按钮
	addFileBtn := widget.NewButton("添加文件", func() {
		// TODO: 实现添加文件逻辑
	})

	// Box区域布局（包含创建按钮和列表）
	boxSection := container.NewBorder(
		createBoxBtn, // 顶部放置创建按钮
		nil,
		nil,
		nil,
		boxList,
	)

	// 主布局
	content := container.NewHSplit(
		boxSection,
		container.NewBorder(nil, addFileBtn, nil, nil, fileList),
	)

	return content
}
