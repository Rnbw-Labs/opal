package gpu

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"opal/render_math"
)

type Shader struct {
	ID uint32
}

func (s Shader) Use() {
	gl.UseProgram(s.ID)
}

func (s Shader) SetMatrix4(name string, mat render_math.Matrix4) {
	loc := gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(loc, 1, false, &mat.M[0])
}

func (s Shader) SetVector3(name string, v render_math.Vector3) {
	loc := gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
	gl.Uniform3f(loc, v.X, v.Y, v.Z)
}

func (s Shader) SetFloat(name string, f float32) {
	loc := gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
	gl.Uniform1f(loc, f)
}

func (s Shader) SetInt(name string, f int32) {
	loc := gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
	gl.Uniform1i(loc, f)
}

func (s Shader) UseTexture(name string, tex Texture, slot int) {
	tex.Bind(uint32(slot))
	s.SetInt(name, int32(slot))
}

func NewShader(vertexSrc, fragmentSrc string) (Shader, error) {
	vs, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		return Shader{}, err
	}
	fs, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return Shader{}, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLen)

		log := make([]byte, logLen+1)
		gl.GetProgramInfoLog(program, logLen, nil, &log[0])

		return Shader{}, fmt.Errorf("program link error: %s", log)
	}
	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	return Shader{ID: program}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	// Check status
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLen)

		log := make([]byte, logLen+1)
		gl.GetShaderInfoLog(shader, logLen, nil, &log[0])

		return 0, fmt.Errorf("shader compile error: %s", log)
	}

	return shader, nil
}
