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
	"sync"

	"github.com/google/uuid"
)

// DAG ...
type DAG struct {
	ID uuid.UUID

	mu       sync.Mutex
	Vertices map[uuid.UUID]*Vertex
}

// NewDAG ...
func NewDAG() *DAG {
	d := &DAG{
		ID: uuid.New(),

		Vertices: make(map[uuid.UUID]*Vertex, 0),
	}

	return d
}

// AddVertex ...
func (d *DAG) AddVertex(v *Vertex) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Vertices[v.ID] = v

	return nil
}

// AddEdge ...
func (d *DAG) AddEdge(fromVertex *Vertex, toVertex *Vertex) error {
	var ok bool

	d.mu.Lock()
	fromExists := false
	toExists := false
	for _, vertex := range d.Vertices {
		if vertex == fromVertex {
			fromExists = true
		}
		if vertex == toVertex {
			toExists = true
		}
	}
	if fromExists == false {
		return fmt.Errorf("Vertex with the id %v not found", fromVertex.ID)
	}
	if toExists == false {
		return fmt.Errorf("Vertex with the id %v not found", toVertex.ID)
	}
	d.mu.Unlock()

	d.mu.Lock()
	if fromVertex, ok = d.Vertices[fromVertex.ID]; !ok {
		return fmt.Errorf("Vertex with the id %v not found", fromVertex.ID)
	}
	d.mu.Unlock()

	d.mu.Lock()
	if toVertex, ok = d.Vertices[toVertex.ID]; !ok {
		return fmt.Errorf("vertex with the id %v not found", toVertex.ID)
	}
	d.mu.Unlock()

	for _, childVertex := range fromVertex.Children {
		if childVertex == toVertex {
			return fmt.Errorf("edge (%v,%v) already exists", fromVertex.ID, toVertex.ID)
		}
	}
	fromVertex.Children = append(fromVertex.Children, toVertex)

	return nil
}
