package ui

import "allece.com/system/core/opengl/gl11/gl"

func initOpenGL11() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}
