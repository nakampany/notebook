package services

type HelloService struct {
	HelloHandler func() string
}
