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

// ConnectionResponseDataBlockDeviceManagement is the corresponding interface of ConnectionResponseDataBlockDeviceManagement
type ConnectionResponseDataBlockDeviceManagement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ConnectionResponseDataBlock
}

// ConnectionResponseDataBlockDeviceManagementExactly can be used when we want exactly this type and not a type which fulfills ConnectionResponseDataBlockDeviceManagement.
// This is useful for switch cases.
type ConnectionResponseDataBlockDeviceManagementExactly interface {
	ConnectionResponseDataBlockDeviceManagement
	isConnectionResponseDataBlockDeviceManagement() bool
}

// _ConnectionResponseDataBlockDeviceManagement is the data-structure of this message
type _ConnectionResponseDataBlockDeviceManagement struct {
	*_ConnectionResponseDataBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionResponseDataBlockDeviceManagement) GetConnectionType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionResponseDataBlockDeviceManagement) InitializeParent(parent ConnectionResponseDataBlock) {
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetParent() ConnectionResponseDataBlock {
	return m._ConnectionResponseDataBlock
}

// NewConnectionResponseDataBlockDeviceManagement factory function for _ConnectionResponseDataBlockDeviceManagement
func NewConnectionResponseDataBlockDeviceManagement() *_ConnectionResponseDataBlockDeviceManagement {
	_result := &_ConnectionResponseDataBlockDeviceManagement{
		_ConnectionResponseDataBlock: NewConnectionResponseDataBlock(),
	}
	_result._ConnectionResponseDataBlock._ConnectionResponseDataBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionResponseDataBlockDeviceManagement(structType interface{}) ConnectionResponseDataBlockDeviceManagement {
	if casted, ok := structType.(ConnectionResponseDataBlockDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionResponseDataBlockDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetTypeName() string {
	return "ConnectionResponseDataBlockDeviceManagement"
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ConnectionResponseDataBlockDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionResponseDataBlockDeviceManagementParse(theBytes []byte) (ConnectionResponseDataBlockDeviceManagement, error) {
	return ConnectionResponseDataBlockDeviceManagementParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ConnectionResponseDataBlockDeviceManagementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionResponseDataBlockDeviceManagement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlockDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionResponseDataBlockDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlockDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionResponseDataBlockDeviceManagement")
	}

	// Create a partially initialized instance
	_child := &_ConnectionResponseDataBlockDeviceManagement{
		_ConnectionResponseDataBlock: &_ConnectionResponseDataBlock{},
	}
	_child._ConnectionResponseDataBlock._ConnectionResponseDataBlockChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionResponseDataBlockDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlockDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionResponseDataBlockDeviceManagement")
		}

		if popErr := writeBuffer.PopContext("ConnectionResponseDataBlockDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionResponseDataBlockDeviceManagement")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionResponseDataBlockDeviceManagement) isConnectionResponseDataBlockDeviceManagement() bool {
	return true
}

func (m *_ConnectionResponseDataBlockDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
