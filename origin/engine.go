package origin

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"github.com/q676718008/regin/base"
)

type EngineDispatcher struct {
	origin   *gin.Engine
	response *ResponseHandler
}

// 定义RouterHandler
var Engine *EngineDispatcher

func init() {
	Engine = &EngineDispatcher{
		origin:   gin.Default(),
		response: Response,
	}
}

// Run HttpServer
func (ed *EngineDispatcher) HttpServer(server base.WebServer) {
	ed.origin.Any("/:module/:action", func(c *gin.Context) {
		// Error catch.
		defer func() {
			if err := server.GetError(); err != nil{
				result := base.ResultInvoker.CreateJson(200,"")
				result.SetData("code",10000)
				result.SetData("msg",err.Error())
				ed.response.Output(c,result)
			}
		}()
		defer server.ErrorCatch()
		ed.response.Output(c, server.Work(base.RequestInvoker.Factory(c)))
	})
	ed.origin.Run(server.Addr()) // Run HttpServer
}

// Run HttpsServer
func (ed *EngineDispatcher) HttpsServer(server base.WebServer) {
	// Register middleware.
	ed.origin.Use(func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "127.0.0.1:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			panic(err.Error())
		}
		c.Next()
	})
	// Register router.
	ed.origin.Any("/:module/:action", func(c *gin.Context) {
		request := base.RequestInvoker.Factory(c)
		result := server.Work(request)
		ed.response.Output(c, result)
	})
	// Run server.
	ed.origin.RunTLS("127.0.0.1:8080", "cert.pem", "key.pem")
}
