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

// ConnectedAddressItem is the corresponding interface of ConnectedAddressItem
type ConnectedAddressItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetConnectionId returns ConnectionId (property field)
	GetConnectionId() uint32
}

// ConnectedAddressItemExactly can be used when we want exactly this type and not a type which fulfills ConnectedAddressItem.
// This is useful for switch cases.
type ConnectedAddressItemExactly interface {
	ConnectedAddressItem
	isConnectedAddressItem() bool
}

// _ConnectedAddressItem is the data-structure of this message
type _ConnectedAddressItem struct {
	*_TypeId
	ConnectionId uint32
	// Reserved Fields
	reservedField0 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectedAddressItem) GetId() uint16 {
	return 0x00A1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectedAddressItem) InitializeParent(parent TypeId) {}

func (m *_ConnectedAddressItem) GetParent() TypeId {
	return m._TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ConnectedAddressItem) GetConnectionId() uint32 {
	return m.ConnectionId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConnectedAddressItem factory function for _ConnectedAddressItem
func NewConnectedAddressItem(connectionId uint32) *_ConnectedAddressItem {
	_result := &_ConnectedAddressItem{
		ConnectionId: connectionId,
		_TypeId:      NewTypeId(),
	}
	_result._TypeId._TypeIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectedAddressItem(structType interface{}) ConnectedAddressItem {
	if casted, ok := structType.(ConnectedAddressItem); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectedAddressItem); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectedAddressItem) GetTypeName() string {
	return "ConnectedAddressItem"
}

func (m *_ConnectedAddressItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (connectionId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ConnectedAddressItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectedAddressItemParse(theBytes []byte) (ConnectedAddressItem, error) {
	return ConnectedAddressItemParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ConnectedAddressItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectedAddressItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectedAddressItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectedAddressItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ConnectedAddressItem")
		}
		if reserved != uint16(0x0004) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0004),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (connectionId)
	_connectionId, _connectionIdErr := readBuffer.ReadUint32("connectionId", 32)
	if _connectionIdErr != nil {
		return nil, errors.Wrap(_connectionIdErr, "Error parsing 'connectionId' field of ConnectedAddressItem")
	}
	connectionId := _connectionId

	if closeErr := readBuffer.CloseContext("ConnectedAddressItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectedAddressItem")
	}

	// Create a partially initialized instance
	_child := &_ConnectedAddressItem{
		_TypeId:        &_TypeId{},
		ConnectionId:   connectionId,
		reservedField0: reservedField0,
	}
	_child._TypeId._TypeIdChildRequirements = _child
	return _child, nil
}

func (m *_ConnectedAddressItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectedAddressItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectedAddressItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectedAddressItem")
		}

		// Reserved Field (reserved)
		{
			var reserved uint16 = uint16(0x0004)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint16(0x0004),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint16("reserved", 16, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (connectionId)
		connectionId := uint32(m.GetConnectionId())
		_connectionIdErr := writeBuffer.WriteUint32("connectionId", 32, (connectionId))
		if _connectionIdErr != nil {
			return errors.Wrap(_connectionIdErr, "Error serializing 'connectionId' field")
		}

		if popErr := writeBuffer.PopContext("ConnectedAddressItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectedAddressItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectedAddressItem) isConnectedAddressItem() bool {
	return true
}

func (m *_ConnectedAddressItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
