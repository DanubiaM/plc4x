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

// BACnetOptionalCharacterStringNull is the corresponding interface of BACnetOptionalCharacterStringNull
type BACnetOptionalCharacterStringNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalCharacterString
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetOptionalCharacterStringNullExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalCharacterStringNull.
// This is useful for switch cases.
type BACnetOptionalCharacterStringNullExactly interface {
	BACnetOptionalCharacterStringNull
	isBACnetOptionalCharacterStringNull() bool
}

// _BACnetOptionalCharacterStringNull is the data-structure of this message
type _BACnetOptionalCharacterStringNull struct {
	*_BACnetOptionalCharacterString
	NullValue BACnetApplicationTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalCharacterStringNull) InitializeParent(parent BACnetOptionalCharacterString, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalCharacterStringNull) GetParent() BACnetOptionalCharacterString {
	return m._BACnetOptionalCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalCharacterStringNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalCharacterStringNull factory function for _BACnetOptionalCharacterStringNull
func NewBACnetOptionalCharacterStringNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetOptionalCharacterStringNull {
	_result := &_BACnetOptionalCharacterStringNull{
		NullValue:                      nullValue,
		_BACnetOptionalCharacterString: NewBACnetOptionalCharacterString(peekedTagHeader),
	}
	_result._BACnetOptionalCharacterString._BACnetOptionalCharacterStringChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalCharacterStringNull(structType interface{}) BACnetOptionalCharacterStringNull {
	if casted, ok := structType.(BACnetOptionalCharacterStringNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterStringNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalCharacterStringNull) GetTypeName() string {
	return "BACnetOptionalCharacterStringNull"
}

func (m *_BACnetOptionalCharacterStringNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalCharacterStringNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOptionalCharacterStringNullParse(theBytes []byte) (BACnetOptionalCharacterStringNull, error) {
	return BACnetOptionalCharacterStringNullParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalCharacterStringNullParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalCharacterStringNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterStringNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterStringNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
	_nullValue, _nullValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field of BACnetOptionalCharacterStringNull")
	}
	nullValue := _nullValue.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterStringNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterStringNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalCharacterStringNull{
		_BACnetOptionalCharacterString: &_BACnetOptionalCharacterString{},
		NullValue:                      nullValue,
	}
	_child._BACnetOptionalCharacterString._BACnetOptionalCharacterStringChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalCharacterStringNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalCharacterStringNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterStringNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterStringNull")
		}

		// Simple Field (nullValue)
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		_nullValueErr := writeBuffer.WriteSerializable(ctx, m.GetNullValue())
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalCharacterStringNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterStringNull")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalCharacterStringNull) isBACnetOptionalCharacterStringNull() bool {
	return true
}

func (m *_BACnetOptionalCharacterStringNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
