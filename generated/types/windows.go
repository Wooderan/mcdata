// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package types

import "encoding/json"
import "fmt"
import "reflect"

type Windows []struct {
	// The unique identifier for the window
	Id string `json:"id" yaml:"id" mapstructure:"id"`

	// The default displayed name of the window
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// OpenedWith corresponds to the JSON schema field "openedWith".
	OpenedWith []WindowsElemOpenedWithElem `json:"openedWith,omitempty" yaml:"openedWith,omitempty" mapstructure:"openedWith,omitempty"`

	// Names of the properties of the window
	Properties []string `json:"properties,omitempty" yaml:"properties,omitempty" mapstructure:"properties,omitempty"`

	// The slots displayed in the window
	Slots []WindowsElemSlotsElem `json:"slots,omitempty" yaml:"slots,omitempty" mapstructure:"slots,omitempty"`
}

type WindowsElemOpenedWithElem struct {
	// Id corresponds to the JSON schema field "id".
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// Type corresponds to the JSON schema field "type".
	Type WindowsElemOpenedWithElemType `json:"type" yaml:"type" mapstructure:"type"`
}

type WindowsElemOpenedWithElemType string

const WindowsElemOpenedWithElemTypeBlock WindowsElemOpenedWithElemType = "block"
const WindowsElemOpenedWithElemTypeEntity WindowsElemOpenedWithElemType = "entity"
const WindowsElemOpenedWithElemTypeItem WindowsElemOpenedWithElemType = "item"

var enumValues_WindowsElemOpenedWithElemType = []interface{}{
	"item",
	"entity",
	"block",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *WindowsElemOpenedWithElemType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_WindowsElemOpenedWithElemType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_WindowsElemOpenedWithElemType, v)
	}
	*j = WindowsElemOpenedWithElemType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *WindowsElemOpenedWithElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in WindowsElemOpenedWithElem: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in WindowsElemOpenedWithElem: required")
	}
	type Plain WindowsElemOpenedWithElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = WindowsElemOpenedWithElem(plain)
	return nil
}

// A slot or slot range in the window
type WindowsElemSlotsElem struct {
	// The position of the slot or begin of the slot range
	Index int `json:"index" yaml:"index" mapstructure:"index"`

	// The name of the slot or slot range
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// The size of the slot range
	Size *int `json:"size,omitempty" yaml:"size,omitempty" mapstructure:"size,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *WindowsElemSlotsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["index"]; raw != nil && !ok {
		return fmt.Errorf("field index in WindowsElemSlotsElem: required")
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in WindowsElemSlotsElem: required")
	}
	type Plain WindowsElemSlotsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = WindowsElemSlotsElem(plain)
	return nil
}
