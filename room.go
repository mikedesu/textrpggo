package main

type Room struct {
    name string
    description string

    // add exits
    exits []int
}



func roomCreateNewRoom(name string) Room {
    var r Room
    r.name = name
    return r
}


func roomAddDescription(r *Room, description string) {
    r.description = description
}


func roomAddExit(r *Room, exit int) {
    r.exits = append(r.exits, exit)
}

