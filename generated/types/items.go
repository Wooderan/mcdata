// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package types

import "encoding/json"
import "fmt"

type Items []struct {
	// The display name of an item
	DisplayName string `json:"displayName" yaml:"displayName" mapstructure:"displayName"`

	// describes categories of enchants this item can use
	EnchantCategories []string `json:"enchantCategories,omitempty" yaml:"enchantCategories,omitempty" mapstructure:"enchantCategories,omitempty"`

	// The unique identifier for an item
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// the amount of durability an item has before being damaged/used
	MaxDurability *int `json:"maxDurability,omitempty" yaml:"maxDurability,omitempty" mapstructure:"maxDurability,omitempty"`

	// The name of an item
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// describes what items this item can be fixed with in an anvil
	RepairWith []string `json:"repairWith,omitempty" yaml:"repairWith,omitempty" mapstructure:"repairWith,omitempty"`

	// Stack size for an item
	StackSize int `json:"stackSize" yaml:"stackSize" mapstructure:"stackSize"`

	// Variations corresponds to the JSON schema field "variations".
	Variations []ItemsElemVariationsElem `json:"variations,omitempty" yaml:"variations,omitempty" mapstructure:"variations,omitempty"`

	AdditionalProperties interface{}
}

type ItemsElemVariationsElem struct {
	// DisplayName corresponds to the JSON schema field "displayName".
	DisplayName string `json:"displayName" yaml:"displayName" mapstructure:"displayName"`

	// Metadata corresponds to the JSON schema field "metadata".
	Metadata int `json:"metadata" yaml:"metadata" mapstructure:"metadata"`

	AdditionalProperties interface{}
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ItemsElemVariationsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["displayName"]; raw != nil && !ok {
		return fmt.Errorf("field displayName in ItemsElemVariationsElem: required")
	}
	if _, ok := raw["metadata"]; raw != nil && !ok {
		return fmt.Errorf("field metadata in ItemsElemVariationsElem: required")
	}
	type Plain ItemsElemVariationsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ItemsElemVariationsElem(plain)
	return nil
}