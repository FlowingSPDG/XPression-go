package main

import (
	XPression "github.com/FlowingSPDG/XPression-go"
	"github.com/c-bata/go-prompt"
)


func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "CLFB",Description: "Clears framebuffer number buffer. For example, CLFB 0000 clears framebuffer 1."},
		{Text: "CLRA",Description: "Clears all framebuffers."},
		{Text: "CUE",Description: "Prepares take item takeid to go to air next in framebuffer number buffer on layer number layer. The take item is not taken to air, but is prepared to be taken to air without any frame delay. For example, CUE 3:2:-5 prepares to load the take item 3 into the framebuffer 3 and onto layer -5."},
		{Text: "DOWN",Description: "Move the current selection in the sequencer to the item below it in the list."},
		{Text: "FOCUS ",Description: "Set the sequencer focus to the take item number takeid. For example, FOCUS 0005 set the focus to take item 0005."},
		{Text: "GPI ",Description: "Trigger the simulated GPI input gpi. This is treated as if the GPI input were triggered externally. For example, GPI 5 triggers GPI input 5."},
		{Text: "LAYEROFF ",Description: "Takes a scene in framebuffer number buffer on layer number layer off air using the defined out transition. For example, LAYEROFF 0000:2 removes the scene on layer 2 of framebuffer 0000 (the first framebuffer)."},
		{Text: "NEXT",Description: "Take the current take item in the sequencer to air and advance the current selection to the next item in the list."},
		{Text: "READ",Description: "Take the current selection in the sequencer to air."},
		{Text: "RESUME ",Description: "Resumes all layers in framebuffer number buffer. For example, RESUME 0000 resumes all layers in framebuffer 1."},
		{Text: "SEQI",Description: "Loads the take item takeid to air on layer number layer to the output channel selected in the template. The Sequencer focus moves to this item. For example, SEQI 0005:7 loads the take item 0005 onto layer 7."},
		{Text: "SEQO",Description: "Takes the take item takeid off-air. For example, SEQO 0005 takes the template with TakeID 5 off-air."},
		{Text: "SWAP",Description: "Loads all the take items that are currently in the cued state to air in framebuffer number buffer. If a framebuffer is not specified, all cued take items in all framebuffers are taken to air. For example, SWAP 0 takes all the cued take items in framebuffer 1 to air."},
		{Text: "TAKE",Description: "Loads take item takeid to air in framebuffer number buffer on layer number layer. The Sequencer focus does not move to this item. For example, TAKE 5:0:7 loads the template with TakeID 5 into framebuffer 1 and onto layer 7."},
		{Text: "UNCUEALL",Description: "Removes all cued items from the cued state."},
		{Text: "UNCUE",Description: "Remove item with take id takeid from the cued state."},
		{Text:"UP",Description: "Move the current selection in the sequencer to the item above it in the list."},
		{Text: "UPNEXT",Description: "Sets the preview to the take item takeid in the sequencer without moving the focus bar."},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	xp, err := XPression.New("localhost",7788)
	if err != nil {
		panic(err)
	}

	for {
		cmd := prompt.Input("XPression >>> ", completer)
		if err := xp.Write(cmd); err != nil {
			panic(err)
		}
	}
}
