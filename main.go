package main

import (
  "os"
  "fmt"
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
    port := fmt.Sprintf(":%s", os.Getenv("PORT"))
    h := server.Default(
      server.WithHostPorts(port),
    )

    h.GET("/", func(c context.Context, ctx *app.RequestContext) {
            ctx.JSON(consts.StatusOK, utils.H{"ip": ctx.ClientIP()})
    })

    h.Spin()
}
