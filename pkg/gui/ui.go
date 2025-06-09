package gui

import (
	"box/internal/core/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// 存储当前显示的Boxes

func StartGUI(boxService service.BoxService) {
	app := app.New()
	window := app.NewWindow("Box Management")

	// 初始化状态管理
	state := NewUIState(boxService, window)

	// 初始化界面组件
	components := NewUIComponents(window, state)

	// 设置事件处理
	SetupEventHandlers(components, state, window)

	// 设置布局
	window.SetContent(components.BuildLayout())

	// 初始加载
	state.RefreshAll()

	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
