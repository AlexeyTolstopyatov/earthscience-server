package app

type Handler interface {
	Instance()
}

type App struct {
	address string
	port    string
}
