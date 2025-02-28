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

// ApduDataMemoryWrite is the corresponding interface of ApduDataMemoryWrite
type ApduDataMemoryWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataMemoryWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataMemoryWrite.
// This is useful for switch cases.
type ApduDataMemoryWriteExactly interface {
	ApduDataMemoryWrite
	isApduDataMemoryWrite() bool
}

// _ApduDataMemoryWrite is the data-structure of this message
type _ApduDataMemoryWrite struct {
	*_ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryWrite) GetApciType() uint8 {
	return 0xA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryWrite) InitializeParent(parent ApduData) {}

func (m *_ApduDataMemoryWrite) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataMemoryWrite factory function for _ApduDataMemoryWrite
func NewApduDataMemoryWrite(dataLength uint8) *_ApduDataMemoryWrite {
	_result := &_ApduDataMemoryWrite{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataMemoryWrite(structType interface{}) ApduDataMemoryWrite {
	if casted, ok := structType.(ApduDataMemoryWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryWrite) GetTypeName() string {
	return "ApduDataMemoryWrite"
}

func (m *_ApduDataMemoryWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataMemoryWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataMemoryWriteParse(theBytes []byte, dataLength uint8) (ApduDataMemoryWrite, error) {
	return ApduDataMemoryWriteParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataMemoryWriteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataMemoryWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataMemoryWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataMemoryWrite{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataMemoryWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataMemoryWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryWrite")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataMemoryWrite) isApduDataMemoryWrite() bool {
	return true
}

func (m *_ApduDataMemoryWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
