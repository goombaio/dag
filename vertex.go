// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package dag

import (
	"fmt"

	"github.com/goombaio/orderedset"
)

// Vertex ...
type Vertex struct {
	ID       string
	Parents  *orderedset.OrderedSet
	Children *orderedset.OrderedSet
}

// NewVertex ...
func NewVertex(id string, value interface{}) *Vertex {
	v := &Vertex{
		ID:       id,
		Parents:  orderedset.NewOrderedSet(),
		Children: orderedset.NewOrderedSet(),
	}

	return v
}

// String implements stringer interface and prints an string representation
// of this instance.
func (v *Vertex) String() string {
	result := fmt.Sprintf("Vertex ID: %s - Parents: %d - Children: %d\n", v.ID, v.Parents.Size(), v.Children.Size())

	return result
}
