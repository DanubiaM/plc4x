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

// MResetInd is the corresponding interface of MResetInd
type MResetInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MResetIndExactly can be used when we want exactly this type and not a type which fulfills MResetInd.
// This is useful for switch cases.
type MResetIndExactly interface {
	MResetInd
	isMResetInd() bool
}

// _MResetInd is the data-structure of this message
type _MResetInd struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MResetInd) GetMessageCode() uint8 {
	return 0xF0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MResetInd) InitializeParent(parent CEMI) {}

func (m *_MResetInd) GetParent() CEMI {
	return m._CEMI
}

// NewMResetInd factory function for _MResetInd
func NewMResetInd(size uint16) *_MResetInd {
	_result := &_MResetInd{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMResetInd(structType interface{}) MResetInd {
	if casted, ok := structType.(MResetInd); ok {
		return casted
	}
	if casted, ok := structType.(*MResetInd); ok {
		return *casted
	}
	return nil
}

func (m *_MResetInd) GetTypeName() string {
	return "MResetInd"
}

func (m *_MResetInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MResetInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MResetIndParse(theBytes []byte, size uint16) (MResetInd, error) {
	return MResetIndParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), size)
}

func MResetIndParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (MResetInd, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MResetInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MResetInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MResetInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MResetInd")
	}

	// Create a partially initialized instance
	_child := &_MResetInd{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MResetInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MResetInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MResetInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MResetInd")
		}

		if popErr := writeBuffer.PopContext("MResetInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MResetInd")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MResetInd) isMResetInd() bool {
	return true
}

func (m *_MResetInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
