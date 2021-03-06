/*
Copyright 2014 go-trello authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trello

import "net/http"

type Client struct {
	// Auth

	client   *http.Client
	endpoint string
	version  string
	authenticated string
}
// You can set "" any values to access to public information.
func NewClient(token string, consumerKey string) (*Client, error) {
	version := "1"
	endpoint := "https://api.trello.com/" + version
	filters := "createCard,createBoard,commentCard,updateCheckItemStateOnCard,addChecklistToCard,createList,updateCard"
	client := Client{ client: http.DefaultClient, endpoint: endpoint, version: version}

	if token != "" && consumerKey != "" {
		client.authenticated = "?key="+consumerKey+"&token="+token+"&filter="+filters
	}

	return &client, nil
}
