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

// BACnetConstructedDataSetpoint is the corresponding interface of BACnetConstructedDataSetpoint
type BACnetConstructedDataSetpoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSetpoint returns Setpoint (property field)
	GetSetpoint() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataSetpointExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSetpoint.
// This is useful for switch cases.
type BACnetConstructedDataSetpointExactly interface {
	BACnetConstructedDataSetpoint
	isBACnetConstructedDataSetpoint() bool
}

// _BACnetConstructedDataSetpoint is the data-structure of this message
type _BACnetConstructedDataSetpoint struct {
	*_BACnetConstructedData
	Setpoint BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSetpoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSetpoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SETPOINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSetpoint) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSetpoint) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSetpoint) GetSetpoint() BACnetApplicationTagReal {
	return m.Setpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSetpoint) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetSetpoint())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSetpoint factory function for _BACnetConstructedDataSetpoint
func NewBACnetConstructedDataSetpoint(setpoint BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSetpoint {
	_result := &_BACnetConstructedDataSetpoint{
		Setpoint:               setpoint,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSetpoint(structType interface{}) BACnetConstructedDataSetpoint {
	if casted, ok := structType.(BACnetConstructedDataSetpoint); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSetpoint); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSetpoint) GetTypeName() string {
	return "BACnetConstructedDataSetpoint"
}

func (m *_BACnetConstructedDataSetpoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (setpoint)
	lengthInBits += m.Setpoint.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSetpoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSetpointParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSetpoint, error) {
	return BACnetConstructedDataSetpointParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSetpointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSetpoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSetpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSetpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (setpoint)
	if pullErr := readBuffer.PullContext("setpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setpoint")
	}
	_setpoint, _setpointErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _setpointErr != nil {
		return nil, errors.Wrap(_setpointErr, "Error parsing 'setpoint' field of BACnetConstructedDataSetpoint")
	}
	setpoint := _setpoint.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("setpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setpoint")
	}

	// Virtual field
	_actualValue := setpoint
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSetpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSetpoint")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSetpoint{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Setpoint: setpoint,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSetpoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSetpoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSetpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSetpoint")
		}

		// Simple Field (setpoint)
		if pushErr := writeBuffer.PushContext("setpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for setpoint")
		}
		_setpointErr := writeBuffer.WriteSerializable(ctx, m.GetSetpoint())
		if popErr := writeBuffer.PopContext("setpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for setpoint")
		}
		if _setpointErr != nil {
			return errors.Wrap(_setpointErr, "Error serializing 'setpoint' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSetpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSetpoint")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSetpoint) isBACnetConstructedDataSetpoint() bool {
	return true
}

func (m *_BACnetConstructedDataSetpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
