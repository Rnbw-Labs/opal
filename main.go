package main

import (
	"github.com/rnbw-labs/opal/camera"
	"github.com/rnbw-labs/opal/cpu"
	"github.com/rnbw-labs/opal/gmath"
	gpu "github.com/rnbw-labs/opal/gpu"
	"github.com/rnbw-labs/opal/renderer"
	"github.com/rnbw-labs/opal/window"
	"os"
)

func main() {
	// Set up game window
	gameWindow := window.New(600, 400, "Test Window")
	gameWindow.SetVsync(true)
	gameWindow.LockCursor()

	newFragmentBytes, err := os.ReadFile("./gpu/shaders/basic_fragment.glsl")
	if err != nil {
		panic(err)
	}
	newFragmentSource := string(newFragmentBytes)

	// Compile basic shaders
	shader, err := gpu.NewShader(gpu.BasicVertexShaderSource, newFragmentSource)
	if err != nil {
		panic(err)
	}

	// Generate a fake cube mesh with size 1
	mesh := cpu.CubeMesh(1)

	// UploadMesh cube mesh to the GPU
	gpuMesh := gpu.UploadMesh(*mesh)

	// Generate a new basic camera and move it slightly
	myCamera := camera.NewBasicCamera(600, 400)
	myCamera.Position.Z = 3
	myCamera.Position.X = 1

	// Initialize a renderer
	myRenderer := renderer.Renderer{Camera: myCamera}

	texture, err := cpu.LoadPNG("./gpu/textures/test.png")
	if err != nil {
		panic(err)
	}
	gpuTexture := gpu.UploadTexture(texture)

	// Main render loop
	for gameWindow.ShouldRun() {
		// Wipe window clear
		gameWindow.Clear()
		// Draw my stuff to it
		myRenderer.DrawMesh(*gpuMesh, shader, *gpuTexture, gmath.Translate(gmath.Vector3{}))
		// Swap
		input := gameWindow.Swap()
		if input.Pushing("W") {
			myCamera.MoveForward(0.05)
		}
		if input.Pushing("S") {
			myCamera.MoveForward(-0.05)
		}
		if input.Pushing("D") {
			myCamera.MoveRight(0.05)
		}
		if input.Pushing("A") {
			myCamera.MoveRight(-0.05)
		}
		myCamera.Pitch(-input.Mouse.Y * 0.001)
		myCamera.Yaw(-input.Mouse.X * 0.001)
	}
}
