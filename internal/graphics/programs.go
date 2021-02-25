package graphics

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/sirupsen/logrus"
	"strings"
)

// todo move opengl shader scripts to separate files
// todo refactoring directories

type Programs struct {
	SimpleColor uint32
	Light       uint32
}

const lightVertexShader string = `
    #version 410
	layout (location = 0) in vec3 vert;
    uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;
    void main() {
        gl_Position = projection * view * model * vec4(vert, 1.0);
    }
` + "\x00"

const lightFragmentShader string = `
    #version 410
	out vec4 FragColor;
	
	void main()
	{
		FragColor = vec4(1.0); // set alle 4 vector values to 1.0
	}
` + "\x00"

const vertexShader string = `
    #version 410
	layout (location = 0) in vec3 aPos;
	layout (location = 1) in vec3 aNormal;
	
	out vec3 FragPos;
	out vec3 Normal;
	
	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;
	
	void main()
	{
		FragPos = vec3(model * vec4(aPos, 1.0));
		Normal = mat3(transpose(inverse(model))) * aNormal;  
		
		gl_Position = projection * view * vec4(FragPos, 1.0);
	}
` + "\x00"

const fragmentShader string = `
    #version 410
	out vec4 FragColor;

	in vec3 Normal;  
	in vec3 FragPos;  
	  
	uniform vec3 lightPos; 
	uniform vec3 viewPos; 
	uniform vec3 lightColor;
	uniform vec3 objectColor;
	
	void main()
	{
		// ambient
		float ambientStrength = 0.01;
		vec3 ambient = ambientStrength * lightColor;
		
		// diffuse 
		vec3 norm = normalize(Normal);
		vec3 lightDir = normalize(lightPos - FragPos);
		float diff = max(dot(norm, lightDir), 0.0);
		vec3 diffuse = diff * lightColor;
		
		// specular
		float specularStrength = 0.5;
		vec3 viewDir = normalize(viewPos - FragPos);
		vec3 reflectDir = reflect(-lightDir, norm);  
		float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
		vec3 specular = specularStrength * spec * lightColor;  
			
		vec3 result = (ambient + diffuse + specular) * objectColor;
		FragColor = vec4(result, 1.0);
	} 
` + "\x00"

func NewPrograms() *Programs {
	p := &Programs{}
	p.SimpleColor = p.buildProgram(vertexShader, fragmentShader)
	p.Light = p.buildProgram(lightVertexShader, lightFragmentShader)

	return p
}

func (c *Programs) buildProgram(vertex, fragment string) uint32 {
	vertexShader, err := c.compileShader(vertex, gl.VERTEX_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	fragmentShader, err := c.compileShader(fragment, gl.FRAGMENT_SHADER)
	if err != nil {
		logrus.WithError(err).Error("error on compile vertex shader")
		return 0
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return prog
}

func (c *Programs) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
