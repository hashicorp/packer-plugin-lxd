// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incus

import (
	"fmt"
)

type Artifact struct {
	id string

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return nil
}

func (a *Artifact) Id() string {
	return a.id
}

func (a *Artifact) String() string {
	return fmt.Sprintf("image: %s", a.id)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	_, err := IncusCommand("image", "delete", a.id)
	return err
}
