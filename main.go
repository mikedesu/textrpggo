package main

import (
    "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
    "strings"
    "strconv"
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
    SetBorders(true).
    AddItem(newPrimitive("Header"), 0, 0, 1, 1, 0, 0, false).
    AddItem(text, 1, 0, 1, 1, 0, 0, false).
    AddItem(inputField, 2, 0, 1, 1, 0, 0, true)


var currentRoomIndex int = 0
var roomList []Room = []Room{
    Room{"start", "An empty room with a single door."},
    Room{"room2", "A room with a table and a chair."},
}


// a map of commands that we have registered
// string to bool
var commandMap = make(map[string]bool)


func handleQuit(t string) {
    if t == "q" || t == "quit" || t=="exit" {
        app.Stop()
    }
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

            if param != -1 && param < len(roomList) {
                currentRoomIndex = param
                return roomList[currentRoomIndex].description
            } else {
                return "I don't understand that command."
            }
        }
        return "Where do you want to go?"
    } else if t == "help" || t == "h" || t == "?" {
        var s string = "Commands:\n"
        for k, _ := range commandMap {
            s = s + k + "\n"
        }
        return s
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
    if t == "q" || t == "quit" || t=="exit" {
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
    registerCommands()
    app.SetRoot(grid, true)
    inputField.SetDoneFunc(inputFieldDoneFunction)
    app.Run()
}

