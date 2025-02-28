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

// BACnetLogRecordLogDatumAnyValue is the corresponding interface of BACnetLogRecordLogDatumAnyValue
type BACnetLogRecordLogDatumAnyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetAnyValue returns AnyValue (property field)
	GetAnyValue() BACnetConstructedData
}

// BACnetLogRecordLogDatumAnyValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumAnyValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumAnyValueExactly interface {
	BACnetLogRecordLogDatumAnyValue
	isBACnetLogRecordLogDatumAnyValue() bool
}

// _BACnetLogRecordLogDatumAnyValue is the data-structure of this message
type _BACnetLogRecordLogDatumAnyValue struct {
	*_BACnetLogRecordLogDatum
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

func (m *_BACnetLogRecordLogDatumAnyValue) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumAnyValue) GetAnyValue() BACnetConstructedData {
	return m.AnyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumAnyValue factory function for _BACnetLogRecordLogDatumAnyValue
func NewBACnetLogRecordLogDatumAnyValue(anyValue BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumAnyValue {
	_result := &_BACnetLogRecordLogDatumAnyValue{
		AnyValue:                 anyValue,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumAnyValue(structType interface{}) BACnetLogRecordLogDatumAnyValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumAnyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumAnyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumAnyValue"
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (anyValue)
	if m.AnyValue != nil {
		lengthInBits += m.AnyValue.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumAnyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordLogDatumAnyValueParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumAnyValue, error) {
	return BACnetLogRecordLogDatumAnyValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumAnyValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumAnyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumAnyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumAnyValue")
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
		_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(10), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'anyValue' field of BACnetLogRecordLogDatumAnyValue")
		default:
			anyValue = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("anyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for anyValue")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumAnyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumAnyValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumAnyValue{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		AnyValue: anyValue,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumAnyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumAnyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumAnyValue")
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

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumAnyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumAnyValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumAnyValue) isBACnetLogRecordLogDatumAnyValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumAnyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
