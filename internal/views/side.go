package views

import "github.com/ryjen/imgui-go"

const sideViewID = "SideView"

var sideViewSize = imgui.Vec2{
	X: 200, Y: 440,
}

type SideView struct {
	players *PlayersView
	dice *DiceView
	chat *ChatView
}

func NewSideView() *SideView {
	return &SideView{
		players: NewPlayersView(),
		dice: NewDiceView(),
		chat: NewChatView(),
	}
}

func (view *SideView) Render() {

	imgui.SameLine()

	if imgui.BeginChildV(sideViewID, sideViewSize, false, imgui.WindowFlagsAlwaysAutoResize) {

		view.players.Render()
		view.dice.Render()
		view.chat.Render()
	}

	imgui.EndChild()
}

func (view *SideView) Update() {
	view.players.Update()
	view.dice.Update()
	view.chat.Update()
}