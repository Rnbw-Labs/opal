package gpu

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"opal/cpu"
)

func UploadMesh(uploading cpu.Mesh) *Mesh {
	var vao, vbo, ebo uint32

	// First, we generate the vertex array object.
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Upload vertex data.
	vertexFloats := make([]float32, 0, len(uploading.Vertices)*8)
	for _, vertex := range uploading.Vertices {
		vertexFloats = append(vertexFloats, vertex.Unpack()...)
	}
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexFloats)*4, gl.Ptr(vertexFloats), gl.STATIC_DRAW)

	// Upload index data.
	indices := make([]uint32, 0, len(uploading.Triangles)*3)
	for _, triangle := range uploading.Triangles {
		indices = append(indices, triangle.Indices[0], triangle.Indices[1], triangle.Indices[2])
	}
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Vertex attribute pointers.
	stride := int32(8 * 4) // 8 floats per vertex
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, stride, 0)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, stride, 12)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, stride, 24)
	gl.EnableVertexAttribArray(2)

	// Cleanup.
	gl.BindVertexArray(0)

	return &Mesh{
		VAO:        vao,
		VBO:        vbo,
		EBO:        ebo,
		IndexCount: uint32(len(indices)),
	}
}

func UploadTexture(t *cpu.Texture) *Texture {
	var id uint32
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(t.Width),
		int32(t.Height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(t.Pixels),
	)

	gl.GenerateMipmap(gl.TEXTURE_2D)
	return &Texture{ID: id}
}
