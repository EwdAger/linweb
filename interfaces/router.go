package interfaces

type HandlerFunc func(IContext)

type IRouter interface {
	Handle(c IContext)
	Get(part string, handler HandlerFunc)
	Post(part string, handler HandlerFunc)
}