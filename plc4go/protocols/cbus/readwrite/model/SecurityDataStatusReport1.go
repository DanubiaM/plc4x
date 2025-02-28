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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataStatusReport1 is the corresponding interface of SecurityDataStatusReport1
type SecurityDataStatusReport1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetArmCodeType returns ArmCodeType (property field)
	GetArmCodeType() SecurityArmCode
	// GetTamperStatus returns TamperStatus (property field)
	GetTamperStatus() TamperStatus
	// GetPanicStatus returns PanicStatus (property field)
	GetPanicStatus() PanicStatus
	// GetZoneStatus returns ZoneStatus (property field)
	GetZoneStatus() []ZoneStatus
}

// SecurityDataStatusReport1Exactly can be used when we want exactly this type and not a type which fulfills SecurityDataStatusReport1.
// This is useful for switch cases.
type SecurityDataStatusReport1Exactly interface {
	SecurityDataStatusReport1
	isSecurityDataStatusReport1() bool
}

// _SecurityDataStatusReport1 is the data-structure of this message
type _SecurityDataStatusReport1 struct {
	*_SecurityData
	ArmCodeType  SecurityArmCode
	TamperStatus TamperStatus
	PanicStatus  PanicStatus
	ZoneStatus   []ZoneStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatusReport1) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataStatusReport1) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataStatusReport1) GetArmCodeType() SecurityArmCode {
	return m.ArmCodeType
}

func (m *_SecurityDataStatusReport1) GetTamperStatus() TamperStatus {
	return m.TamperStatus
}

func (m *_SecurityDataStatusReport1) GetPanicStatus() PanicStatus {
	return m.PanicStatus
}

func (m *_SecurityDataStatusReport1) GetZoneStatus() []ZoneStatus {
	return m.ZoneStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataStatusReport1 factory function for _SecurityDataStatusReport1
func NewSecurityDataStatusReport1(armCodeType SecurityArmCode, tamperStatus TamperStatus, panicStatus PanicStatus, zoneStatus []ZoneStatus, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatusReport1 {
	_result := &_SecurityDataStatusReport1{
		ArmCodeType:   armCodeType,
		TamperStatus:  tamperStatus,
		PanicStatus:   panicStatus,
		ZoneStatus:    zoneStatus,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatusReport1(structType interface{}) SecurityDataStatusReport1 {
	if casted, ok := structType.(SecurityDataStatusReport1); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatusReport1); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatusReport1) GetTypeName() string {
	return "SecurityDataStatusReport1"
}

func (m *_SecurityDataStatusReport1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (armCodeType)
	lengthInBits += m.ArmCodeType.GetLengthInBits(ctx)

	// Simple field (tamperStatus)
	lengthInBits += m.TamperStatus.GetLengthInBits(ctx)

	// Simple field (panicStatus)
	lengthInBits += m.PanicStatus.GetLengthInBits(ctx)

	// Array field
	if len(m.ZoneStatus) > 0 {
		for _curItem, element := range m.ZoneStatus {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.ZoneStatus), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityDataStatusReport1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataStatusReport1Parse(theBytes []byte) (SecurityDataStatusReport1, error) {
	return SecurityDataStatusReport1ParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataStatusReport1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataStatusReport1, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatusReport1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatusReport1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (armCodeType)
	if pullErr := readBuffer.PullContext("armCodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for armCodeType")
	}
	_armCodeType, _armCodeTypeErr := SecurityArmCodeParseWithBuffer(ctx, readBuffer)
	if _armCodeTypeErr != nil {
		return nil, errors.Wrap(_armCodeTypeErr, "Error parsing 'armCodeType' field of SecurityDataStatusReport1")
	}
	armCodeType := _armCodeType.(SecurityArmCode)
	if closeErr := readBuffer.CloseContext("armCodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for armCodeType")
	}

	// Simple Field (tamperStatus)
	if pullErr := readBuffer.PullContext("tamperStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tamperStatus")
	}
	_tamperStatus, _tamperStatusErr := TamperStatusParseWithBuffer(ctx, readBuffer)
	if _tamperStatusErr != nil {
		return nil, errors.Wrap(_tamperStatusErr, "Error parsing 'tamperStatus' field of SecurityDataStatusReport1")
	}
	tamperStatus := _tamperStatus.(TamperStatus)
	if closeErr := readBuffer.CloseContext("tamperStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tamperStatus")
	}

	// Simple Field (panicStatus)
	if pullErr := readBuffer.PullContext("panicStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for panicStatus")
	}
	_panicStatus, _panicStatusErr := PanicStatusParseWithBuffer(ctx, readBuffer)
	if _panicStatusErr != nil {
		return nil, errors.Wrap(_panicStatusErr, "Error parsing 'panicStatus' field of SecurityDataStatusReport1")
	}
	panicStatus := _panicStatus.(PanicStatus)
	if closeErr := readBuffer.CloseContext("panicStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for panicStatus")
	}

	// Array field (zoneStatus)
	if pullErr := readBuffer.PullContext("zoneStatus", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneStatus")
	}
	// Count array
	zoneStatus := make([]ZoneStatus, uint16(32))
	// This happens when the size is set conditional to 0
	if len(zoneStatus) == 0 {
		zoneStatus = nil
	}
	{
		_numItems := uint16(uint16(32))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ZoneStatusParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'zoneStatus' field of SecurityDataStatusReport1")
			}
			zoneStatus[_curItem] = _item.(ZoneStatus)
		}
	}
	if closeErr := readBuffer.CloseContext("zoneStatus", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneStatus")
	}

	if closeErr := readBuffer.CloseContext("SecurityDataStatusReport1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatusReport1")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataStatusReport1{
		_SecurityData: &_SecurityData{},
		ArmCodeType:   armCodeType,
		TamperStatus:  tamperStatus,
		PanicStatus:   panicStatus,
		ZoneStatus:    zoneStatus,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataStatusReport1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatusReport1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatusReport1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatusReport1")
		}

		// Simple Field (armCodeType)
		if pushErr := writeBuffer.PushContext("armCodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for armCodeType")
		}
		_armCodeTypeErr := writeBuffer.WriteSerializable(ctx, m.GetArmCodeType())
		if popErr := writeBuffer.PopContext("armCodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for armCodeType")
		}
		if _armCodeTypeErr != nil {
			return errors.Wrap(_armCodeTypeErr, "Error serializing 'armCodeType' field")
		}

		// Simple Field (tamperStatus)
		if pushErr := writeBuffer.PushContext("tamperStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for tamperStatus")
		}
		_tamperStatusErr := writeBuffer.WriteSerializable(ctx, m.GetTamperStatus())
		if popErr := writeBuffer.PopContext("tamperStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for tamperStatus")
		}
		if _tamperStatusErr != nil {
			return errors.Wrap(_tamperStatusErr, "Error serializing 'tamperStatus' field")
		}

		// Simple Field (panicStatus)
		if pushErr := writeBuffer.PushContext("panicStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for panicStatus")
		}
		_panicStatusErr := writeBuffer.WriteSerializable(ctx, m.GetPanicStatus())
		if popErr := writeBuffer.PopContext("panicStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for panicStatus")
		}
		if _panicStatusErr != nil {
			return errors.Wrap(_panicStatusErr, "Error serializing 'panicStatus' field")
		}

		// Array Field (zoneStatus)
		if pushErr := writeBuffer.PushContext("zoneStatus", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneStatus")
		}
		for _curItem, _element := range m.GetZoneStatus() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetZoneStatus()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'zoneStatus' field")
			}
		}
		if popErr := writeBuffer.PopContext("zoneStatus", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneStatus")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatusReport1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatusReport1")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatusReport1) isSecurityDataStatusReport1() bool {
	return true
}

func (m *_SecurityDataStatusReport1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
