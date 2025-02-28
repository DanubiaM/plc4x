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

// BACnetChannelValueReal is the corresponding interface of BACnetChannelValueReal
type BACnetChannelValueReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// BACnetChannelValueRealExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueReal.
// This is useful for switch cases.
type BACnetChannelValueRealExactly interface {
	BACnetChannelValueReal
	isBACnetChannelValueReal() bool
}

// _BACnetChannelValueReal is the data-structure of this message
type _BACnetChannelValueReal struct {
	*_BACnetChannelValue
	RealValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueReal) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueReal) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueReal factory function for _BACnetChannelValueReal
func NewBACnetChannelValueReal(realValue BACnetApplicationTagReal, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueReal {
	_result := &_BACnetChannelValueReal{
		RealValue:           realValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueReal(structType interface{}) BACnetChannelValueReal {
	if casted, ok := structType.(BACnetChannelValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueReal) GetTypeName() string {
	return "BACnetChannelValueReal"
}

func (m *_BACnetChannelValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueRealParse(theBytes []byte) (BACnetChannelValueReal, error) {
	return BACnetChannelValueRealParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueRealParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
	_realValue, _realValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field of BACnetChannelValueReal")
	}
	realValue := _realValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueReal{
		_BACnetChannelValue: &_BACnetChannelValue{},
		RealValue:           realValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueReal")
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realValue")
		}
		_realValueErr := writeBuffer.WriteSerializable(ctx, m.GetRealValue())
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realValue")
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueReal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueReal) isBACnetChannelValueReal() bool {
	return true
}

func (m *_BACnetChannelValueReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
