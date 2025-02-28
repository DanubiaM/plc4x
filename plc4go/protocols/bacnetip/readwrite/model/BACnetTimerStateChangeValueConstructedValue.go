/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetTimerStateChangeValueConstructedValue is the corresponding interface of BACnetTimerStateChangeValueConstructedValue
type BACnetTimerStateChangeValueConstructedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() BACnetConstructedData
}

// BACnetTimerStateChangeValueConstructedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueConstructedValue.
// This is useful for switch cases.
type BACnetTimerStateChangeValueConstructedValueExactly interface {
	BACnetTimerStateChangeValueConstructedValue
	isBACnetTimerStateChangeValueConstructedValue() bool
}

// _BACnetTimerStateChangeValueConstructedValue is the data-structure of this message
type _BACnetTimerStateChangeValueConstructedValue struct {
	*_BACnetTimerStateChangeValue
	ConstructedValue BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueConstructedValue) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueConstructedValue) GetConstructedValue() BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueConstructedValue factory function for _BACnetTimerStateChangeValueConstructedValue
func NewBACnetTimerStateChangeValueConstructedValue(constructedValue BACnetConstructedData, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueConstructedValue {
	_result := &_BACnetTimerStateChangeValueConstructedValue{
		ConstructedValue:             constructedValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueConstructedValue(structType interface{}) BACnetTimerStateChangeValueConstructedValue {
	if casted, ok := structType.(BACnetTimerStateChangeValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueConstructedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetTypeName() string {
	return "BACnetTimerStateChangeValueConstructedValue"
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueConstructedValueParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueConstructedValue, error) {
	return BACnetTimerStateChangeValueConstructedValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueConstructedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueConstructedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (constructedValue)
	if pullErr := readBuffer.PullContext("constructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for constructedValue")
	}
	_constructedValue, _constructedValueErr := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), nil)
	if _constructedValueErr != nil {
		return nil, errors.Wrap(_constructedValueErr, "Error parsing 'constructedValue' field of BACnetTimerStateChangeValueConstructedValue")
	}
	constructedValue := _constructedValue.(BACnetConstructedData)
	if closeErr := readBuffer.CloseContext("constructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for constructedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueConstructedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueConstructedValue{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		ConstructedValue: constructedValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueConstructedValue")
		}

		// Simple Field (constructedValue)
		if pushErr := writeBuffer.PushContext("constructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for constructedValue")
		}
		_constructedValueErr := writeBuffer.WriteSerializable(ctx, m.GetConstructedValue())
		if popErr := writeBuffer.PopContext("constructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for constructedValue")
		}
		if _constructedValueErr != nil {
			return errors.Wrap(_constructedValueErr, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueConstructedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueConstructedValue) isBACnetTimerStateChangeValueConstructedValue() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueConstructedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
