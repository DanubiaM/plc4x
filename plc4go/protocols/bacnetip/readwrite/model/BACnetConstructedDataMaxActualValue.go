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

// BACnetConstructedDataMaxActualValue is the corresponding interface of BACnetConstructedDataMaxActualValue
type BACnetConstructedDataMaxActualValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxActualValue returns MaxActualValue (property field)
	GetMaxActualValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataMaxActualValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaxActualValue.
// This is useful for switch cases.
type BACnetConstructedDataMaxActualValueExactly interface {
	BACnetConstructedDataMaxActualValue
	isBACnetConstructedDataMaxActualValue() bool
}

// _BACnetConstructedDataMaxActualValue is the data-structure of this message
type _BACnetConstructedDataMaxActualValue struct {
	*_BACnetConstructedData
	MaxActualValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxActualValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_ACTUAL_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxActualValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaxActualValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetMaxActualValue() BACnetApplicationTagReal {
	return m.MaxActualValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxActualValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxActualValue factory function for _BACnetConstructedDataMaxActualValue
func NewBACnetConstructedDataMaxActualValue(maxActualValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxActualValue {
	_result := &_BACnetConstructedDataMaxActualValue{
		MaxActualValue:         maxActualValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxActualValue(structType interface{}) BACnetConstructedDataMaxActualValue {
	if casted, ok := structType.(BACnetConstructedDataMaxActualValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxActualValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxActualValue) GetTypeName() string {
	return "BACnetConstructedDataMaxActualValue"
}

func (m *_BACnetConstructedDataMaxActualValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (maxActualValue)
	lengthInBits += m.MaxActualValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxActualValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataMaxActualValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaxActualValue, error) {
	return BACnetConstructedDataMaxActualValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMaxActualValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaxActualValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxActualValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxActualValue)
	if pullErr := readBuffer.PullContext("maxActualValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxActualValue")
	}
	_maxActualValue, _maxActualValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _maxActualValueErr != nil {
		return nil, errors.Wrap(_maxActualValueErr, "Error parsing 'maxActualValue' field of BACnetConstructedDataMaxActualValue")
	}
	maxActualValue := _maxActualValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("maxActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxActualValue")
	}

	// Virtual field
	_actualValue := maxActualValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxActualValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxActualValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaxActualValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxActualValue: maxActualValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaxActualValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxActualValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxActualValue")
		}

		// Simple Field (maxActualValue)
		if pushErr := writeBuffer.PushContext("maxActualValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxActualValue")
		}
		_maxActualValueErr := writeBuffer.WriteSerializable(ctx, m.GetMaxActualValue())
		if popErr := writeBuffer.PopContext("maxActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxActualValue")
		}
		if _maxActualValueErr != nil {
			return errors.Wrap(_maxActualValueErr, "Error serializing 'maxActualValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxActualValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxActualValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxActualValue) isBACnetConstructedDataMaxActualValue() bool {
	return true
}

func (m *_BACnetConstructedDataMaxActualValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
