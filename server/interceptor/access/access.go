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
package access

import (
	"fmt"
	"github.com/ServiceComb/service-center/pkg/util"
	"github.com/ServiceComb/service-center/pkg/validate"
	"github.com/ServiceComb/service-center/server/core"
	"net/http"
)

var serverName string = core.Service.ServiceName + "/" + core.Service.Version

func Intercept(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("server", serverName)

	util.InitContext(r)

	if !validate.IsRequestURI(r.RequestURI) {
		err := fmt.Errorf("Invalid Request URI %s", r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(util.StringToBytesWithNoCopy(err.Error()))
		return err
	}
	return nil
}

func Log(w http.ResponseWriter, r *http.Request) error {
	util.LogNilOrWarnf(util.GetStartTimeFromContext(r.Context()), "%s %s", r.Method, r.RequestURI)
	return nil
}
