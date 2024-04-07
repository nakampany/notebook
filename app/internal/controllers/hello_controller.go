package controllers

import (
	"io"
	"net/http"
	"note/app/internal/controllers/services"
)

type HelloController struct {
	service services.HelloService
}

// GET /hello のハンドラ
func (c *HelloController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
