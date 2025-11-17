package renderer

import (
	"github.com/rnbw-labs/opal/camera"
	"github.com/rnbw-labs/opal/gmath"
	"github.com/rnbw-labs/opal/gpu"
)

type Renderer struct {
	Camera camera.Camera
}

func (r *Renderer) DrawMesh(mesh gpu.Mesh, shader gpu.Shader, textures map[string]*gpu.Texture, uniforms map[string]any) {
	shader.Use()

	// Bind textures
	i := 0
	for slot, tex := range textures {
		shader.UseTexture(slot, *tex, i)
		i++
	}

	// Set uniforms
	for name, value := range uniforms {
		switch v := value.(type) {
		case gmath.Matrix4:
			shader.SetMatrix4(name, v)
		case float32:
			shader.SetFloat(name, v)
		case gmath.Vector3:
			shader.SetVector3(name, v)
		case int32:
			shader.SetInt(name, v)
		}
	}

	mesh.Draw()
}
