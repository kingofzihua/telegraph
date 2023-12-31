// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build jsoniter
// +build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

// RawMessage is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
type RawMessage = jsoniter.RawMessage

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkgn/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gitlab.com/skonline/sk-user/sk-user-api-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
