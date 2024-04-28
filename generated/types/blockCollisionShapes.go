// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package types

import "encoding/json"
import "fmt"

type BlockCollisionShapes struct {
	// Each block's collision shape id(s).
	Blocks BlockCollisionShapesBlocks `json:"blocks" yaml:"blocks" mapstructure:"blocks"`

	// Collision shapes by id, each shape being composed of a list of collision boxes.
	Shapes BlockCollisionShapesShapes `json:"shapes" yaml:"shapes" mapstructure:"shapes"`
}

// Each block's collision shape id(s).
type BlockCollisionShapesBlocks map[string]interface{}

// Collision shapes by id, each shape being composed of a list of collision boxes.
type BlockCollisionShapesShapes map[string][][]float64

// UnmarshalJSON implements json.Unmarshaler.
func (j *BlockCollisionShapes) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["blocks"]; raw != nil && !ok {
		return fmt.Errorf("field blocks in BlockCollisionShapes: required")
	}
	if _, ok := raw["shapes"]; raw != nil && !ok {
		return fmt.Errorf("field shapes in BlockCollisionShapes: required")
	}
	type Plain BlockCollisionShapes
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BlockCollisionShapes(plain)
	return nil
}