package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

func newPrimitive(t string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(t)
}

var inputFieldWidth = 64
var app = tview.NewApplication()
var text = tview.NewTextView().
	SetText("Welcome\n(q) to quit")
var inputField = tview.NewInputField().
	SetLabel("Command: ").
	SetFieldWidth(inputFieldWidth).
	SetAcceptanceFunc(tview.InputFieldMaxLength(64))
var grid = tview.NewGrid().
	SetRows(3, 0, 3).
	SetBorders(false).
	AddItem(newPrimitive("evildojo text rpg"), 0, 0, 1, 1, 0, 0, false).
	AddItem(text, 1, 0, 1, 1, 0, 0, false).
	AddItem(inputField, 2, 0, 1, 1, 0, 0, true)

var currentRoomIndex int = 0
var roomList []Room = []Room{}

// a map of commands that we have registered
// string to bool
var commandMap = make(map[string]bool)

func handleQuit(t string) {
	if t == "q" || t == "quit" || t == "exit" {
		app.Stop()
	}
}

func processHelp() string {
	var s string = "Commands:\n"
	for k, _ := range commandMap {
		// grab the first letter of k and wrap it in []
		s = s + "[" + string(k[0]) + "]" + k[1:] + "\n"
		//s = s + k + "\n"
	}
	return s
}

func processGoto(param int) string {
	for _, v := range roomList[currentRoomIndex].exits {
		if v == param {
			currentRoomIndex = param
			var s string = "You are now in " + roomList[currentRoomIndex].name + ".\n" + roomList[currentRoomIndex].description
			return s
		}
	}
	return "Invalid index supplied."

}

func processCommand(t string) string {
	// split t into words
	var words []string = strings.Split(t, " ")

	if words[0] == "look" || words[0] == "l" {
		return roomList[currentRoomIndex].description
	} else if words[0] == "goto" || words[0] == "g" {
		if len(words) > 1 {
			// if the second word is a room index
			//var param int = stringToInt(words[1])
			var param int = -1
			param, err := strconv.Atoi(words[1])
			if err != nil {
				return "Invalid index supplied."
			}

			if param != -1 {
				return processGoto(param)
			} else {
				return "I don't understand that command."
			}
		}
		var s string = ""
		s = "Where do you want to go?\n"
		for _, v := range roomList[currentRoomIndex].exits {
			for j, v2 := range roomList {
				if roomList[v].name == v2.name {
					s = s + strconv.Itoa(j) + " " + v2.name + "\n"
				}
			}
		}
		return s
	} else if t == "help" || t == "h" || t == "?" {
		return processHelp()
	}
	return "I don't understand that command."
}

func safeProcessCommand(t string) string {
	// lowercase t
	t = strings.ToLower(t)
	// if t is a key in commandMap that is true
	if commandMap[t] {
		return processCommand(t)
	} else {
		return "I don't understand that command."
	}
}

func inputFieldDoneFunction(key tcell.Key) {
	t := inputField.GetText()
	// lowercase t
	t = strings.ToLower(t)
	if t == "q" || t == "quit" || t == "exit" {
		app.Stop()
	} else {
		inputField.SetText("")
		text.SetText(processCommand(t))
	}
}

func registerCommands() {
	commandMap["quit"] = true
	commandMap["help"] = true
	commandMap["look"] = true
	commandMap["goto"] = true
}

func main() {
	room0 := roomCreateNewRoom("Starting Room")
	roomAddDescription(&room0, "An empty room.")
	roomAddExit(&room0, 1)

	room1 := roomCreateNewRoom("Command Room")
	roomAddDescription(&room1, "A room with a list of commands.")
	roomAddExit(&room1, 0)

	roomList = append(roomList, room0)
	roomList = append(roomList, room1)

	registerCommands()
	app.SetRoot(grid, true)
	inputField.SetDoneFunc(inputFieldDoneFunction)
	app.Run()
}
