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

// BACnetHostNPort is the corresponding interface of BACnetHostNPort
type BACnetHostNPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHost returns Host (property field)
	GetHost() BACnetHostAddressEnclosed
	// GetPort returns Port (property field)
	GetPort() BACnetContextTagUnsignedInteger
}

// BACnetHostNPortExactly can be used when we want exactly this type and not a type which fulfills BACnetHostNPort.
// This is useful for switch cases.
type BACnetHostNPortExactly interface {
	BACnetHostNPort
	isBACnetHostNPort() bool
}

// _BACnetHostNPort is the data-structure of this message
type _BACnetHostNPort struct {
	Host BACnetHostAddressEnclosed
	Port BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostNPort) GetHost() BACnetHostAddressEnclosed {
	return m.Host
}

func (m *_BACnetHostNPort) GetPort() BACnetContextTagUnsignedInteger {
	return m.Port
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostNPort factory function for _BACnetHostNPort
func NewBACnetHostNPort(host BACnetHostAddressEnclosed, port BACnetContextTagUnsignedInteger) *_BACnetHostNPort {
	return &_BACnetHostNPort{Host: host, Port: port}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostNPort(structType interface{}) BACnetHostNPort {
	if casted, ok := structType.(BACnetHostNPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostNPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostNPort) GetTypeName() string {
	return "BACnetHostNPort"
}

func (m *_BACnetHostNPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (host)
	lengthInBits += m.Host.GetLengthInBits(ctx)

	// Simple field (port)
	lengthInBits += m.Port.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostNPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostNPortParse(theBytes []byte) (BACnetHostNPort, error) {
	return BACnetHostNPortParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetHostNPortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (host)
	if pullErr := readBuffer.PullContext("host"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for host")
	}
	_host, _hostErr := BACnetHostAddressEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _hostErr != nil {
		return nil, errors.Wrap(_hostErr, "Error parsing 'host' field of BACnetHostNPort")
	}
	host := _host.(BACnetHostAddressEnclosed)
	if closeErr := readBuffer.CloseContext("host"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for host")
	}

	// Simple Field (port)
	if pullErr := readBuffer.PullContext("port"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for port")
	}
	_port, _portErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BACnetHostNPort")
	}
	port := _port.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("port"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for port")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostNPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPort")
	}

	// Create the instance
	return &_BACnetHostNPort{
		Host: host,
		Port: port,
	}, nil
}

func (m *_BACnetHostNPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostNPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostNPort"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPort")
	}

	// Simple Field (host)
	if pushErr := writeBuffer.PushContext("host"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for host")
	}
	_hostErr := writeBuffer.WriteSerializable(ctx, m.GetHost())
	if popErr := writeBuffer.PopContext("host"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for host")
	}
	if _hostErr != nil {
		return errors.Wrap(_hostErr, "Error serializing 'host' field")
	}

	// Simple Field (port)
	if pushErr := writeBuffer.PushContext("port"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for port")
	}
	_portErr := writeBuffer.WriteSerializable(ctx, m.GetPort())
	if popErr := writeBuffer.PopContext("port"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for port")
	}
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostNPort"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPort")
	}
	return nil
}

func (m *_BACnetHostNPort) isBACnetHostNPort() bool {
	return true
}

func (m *_BACnetHostNPort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
