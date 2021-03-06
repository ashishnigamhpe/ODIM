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

// Package services ...
package services

import (
	"context"
	"log"
	"net/http"

	authproto "github.com/ODIM-Project/ODIM/lib-utilities/proto/auth"
	sessionproto "github.com/ODIM-Project/ODIM/lib-utilities/proto/session"
	errResponse "github.com/ODIM-Project/ODIM/lib-utilities/response"
)

// IsAuthorized is used to authorize the services using svc-account-session.
// As parameters session token, privileges and oem privileges are passed.
// A RPC call is made with these parameters to the Account-Session service
// to check whether the session is valid and have all the privileges which are
// passed to it. After the RPC response, the function will return status code and
// status message back to the caller.
func IsAuthorized(sessionToken string, privileges, oemPrivileges []string) (int32, string) {
	asService := authproto.NewAuthorizationService(AccountSession, Service.Client())
	response, err := asService.IsAuthorized(
		context.TODO(),
		&authproto.AuthRequest{
			SessionToken:  sessionToken,
			Privileges:    privileges,
			Oemprivileges: oemPrivileges,
		},
	)
	if err != nil && response == nil {
		log.Printf("error: something went wrong with rpc call: %v", err)
		return http.StatusInternalServerError, errResponse.InternalError
	}
	return response.StatusCode, response.StatusMessage
}

// GetSessionUserName will get user name from the session token by rpc call to account-session service
func GetSessionUserName(sessionToken string) (string, error) {
	asService := sessionproto.NewSessionService(AccountSession, Service.Client())
	response, err := asService.GetSessionUserName(
		context.TODO(),
		&sessionproto.SessionRequest{
			SessionToken: sessionToken,
		},
	)
	if err != nil && response == nil {
		log.Printf("error: something went wrong with rpc call: %v", err)
		return "", err
	}
	return response.UserName, err
}
