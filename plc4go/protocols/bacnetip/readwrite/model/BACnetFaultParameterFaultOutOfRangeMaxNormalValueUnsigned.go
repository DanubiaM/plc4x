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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMaxNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(unsignedValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned{
		UnsignedValue: unsignedValue,
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned(structType interface{}) BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedParse(theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, error) {
	return BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned{
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &_BACnetFaultParameterFaultOutOfRangeMaxNormalValue{
			TagNumber: tagNumber,
		},
		UnsignedValue: unsignedValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(ctx, m.GetUnsignedValue())
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) isBACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
