package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/common"
	"github.com/go-chassis/go-chassis/core/lager"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/client/

// Implement grayscale publishing of the application,version  A is you old service ,version B is you
// new service.you want to small request to access you new service to test new service of new function

func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}

	for i := 0; i < 10; i++ {
		req, err := rest.NewRequest("POST", "cse://ROUTERServer/equal")
		if err != nil {
			lager.Logger.Error("new request failed.")
			return
		}
		parm := struct {
			Num  int
			Nums []int
		}{
			Num:  10,
			Nums: []int{2, 5, 3},
		}

		parmByte, _ := json.Marshal(parm)
		req.SetBody(parmByte)

		//req.SetHeader("Chassis", "info")
		defer req.Close()
		ctx := context.WithValue(context.TODO(), common.ContextHeaderKey{}, map[string]string{
			"user": "peter",
		})
		resp, err := core.NewRestInvoker().ContextDo(ctx, req)
		if err != nil {
			lager.Logger.Error("do request failed.")
			return
		}
		defer resp.Close()
		lager.Logger.Info("ROUTER Server equal num [POST]: " + string(resp.ReadBody()))

		time.Sleep(1 * time.Second)
	}
}
