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

// MPropReadCon is the corresponding interface of MPropReadCon
type MPropReadCon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
	// GetData returns Data (property field)
	GetData() uint16
}

// MPropReadConExactly can be used when we want exactly this type and not a type which fulfills MPropReadCon.
// This is useful for switch cases.
type MPropReadConExactly interface {
	MPropReadCon
	isMPropReadCon() bool
}

// _MPropReadCon is the data-structure of this message
type _MPropReadCon struct {
	*_CEMI
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
	Data                uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropReadCon) GetMessageCode() uint8 {
	return 0xFB
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropReadCon) InitializeParent(parent CEMI) {}

func (m *_MPropReadCon) GetParent() CEMI {
	return m._CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MPropReadCon) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *_MPropReadCon) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *_MPropReadCon) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_MPropReadCon) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *_MPropReadCon) GetStartIndex() uint16 {
	return m.StartIndex
}

func (m *_MPropReadCon) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMPropReadCon factory function for _MPropReadCon
func NewMPropReadCon(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, data uint16, size uint16) *_MPropReadCon {
	_result := &_MPropReadCon{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Data:                data,
		_CEMI:               NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropReadCon(structType interface{}) MPropReadCon {
	if casted, ok := structType.(MPropReadCon); ok {
		return casted
	}
	if casted, ok := structType.(*MPropReadCon); ok {
		return *casted
	}
	return nil
}

func (m *_MPropReadCon) GetTypeName() string {
	return "MPropReadCon"
}

func (m *_MPropReadCon) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *_MPropReadCon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MPropReadConParse(theBytes []byte, size uint16) (MPropReadCon, error) {
	return MPropReadConParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), size)
}

func MPropReadConParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (MPropReadCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropReadCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropReadCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceObjectType)
	_interfaceObjectType, _interfaceObjectTypeErr := readBuffer.ReadUint16("interfaceObjectType", 16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.Wrap(_interfaceObjectTypeErr, "Error parsing 'interfaceObjectType' field of MPropReadCon")
	}
	interfaceObjectType := _interfaceObjectType

	// Simple Field (objectInstance)
	_objectInstance, _objectInstanceErr := readBuffer.ReadUint8("objectInstance", 8)
	if _objectInstanceErr != nil {
		return nil, errors.Wrap(_objectInstanceErr, "Error parsing 'objectInstance' field of MPropReadCon")
	}
	objectInstance := _objectInstance

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field of MPropReadCon")
	}
	propertyId := _propertyId

	// Simple Field (numberOfElements)
	_numberOfElements, _numberOfElementsErr := readBuffer.ReadUint8("numberOfElements", 4)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field of MPropReadCon")
	}
	numberOfElements := _numberOfElements

	// Simple Field (startIndex)
	_startIndex, _startIndexErr := readBuffer.ReadUint16("startIndex", 12)
	if _startIndexErr != nil {
		return nil, errors.Wrap(_startIndexErr, "Error parsing 'startIndex' field of MPropReadCon")
	}
	startIndex := _startIndex

	// Simple Field (data)
	_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field of MPropReadCon")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("MPropReadCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropReadCon")
	}

	// Create a partially initialized instance
	_child := &_MPropReadCon{
		_CEMI: &_CEMI{
			Size: size,
		},
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Data:                data,
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MPropReadCon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropReadCon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropReadCon")
		}

		// Simple Field (interfaceObjectType)
		interfaceObjectType := uint16(m.GetInterfaceObjectType())
		_interfaceObjectTypeErr := writeBuffer.WriteUint16("interfaceObjectType", 16, (interfaceObjectType))
		if _interfaceObjectTypeErr != nil {
			return errors.Wrap(_interfaceObjectTypeErr, "Error serializing 'interfaceObjectType' field")
		}

		// Simple Field (objectInstance)
		objectInstance := uint8(m.GetObjectInstance())
		_objectInstanceErr := writeBuffer.WriteUint8("objectInstance", 8, (objectInstance))
		if _objectInstanceErr != nil {
			return errors.Wrap(_objectInstanceErr, "Error serializing 'objectInstance' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (numberOfElements)
		numberOfElements := uint8(m.GetNumberOfElements())
		_numberOfElementsErr := writeBuffer.WriteUint8("numberOfElements", 4, (numberOfElements))
		if _numberOfElementsErr != nil {
			return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
		}

		// Simple Field (startIndex)
		startIndex := uint16(m.GetStartIndex())
		_startIndexErr := writeBuffer.WriteUint16("startIndex", 12, (startIndex))
		if _startIndexErr != nil {
			return errors.Wrap(_startIndexErr, "Error serializing 'startIndex' field")
		}

		// Simple Field (data)
		data := uint16(m.GetData())
		_dataErr := writeBuffer.WriteUint16("data", 16, (data))
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropReadCon")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropReadCon) isMPropReadCon() bool {
	return true
}

func (m *_MPropReadCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
