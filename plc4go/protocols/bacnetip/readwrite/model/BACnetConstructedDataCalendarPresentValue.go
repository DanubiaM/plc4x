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

// BACnetConstructedDataCalendarPresentValue is the corresponding interface of BACnetConstructedDataCalendarPresentValue
type BACnetConstructedDataCalendarPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataCalendarPresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCalendarPresentValue.
// This is useful for switch cases.
type BACnetConstructedDataCalendarPresentValueExactly interface {
	BACnetConstructedDataCalendarPresentValue
	isBACnetConstructedDataCalendarPresentValue() bool
}

// _BACnetConstructedDataCalendarPresentValue is the data-structure of this message
type _BACnetConstructedDataCalendarPresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CALENDAR
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetPresentValue() BACnetApplicationTagBoolean {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCalendarPresentValue factory function for _BACnetConstructedDataCalendarPresentValue
func NewBACnetConstructedDataCalendarPresentValue(presentValue BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCalendarPresentValue {
	_result := &_BACnetConstructedDataCalendarPresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCalendarPresentValue(structType interface{}) BACnetConstructedDataCalendarPresentValue {
	if casted, ok := structType.(BACnetConstructedDataCalendarPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCalendarPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetTypeName() string {
	return "BACnetConstructedDataCalendarPresentValue"
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCalendarPresentValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCalendarPresentValue, error) {
	return BACnetConstructedDataCalendarPresentValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCalendarPresentValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCalendarPresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCalendarPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCalendarPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field of BACnetConstructedDataCalendarPresentValue")
	}
	presentValue := _presentValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Virtual field
	_actualValue := presentValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCalendarPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCalendarPresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCalendarPresentValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PresentValue: presentValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCalendarPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCalendarPresentValue")
		}

		// Simple Field (presentValue)
		if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for presentValue")
		}
		_presentValueErr := writeBuffer.WriteSerializable(ctx, m.GetPresentValue())
		if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for presentValue")
		}
		if _presentValueErr != nil {
			return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCalendarPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCalendarPresentValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCalendarPresentValue) isBACnetConstructedDataCalendarPresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataCalendarPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
