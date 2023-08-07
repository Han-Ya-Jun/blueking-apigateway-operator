/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - API 网关(BlueKing - APIGateway) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *     http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package client

import (
	"fmt"

	"github.com/TencentBlueKing/blueking-apigateway-operator/pkg/constant"
)

// ReportEventReq add event request
type ReportEventReq struct {
	PublishID     string                 `json:"-"`
	BkGatewayName string                 `json:"bk_gateway_name"`
	BkStageName   string                 `json:"bk_stage_name"`
	Name          constant.EventName     `json:"name"`
	Status        constant.EventStatus   `json:"status"`
	Detail        map[string]interface{} `json:"detail"`
}

// Validate ReportEventReq
func (r ReportEventReq) Validate() error {
	// filter sync event publish_id empty)
	if r.PublishID == "" {
		return fmt.Errorf("not need report event: %+v", r)
	}
	return nil
}
