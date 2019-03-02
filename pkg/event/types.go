package event

// Event is an event that is passed through to the state
type Event interface{}

// Keyboard holds the key or key combination of a keyboard event
type Keyboard struct {
	Key string
}

// Mouse holds the action and location of a mouse event
type Mouse struct {
	Action string
	X      int
	Y      int
}

// Resize holds the new width and height of a window after a resize
type Resize struct {
	Width  int
	Height int
}

// Error holds the error message from an error event
type Error struct {
	Msg string
}

// Quit is a quit event that signifies that the program should exit
type Quit struct{}
