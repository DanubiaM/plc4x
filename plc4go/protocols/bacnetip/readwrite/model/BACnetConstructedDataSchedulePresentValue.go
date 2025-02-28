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

// BACnetConstructedDataSchedulePresentValue is the corresponding interface of BACnetConstructedDataSchedulePresentValue
type BACnetConstructedDataSchedulePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetConstructedDataElement
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetConstructedDataElement
}

// BACnetConstructedDataSchedulePresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSchedulePresentValue.
// This is useful for switch cases.
type BACnetConstructedDataSchedulePresentValueExactly interface {
	BACnetConstructedDataSchedulePresentValue
	isBACnetConstructedDataSchedulePresentValue() bool
}

// _BACnetConstructedDataSchedulePresentValue is the data-structure of this message
type _BACnetConstructedDataSchedulePresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetConstructedDataElement
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSchedulePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_SCHEDULE
}

func (m *_BACnetConstructedDataSchedulePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSchedulePresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSchedulePresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSchedulePresentValue) GetPresentValue() BACnetConstructedDataElement {
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

func (m *_BACnetConstructedDataSchedulePresentValue) GetActualValue() BACnetConstructedDataElement {
	ctx := context.Background()
	_ = ctx
	return CastBACnetConstructedDataElement(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSchedulePresentValue factory function for _BACnetConstructedDataSchedulePresentValue
func NewBACnetConstructedDataSchedulePresentValue(presentValue BACnetConstructedDataElement, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSchedulePresentValue {
	_result := &_BACnetConstructedDataSchedulePresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSchedulePresentValue(structType interface{}) BACnetConstructedDataSchedulePresentValue {
	if casted, ok := structType.(BACnetConstructedDataSchedulePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSchedulePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSchedulePresentValue) GetTypeName() string {
	return "BACnetConstructedDataSchedulePresentValue"
}

func (m *_BACnetConstructedDataSchedulePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSchedulePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSchedulePresentValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSchedulePresentValue, error) {
	return BACnetConstructedDataSchedulePresentValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSchedulePresentValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSchedulePresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSchedulePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSchedulePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetConstructedDataElementParseWithBuffer(ctx, readBuffer, BACnetObjectType(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), nil)
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field of BACnetConstructedDataSchedulePresentValue")
	}
	presentValue := _presentValue.(BACnetConstructedDataElement)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Virtual field
	_actualValue := presentValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSchedulePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSchedulePresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSchedulePresentValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PresentValue: presentValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSchedulePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSchedulePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSchedulePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSchedulePresentValue")
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

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSchedulePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSchedulePresentValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSchedulePresentValue) isBACnetConstructedDataSchedulePresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataSchedulePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
