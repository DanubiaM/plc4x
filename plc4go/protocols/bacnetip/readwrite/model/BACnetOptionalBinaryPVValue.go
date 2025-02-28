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

// BACnetOptionalBinaryPVValue is the corresponding interface of BACnetOptionalBinaryPVValue
type BACnetOptionalBinaryPVValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalBinaryPV
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() BACnetBinaryPVTagged
}

// BACnetOptionalBinaryPVValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalBinaryPVValue.
// This is useful for switch cases.
type BACnetOptionalBinaryPVValueExactly interface {
	BACnetOptionalBinaryPVValue
	isBACnetOptionalBinaryPVValue() bool
}

// _BACnetOptionalBinaryPVValue is the data-structure of this message
type _BACnetOptionalBinaryPVValue struct {
	*_BACnetOptionalBinaryPV
	BinaryPv BACnetBinaryPVTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalBinaryPVValue) InitializeParent(parent BACnetOptionalBinaryPV, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalBinaryPVValue) GetParent() BACnetOptionalBinaryPV {
	return m._BACnetOptionalBinaryPV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalBinaryPVValue) GetBinaryPv() BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalBinaryPVValue factory function for _BACnetOptionalBinaryPVValue
func NewBACnetOptionalBinaryPVValue(binaryPv BACnetBinaryPVTagged, peekedTagHeader BACnetTagHeader) *_BACnetOptionalBinaryPVValue {
	_result := &_BACnetOptionalBinaryPVValue{
		BinaryPv:                binaryPv,
		_BACnetOptionalBinaryPV: NewBACnetOptionalBinaryPV(peekedTagHeader),
	}
	_result._BACnetOptionalBinaryPV._BACnetOptionalBinaryPVChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalBinaryPVValue(structType interface{}) BACnetOptionalBinaryPVValue {
	if casted, ok := structType.(BACnetOptionalBinaryPVValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPVValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalBinaryPVValue) GetTypeName() string {
	return "BACnetOptionalBinaryPVValue"
}

func (m *_BACnetOptionalBinaryPVValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalBinaryPVValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOptionalBinaryPVValueParse(theBytes []byte) (BACnetOptionalBinaryPVValue, error) {
	return BACnetOptionalBinaryPVValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalBinaryPVValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalBinaryPVValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPVValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPVValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryPv)
	if pullErr := readBuffer.PullContext("binaryPv"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryPv")
	}
	_binaryPv, _binaryPvErr := BACnetBinaryPVTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _binaryPvErr != nil {
		return nil, errors.Wrap(_binaryPvErr, "Error parsing 'binaryPv' field of BACnetOptionalBinaryPVValue")
	}
	binaryPv := _binaryPv.(BACnetBinaryPVTagged)
	if closeErr := readBuffer.CloseContext("binaryPv"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryPv")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPVValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPVValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalBinaryPVValue{
		_BACnetOptionalBinaryPV: &_BACnetOptionalBinaryPV{},
		BinaryPv:                binaryPv,
	}
	_child._BACnetOptionalBinaryPV._BACnetOptionalBinaryPVChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalBinaryPVValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalBinaryPVValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPVValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPVValue")
		}

		// Simple Field (binaryPv)
		if pushErr := writeBuffer.PushContext("binaryPv"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for binaryPv")
		}
		_binaryPvErr := writeBuffer.WriteSerializable(ctx, m.GetBinaryPv())
		if popErr := writeBuffer.PopContext("binaryPv"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for binaryPv")
		}
		if _binaryPvErr != nil {
			return errors.Wrap(_binaryPvErr, "Error serializing 'binaryPv' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPVValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPVValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalBinaryPVValue) isBACnetOptionalBinaryPVValue() bool {
	return true
}

func (m *_BACnetOptionalBinaryPVValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
