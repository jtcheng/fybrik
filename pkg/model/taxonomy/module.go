// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package taxonomy

import (
	"encoding/json"

	"fybrik.io/fybrik/pkg/serde"
)

type Capability string

type PluginType string

type ActionName string

// +kubebuilder:validation:Enum=read;write;delete;copy
type DataFlow string

const (
	// ReadFlow indicates a data set is being read
	ReadFlow DataFlow = "read"

	// WriteFlow indicates a data set is being written
	WriteFlow DataFlow = "write"

	// DeleteFlow indicates a data set is being deleted
	DeleteFlow DataFlow = "delete"

	// CopyFlow indicates a data set is being copied
	CopyFlow DataFlow = "copy"
)

// +kubebuilder:pruning:PreserveUnknownFields
type Action struct {
	Name                 ActionName       `json:"name"`
	AdditionalProperties serde.Properties `json:"-"`
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		nameKey: o.Name,
	}

	for key, value := range o.AdditionalProperties.Items {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Action) UnmarshalJSON(bytes []byte) (err error) {
	items := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &items); err == nil {
		o.Name = ActionName(items[nameKey].(string))
		delete(items, nameKey)
		if len(items) == 0 {
			items = nil
		}
		o.AdditionalProperties = serde.Properties{Items: items}
	}
	return err
}
