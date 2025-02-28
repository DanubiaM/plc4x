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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogDataEntryAnyValue is the corresponding interface of BACnetLogDataLogDataEntryAnyValue
type BACnetLogDataLogDataEntryAnyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetAnyValue returns AnyValue (property field)
	GetAnyValue() BACnetConstructedData
}

// BACnetLogDataLogDataEntryAnyValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryAnyValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryAnyValueExactly interface {
	BACnetLogDataLogDataEntryAnyValue
	isBACnetLogDataLogDataEntryAnyValue() bool
}

// _BACnetLogDataLogDataEntryAnyValue is the data-structure of this message
type _BACnetLogDataLogDataEntryAnyValue struct {
	*_BACnetLogDataLogDataEntry
	AnyValue BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryAnyValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryAnyValue) GetAnyValue() BACnetConstructedData {
	return m.AnyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryAnyValue factory function for _BACnetLogDataLogDataEntryAnyValue
func NewBACnetLogDataLogDataEntryAnyValue(anyValue BACnetConstructedData, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryAnyValue {
	_result := &_BACnetLogDataLogDataEntryAnyValue{
		AnyValue:                   anyValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryAnyValue(structType interface{}) BACnetLogDataLogDataEntryAnyValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryAnyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryAnyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryAnyValue"
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (anyValue)
	if m.AnyValue != nil {
		lengthInBits += m.AnyValue.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryAnyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogDataEntryAnyValueParse(theBytes []byte) (BACnetLogDataLogDataEntryAnyValue, error) {
	return BACnetLogDataLogDataEntryAnyValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogDataLogDataEntryAnyValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryAnyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryAnyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryAnyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (anyValue) (Can be skipped, if a given expression evaluates to false)
	var anyValue BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("anyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for anyValue")
		}
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(8), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'anyValue' field of BACnetLogDataLogDataEntryAnyValue")
		default:
			anyValue = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("anyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for anyValue")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryAnyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryAnyValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryAnyValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		AnyValue:                   anyValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryAnyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryAnyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryAnyValue")
		}

		// Optional Field (anyValue) (Can be skipped, if the value is null)
		var anyValue BACnetConstructedData = nil
		if m.GetAnyValue() != nil {
			if pushErr := writeBuffer.PushContext("anyValue"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for anyValue")
			}
			anyValue = m.GetAnyValue()
			_anyValueErr := writeBuffer.WriteSerializable(ctx, anyValue)
			if popErr := writeBuffer.PopContext("anyValue"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for anyValue")
			}
			if _anyValueErr != nil {
				return errors.Wrap(_anyValueErr, "Error serializing 'anyValue' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryAnyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryAnyValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryAnyValue) isBACnetLogDataLogDataEntryAnyValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryAnyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
