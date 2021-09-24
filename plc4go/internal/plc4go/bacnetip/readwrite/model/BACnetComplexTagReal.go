/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetComplexTagReal struct {
	Value  float32
	Parent *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagReal interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagReal) DataType() BACnetDataType {
	return BACnetDataType_REAL
}

func (m *BACnetComplexTagReal) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagReal(value float32, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagReal{
		Value:  value,
		Parent: NewBACnetComplexTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagReal(structType interface{}) *BACnetComplexTagReal {
	castFunc := func(typ interface{}) *BACnetComplexTagReal {
		if casted, ok := typ.(BACnetComplexTagReal); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagReal); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagReal(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagReal(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagReal) GetTypeName() string {
	return "BACnetComplexTagReal"
}

func (m *BACnetComplexTagReal) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagReal) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *BACnetComplexTagReal) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagRealParse(readBuffer utils.ReadBuffer, lengthValueType uint8, extLength uint8) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagReal"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	value, _valueErr := readBuffer.ReadFloat32("value", true, 8, 23)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagReal{
		Value:  value,
		Parent: &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagReal) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagReal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := float32(m.Value)
		_valueErr := writeBuffer.WriteFloat32("value", 32, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagReal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
