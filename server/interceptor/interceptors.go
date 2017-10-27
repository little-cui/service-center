//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package interceptor

import (
	"github.com/ServiceComb/service-center/pkg/chain"
	roa "github.com/ServiceComb/service-center/pkg/rest"
	"github.com/ServiceComb/service-center/pkg/util"
	"net/http"
	"reflect"
	"runtime"
)

type InterceptorFunc func(http.ResponseWriter, *http.Request) error

func (f InterceptorFunc) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

type Interceptor struct {
	function InterceptorFunc
}

// Invoke performs the given interception.
// val is a pointer to the App Controller.
func (i *Interceptor) Handle(inv *chain.Invocation) {
	ctx := inv.HandlerContext()
	err := i.function(ctx[roa.CTX_SERVER_RESP].(http.ResponseWriter), ctx[roa.CTX_SERVER_REQ].(*http.Request))
	if err != nil {
		inv.Fail(err)
		return
	}
	inv.Next()
}

// InterceptFunc installs a general interceptor.
// This can be applied to any Controller.
// It must have the signature of:
//   func example(c *revel.Controller) revel.Result
func RegisterInterceptFunc(intc InterceptorFunc) {
	chain.RegisterHandler(roa.SERVER_HANDLER_CATALOG, &Interceptor{
		function: intc,
	})
	util.Logger().Infof("Intercept %s", intc.Name())
}
