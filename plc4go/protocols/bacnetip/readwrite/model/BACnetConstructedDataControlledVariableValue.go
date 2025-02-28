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

// BACnetConstructedDataControlledVariableValue is the corresponding interface of BACnetConstructedDataControlledVariableValue
type BACnetConstructedDataControlledVariableValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetControlledVariableValue returns ControlledVariableValue (property field)
	GetControlledVariableValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataControlledVariableValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataControlledVariableValue.
// This is useful for switch cases.
type BACnetConstructedDataControlledVariableValueExactly interface {
	BACnetConstructedDataControlledVariableValue
	isBACnetConstructedDataControlledVariableValue() bool
}

// _BACnetConstructedDataControlledVariableValue is the data-structure of this message
type _BACnetConstructedDataControlledVariableValue struct {
	*_BACnetConstructedData
	ControlledVariableValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataControlledVariableValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetControlledVariableValue() BACnetApplicationTagReal {
	return m.ControlledVariableValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetControlledVariableValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableValue factory function for _BACnetConstructedDataControlledVariableValue
func NewBACnetConstructedDataControlledVariableValue(controlledVariableValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableValue {
	_result := &_BACnetConstructedDataControlledVariableValue{
		ControlledVariableValue: controlledVariableValue,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableValue(structType interface{}) BACnetConstructedDataControlledVariableValue {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableValue) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableValue"
}

func (m *_BACnetConstructedDataControlledVariableValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (controlledVariableValue)
	lengthInBits += m.ControlledVariableValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataControlledVariableValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableValue, error) {
	return BACnetConstructedDataControlledVariableValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataControlledVariableValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlledVariableValue)
	if pullErr := readBuffer.PullContext("controlledVariableValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlledVariableValue")
	}
	_controlledVariableValue, _controlledVariableValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _controlledVariableValueErr != nil {
		return nil, errors.Wrap(_controlledVariableValueErr, "Error parsing 'controlledVariableValue' field of BACnetConstructedDataControlledVariableValue")
	}
	controlledVariableValue := _controlledVariableValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("controlledVariableValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlledVariableValue")
	}

	// Virtual field
	_actualValue := controlledVariableValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataControlledVariableValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ControlledVariableValue: controlledVariableValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataControlledVariableValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableValue")
		}

		// Simple Field (controlledVariableValue)
		if pushErr := writeBuffer.PushContext("controlledVariableValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for controlledVariableValue")
		}
		_controlledVariableValueErr := writeBuffer.WriteSerializable(ctx, m.GetControlledVariableValue())
		if popErr := writeBuffer.PopContext("controlledVariableValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for controlledVariableValue")
		}
		if _controlledVariableValueErr != nil {
			return errors.Wrap(_controlledVariableValueErr, "Error serializing 'controlledVariableValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableValue) isBACnetConstructedDataControlledVariableValue() bool {
	return true
}

func (m *_BACnetConstructedDataControlledVariableValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
