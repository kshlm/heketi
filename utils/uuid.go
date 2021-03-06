//
// Copyright (c) 2014 The heketi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

// From http://www.ashishbanerjee.com/home/go/go-generate-uuid

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/lpabon/godbc"
)

func GenUUID() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	godbc.Check(n == len(uuid), n, len(uuid))
	godbc.Check(err == nil, err)

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}
