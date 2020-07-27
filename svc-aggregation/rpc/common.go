//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

package rpc

import (
	"encoding/json"

	"github.com/bharath-b-hpe/odimra/lib-utilities/common"
	aggregatorproto "github.com/bharath-b-hpe/odimra/lib-utilities/proto/aggregator"
	"github.com/bharath-b-hpe/odimra/lib-utilities/response"
	"github.com/bharath-b-hpe/odimra/lib-utilities/services"
	"github.com/bharath-b-hpe/odimra/svc-aggregation/agcommon"
	"github.com/bharath-b-hpe/odimra/svc-aggregation/agmessagebus"
	"github.com/bharath-b-hpe/odimra/svc-aggregation/agmodel"
	"github.com/bharath-b-hpe/odimra/svc-aggregation/system"
	"github.com/bharath-b-hpe/odimra/svc-plugin-rest-client/pmbhandle"
)

// Aggregator struct helps to register service
type Aggregator struct {
	connector *system.ExternalInterface
}

// GetAggregator intializes all the required connection functions for the aggregation execution
func GetAggregator() *Aggregator {
	return &Aggregator{
		connector: &system.ExternalInterface{
			ContactClient:           pmbhandle.ContactPlugin,
			Auth:                    services.IsAuthorized,
			GetSessionUserName:      services.GetSessionUserName,
			CreateTask:              services.CreateTask,
			CreateChildTask:         services.CreateChildTask,
			UpdateTask:              system.UpdateTaskData,
			CreateSubcription:       system.CreateDefaultEventSubscription,
			PublishEvent:            system.PublishEvent,
			GetPluginStatus:         agcommon.GetPluginStatus,
			SubscribeToEMB:          services.SubscribeToEMB,
			EncryptPassword:         common.EncryptWithPublicKey,
			DecryptPassword:         common.DecryptWithPrivateKey,
			DeleteComputeSystem:     agmodel.DeleteComputeSystem,
			DeleteSystem:            agmodel.DeleteSystem,
			DeleteEventSubscription: services.DeleteSubscription,
			EventNotification:       agmessagebus.Publish,
		},
	}
}

func generateResponse(rpcResp response.RPC, aggResp *aggregatorproto.AggregatorResponse) {
	bytes, _ := json.Marshal(rpcResp.Body)
	*aggResp = aggregatorproto.AggregatorResponse{
		StatusCode:    rpcResp.StatusCode,
		StatusMessage: rpcResp.StatusMessage,
		Header:        rpcResp.Header,
		Body:          bytes,
	}
}

func generateTaskRespone(taskID, taskURI string, rpcResp *response.RPC) {
	commonResponse := response.Response{
		OdataType:    "#Task.v1_4_2.Task",
		ID:           taskID,
		Name:         "Task " + taskID,
		OdataContext: "/redfish/v1/$metadata#Task.Task",
		OdataID:      taskURI,
	}
	commonResponse.MessageArgs = []string{taskID}
	commonResponse.CreateGenericResponse(rpcResp.StatusMessage)
	rpcResp.Body = commonResponse
}