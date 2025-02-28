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

// S7AddressAny is the corresponding interface of S7AddressAny
type S7AddressAny interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Address
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() TransportSize
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint16
	// GetDbNumber returns DbNumber (property field)
	GetDbNumber() uint16
	// GetArea returns Area (property field)
	GetArea() MemoryArea
	// GetByteAddress returns ByteAddress (property field)
	GetByteAddress() uint16
	// GetBitAddress returns BitAddress (property field)
	GetBitAddress() uint8
}

// S7AddressAnyExactly can be used when we want exactly this type and not a type which fulfills S7AddressAny.
// This is useful for switch cases.
type S7AddressAnyExactly interface {
	S7AddressAny
	isS7AddressAny() bool
}

// _S7AddressAny is the data-structure of this message
type _S7AddressAny struct {
	*_S7Address
	TransportSize    TransportSize
	NumberOfElements uint16
	DbNumber         uint16
	Area             MemoryArea
	ByteAddress      uint16
	BitAddress       uint8
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7AddressAny) GetAddressType() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7AddressAny) InitializeParent(parent S7Address) {}

func (m *_S7AddressAny) GetParent() S7Address {
	return m._S7Address
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7AddressAny) GetTransportSize() TransportSize {
	return m.TransportSize
}

func (m *_S7AddressAny) GetNumberOfElements() uint16 {
	return m.NumberOfElements
}

func (m *_S7AddressAny) GetDbNumber() uint16 {
	return m.DbNumber
}

func (m *_S7AddressAny) GetArea() MemoryArea {
	return m.Area
}

func (m *_S7AddressAny) GetByteAddress() uint16 {
	return m.ByteAddress
}

func (m *_S7AddressAny) GetBitAddress() uint8 {
	return m.BitAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7AddressAny factory function for _S7AddressAny
func NewS7AddressAny(transportSize TransportSize, numberOfElements uint16, dbNumber uint16, area MemoryArea, byteAddress uint16, bitAddress uint8) *_S7AddressAny {
	_result := &_S7AddressAny{
		TransportSize:    transportSize,
		NumberOfElements: numberOfElements,
		DbNumber:         dbNumber,
		Area:             area,
		ByteAddress:      byteAddress,
		BitAddress:       bitAddress,
		_S7Address:       NewS7Address(),
	}
	_result._S7Address._S7AddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7AddressAny(structType interface{}) S7AddressAny {
	if casted, ok := structType.(S7AddressAny); ok {
		return casted
	}
	if casted, ok := structType.(*S7AddressAny); ok {
		return *casted
	}
	return nil
}

func (m *_S7AddressAny) GetTypeName() string {
	return "S7AddressAny"
}

func (m *_S7AddressAny) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Enum Field (transportSize)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 16

	// Simple field (dbNumber)
	lengthInBits += 16

	// Simple field (area)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 5

	// Simple field (byteAddress)
	lengthInBits += 16

	// Simple field (bitAddress)
	lengthInBits += 3

	return lengthInBits
}

func (m *_S7AddressAny) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7AddressAnyParse(theBytes []byte) (S7AddressAny, error) {
	return S7AddressAnyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func S7AddressAnyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (S7AddressAny, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7AddressAny"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7AddressAny")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	// Enum field (transportSize)
	transportSizeCode, _transportSizeCodeErr := readBuffer.ReadUint8("TransportSize", 8)
	if _transportSizeCodeErr != nil {
		return nil, errors.Wrap(_transportSizeCodeErr, "Error serializing 'transportSize' field")
	}
	transportSize, _transportSizeErr := TransportSizeFirstEnumForFieldCode(transportSizeCode)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	// Simple Field (numberOfElements)
	_numberOfElements, _numberOfElementsErr := readBuffer.ReadUint16("numberOfElements", 16)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field of S7AddressAny")
	}
	numberOfElements := _numberOfElements

	// Simple Field (dbNumber)
	_dbNumber, _dbNumberErr := readBuffer.ReadUint16("dbNumber", 16)
	if _dbNumberErr != nil {
		return nil, errors.Wrap(_dbNumberErr, "Error parsing 'dbNumber' field of S7AddressAny")
	}
	dbNumber := _dbNumber

	// Simple Field (area)
	if pullErr := readBuffer.PullContext("area"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for area")
	}
	_area, _areaErr := MemoryAreaParseWithBuffer(ctx, readBuffer)
	if _areaErr != nil {
		return nil, errors.Wrap(_areaErr, "Error parsing 'area' field of S7AddressAny")
	}
	area := _area
	if closeErr := readBuffer.CloseContext("area"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for area")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 5)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7AddressAny")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (byteAddress)
	_byteAddress, _byteAddressErr := readBuffer.ReadUint16("byteAddress", 16)
	if _byteAddressErr != nil {
		return nil, errors.Wrap(_byteAddressErr, "Error parsing 'byteAddress' field of S7AddressAny")
	}
	byteAddress := _byteAddress

	// Simple Field (bitAddress)
	_bitAddress, _bitAddressErr := readBuffer.ReadUint8("bitAddress", 3)
	if _bitAddressErr != nil {
		return nil, errors.Wrap(_bitAddressErr, "Error parsing 'bitAddress' field of S7AddressAny")
	}
	bitAddress := _bitAddress

	if closeErr := readBuffer.CloseContext("S7AddressAny"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7AddressAny")
	}

	// Create a partially initialized instance
	_child := &_S7AddressAny{
		_S7Address:       &_S7Address{},
		TransportSize:    transportSize,
		NumberOfElements: numberOfElements,
		DbNumber:         dbNumber,
		Area:             area,
		ByteAddress:      byteAddress,
		BitAddress:       bitAddress,
		reservedField0:   reservedField0,
	}
	_child._S7Address._S7AddressChildRequirements = _child
	return _child, nil
}

func (m *_S7AddressAny) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7AddressAny) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7AddressAny"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7AddressAny")
		}

		if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportSize")
		}
		// Enum field (transportSize)
		_transportSizeErr := writeBuffer.WriteUint8("TransportSize", 8, m.TransportSize.Code(), utils.WithAdditionalStringRepresentation(m.GetTransportSize().PLC4XEnumName()))
		if _transportSizeErr != nil {
			return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
		}
		if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportSize")
		}

		// Simple Field (numberOfElements)
		numberOfElements := uint16(m.GetNumberOfElements())
		_numberOfElementsErr := writeBuffer.WriteUint16("numberOfElements", 16, (numberOfElements))
		if _numberOfElementsErr != nil {
			return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
		}

		// Simple Field (dbNumber)
		dbNumber := uint16(m.GetDbNumber())
		_dbNumberErr := writeBuffer.WriteUint16("dbNumber", 16, (dbNumber))
		if _dbNumberErr != nil {
			return errors.Wrap(_dbNumberErr, "Error serializing 'dbNumber' field")
		}

		// Simple Field (area)
		if pushErr := writeBuffer.PushContext("area"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for area")
		}
		_areaErr := writeBuffer.WriteSerializable(ctx, m.GetArea())
		if popErr := writeBuffer.PopContext("area"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for area")
		}
		if _areaErr != nil {
			return errors.Wrap(_areaErr, "Error serializing 'area' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 5, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (byteAddress)
		byteAddress := uint16(m.GetByteAddress())
		_byteAddressErr := writeBuffer.WriteUint16("byteAddress", 16, (byteAddress))
		if _byteAddressErr != nil {
			return errors.Wrap(_byteAddressErr, "Error serializing 'byteAddress' field")
		}

		// Simple Field (bitAddress)
		bitAddress := uint8(m.GetBitAddress())
		_bitAddressErr := writeBuffer.WriteUint8("bitAddress", 3, (bitAddress))
		if _bitAddressErr != nil {
			return errors.Wrap(_bitAddressErr, "Error serializing 'bitAddress' field")
		}

		if popErr := writeBuffer.PopContext("S7AddressAny"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7AddressAny")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7AddressAny) isS7AddressAny() bool {
	return true
}

func (m *_S7AddressAny) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
