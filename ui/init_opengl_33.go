package ui

import "allece.com/system/core/go-gl/gl/v3.3-core/gl"

func InitOpenGL33() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}
