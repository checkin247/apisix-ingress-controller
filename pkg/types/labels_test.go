// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelsIsSubsetOf(t *testing.T) {
	l := Labels{}
	f := Labels{
		"version": "v1",
		"env":     "prod",
	}
	assert.Equal(t, true, l.IsSubsetOf(f))
	l["env"] = "prod"
	assert.Equal(t, true, l.IsSubsetOf(f))
	l["env"] = "qa"
	assert.Equal(t, false, l.IsSubsetOf(f))
}
