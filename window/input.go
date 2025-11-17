package window

import "github.com/go-gl/glfw/v3.3/glfw"

type Input struct {
	Keys  []string
	Mouse Mouse
}

func (i Input) Pushing(key string) bool {
	for _, k := range i.Keys {
		if k == key {
			return true
		}
	}
	return false
}

type Mouse struct {
	X, Y float32
}

func NewInput() Input {
	return Input{
		Keys: []string{},
		Mouse: Mouse{
			X: 0,
			Y: 0,
		},
	}
}

var glfwKeyToString = map[glfw.Key]string{
	// Letters
	glfw.KeyA: "A", glfw.KeyB: "B", glfw.KeyC: "C", glfw.KeyD: "D",
	glfw.KeyE: "E", glfw.KeyF: "F", glfw.KeyG: "G", glfw.KeyH: "H",
	glfw.KeyI: "I", glfw.KeyJ: "J", glfw.KeyK: "K", glfw.KeyL: "L",
	glfw.KeyM: "M", glfw.KeyN: "N", glfw.KeyO: "O", glfw.KeyP: "P",
	glfw.KeyQ: "Q", glfw.KeyR: "R", glfw.KeyS: "S", glfw.KeyT: "T",
	glfw.KeyU: "U", glfw.KeyV: "V", glfw.KeyW: "W", glfw.KeyX: "X",
	glfw.KeyY: "Y", glfw.KeyZ: "Z",

	// Numbers
	glfw.Key0: "0", glfw.Key1: "1", glfw.Key2: "2", glfw.Key3: "3",
	glfw.Key4: "4", glfw.Key5: "5", glfw.Key6: "6", glfw.Key7: "7",
	glfw.Key8: "8", glfw.Key9: "9",

	// Modifiers
	glfw.KeyLeftShift:    "LEFT_SHIFT",
	glfw.KeyRightShift:   "RIGHT_SHIFT",
	glfw.KeyLeftControl:  "LEFT_CONTROL",
	glfw.KeyRightControl: "RIGHT_CONTROL",
	glfw.KeyLeftAlt:      "LEFT_ALT",
	glfw.KeyRightAlt:     "RIGHT_ALT",

	// Punctuation / Symbols
	glfw.KeyComma:       ".",
	glfw.KeyPeriod:      ".",
	glfw.KeySemicolon:   ";",
	glfw.KeyApostrophe:  "'",
	glfw.KeySlash:       "/",
	glfw.KeyBackslash:   "\\",
	glfw.KeyMinus:       "-",
	glfw.KeyEqual:       "=",
	glfw.KeyGraveAccent: "`",

	// Special / navigation
	glfw.KeySpace:     "SPACE",
	glfw.KeyEnter:     "ENTER",
	glfw.KeyEscape:    "ESCAPE",
	glfw.KeyTab:       "TAB",
	glfw.KeyBackspace: "BACKSPACE",
	glfw.KeyLeft:      "LEFT",
	glfw.KeyRight:     "RIGHT",
	glfw.KeyUp:        "UP",
	glfw.KeyDown:      "DOWN",
	glfw.KeyDelete:    "DELETE",
	glfw.KeyHome:      "HOME",
	glfw.KeyEnd:       "END",
	glfw.KeyPageUp:    "PAGEUP",
	glfw.KeyPageDown:  "PAGEDOWN",
}
