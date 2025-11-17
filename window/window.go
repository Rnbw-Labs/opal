// Package window contains a Window struct and relevant
// factories and functions.
package window

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"time"
)

// Window represents an application window.
type Window struct {
	width, height int
	title         string
	targetFPS     int
	vsync         bool
	internal      *glfw.Window
	lastOp        time.Time
}

// The New function generates and returns a new *Window.
func New(width, height int, title string) *Window {
	runtime.LockOSThread()
	// First, we initialize.
	if err := glfw.Init(); err != nil {
		// If it fails, not much we can do. LOL.
		panic(err)
	}

	// BTW, window, we are using version 4.6...
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	// ...core profile...
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// ...and let me resize pretty please!
	glfw.WindowHint(glfw.Resizable, glfw.True)

	// Whip up a new window and make it the current one.
	internalWindow, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	internalWindow.MakeContextCurrent()

	// Initialize OpenGL!
	if err := gl.Init(); err != nil {
		panic(err)
	}
	// Set the clear color.
	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	gl.Enable(gl.DEPTH_TEST)

	return &Window{
		width:    width,
		height:   height,
		title:    title,
		internal: internalWindow,
		vsync:    false,
	}
}

func (w *Window) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	w.lastOp = time.Now()
}

var lastX, lastY float32

func (w *Window) Swap() Input {
	w.internal.SwapBuffers()
	glfw.PollEvents()

	keys := make([]string, 5)
	for key, name := range glfwKeyToString {
		if w.IsKeyPressed(key) {
			keys = append(keys, name)
		}
	}
	elapsed := time.Since(w.lastOp)
	if w.targetFPS > 0 {
		sleepTime := (time.Second / time.Duration(w.targetFPS)) - elapsed
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
	x, y := w.internal.GetCursorPos()
	deltaX := float32(x) - lastX
	deltaY := float32(y) - lastY
	lastX = float32(x)
	lastY = float32(y)
	return Input{
		Keys: keys,
		Mouse: Mouse{
			X: deltaX,
			Y: deltaY,
		},
	}
}

func (w *Window) ShouldRun() bool {
	return !w.internal.ShouldClose()
}

func (w *Window) Close() {
	w.internal.SetShouldClose(true)
}

func (w *Window) SetVsync(on bool) {
	w.internal.MakeContextCurrent()
	if on {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	w.vsync = on
}

func (w *Window) SetTargetFPS(fps int) {
	w.targetFPS = fps
}

func (w *Window) ReleaseTargetFPS() {
	w.targetFPS = 0
}

func (w *Window) IsKeyPressed(key glfw.Key) bool {
	return w.internal.GetKey(key) == glfw.Press
}

func (w *Window) LockCursor() {
	w.internal.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
}

func (w *Window) UnlockCursor() {
	w.internal.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
}
