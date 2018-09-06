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

// Vertex ...
type Vertex struct {
	ID       string
	Parents  []*Vertex
	Children []*Vertex
}

// NewVertex ...
func NewVertex(id string, value interface{}) *Vertex {
	v := &Vertex{
		ID:       id,
		Parents:  make([]*Vertex, 0),
		Children: make([]*Vertex, 0),
	}

	return v
}

// ParentsIDs ...
func (v *Vertex) ParentsIDs() []string {
	var ids []string
	for _, vertex := range v.Parents {
		ids = append(ids, vertex.ID)
	}

	return ids
}

// ChidrenIDs ...
func (v *Vertex) ChidrenIDs() []string {
	var ids []string
	for _, vertex := range v.Children {
		ids = append(ids, vertex.ID)
	}

	return ids
}
