// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

// RawMessage is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
type RawMessage = json.RawMessage

var (
	// Marshal is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
