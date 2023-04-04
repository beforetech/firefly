// Copyright © 2023 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postTokenPoolPublish = &ffapi.Route{
	Name:   "postTokenPoolPublish",
	Path:   "tokens/pools/{nameOrId}/publish",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsTokenPoolNameOrID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsPostTokenPool,
	JSONInputValue:  func() interface{} { return &core.EmptyInput{} },
	JSONOutputValue: func() interface{} { return &core.TokenPool{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			var pool *core.TokenPool
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			pool, err = cr.or.Assets().GetTokenPoolByNameOrID(cr.ctx, r.PP["nameOrId"])
			if err == nil {
				err = cr.or.DefinitionSender().PublishTokenPool(cr.ctx, &core.TokenPoolDefinition{
					Pool: pool,
				}, waitConfirm)
			}
			return pool, err
		},
	},
}
