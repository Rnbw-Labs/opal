package renderer

import (
	"opal/camera"
	"opal/gpu"
	"opal/render_math"
)

type Renderer struct {
	Camera camera.Camera
}

func (r *Renderer) DrawMesh(mesh gpu.Mesh, shader gpu.Shader, texture gpu.Texture, modelTransformation render_math.Matrix4) {
	shader.Use()
	//shader.UseTexture("tex0", texture, 0)
	shader.SetMatrix4("model", modelTransformation)
	shader.SetMatrix4("view", r.Camera.ViewMatrix())
	shader.SetMatrix4("projection", r.Camera.ProjectionMatrix())
	mesh.Draw()
}
