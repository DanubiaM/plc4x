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

// SecurityDataSystemArmedDisarmed is the corresponding interface of SecurityDataSystemArmedDisarmed
type SecurityDataSystemArmedDisarmed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetArmCodeType returns ArmCodeType (property field)
	GetArmCodeType() SecurityArmCode
}

// SecurityDataSystemArmedDisarmedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataSystemArmedDisarmed.
// This is useful for switch cases.
type SecurityDataSystemArmedDisarmedExactly interface {
	SecurityDataSystemArmedDisarmed
	isSecurityDataSystemArmedDisarmed() bool
}

// _SecurityDataSystemArmedDisarmed is the data-structure of this message
type _SecurityDataSystemArmedDisarmed struct {
	*_SecurityData
	ArmCodeType SecurityArmCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataSystemArmedDisarmed) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataSystemArmedDisarmed) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataSystemArmedDisarmed) GetArmCodeType() SecurityArmCode {
	return m.ArmCodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataSystemArmedDisarmed factory function for _SecurityDataSystemArmedDisarmed
func NewSecurityDataSystemArmedDisarmed(armCodeType SecurityArmCode, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataSystemArmedDisarmed {
	_result := &_SecurityDataSystemArmedDisarmed{
		ArmCodeType:   armCodeType,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataSystemArmedDisarmed(structType interface{}) SecurityDataSystemArmedDisarmed {
	if casted, ok := structType.(SecurityDataSystemArmedDisarmed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataSystemArmedDisarmed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataSystemArmedDisarmed) GetTypeName() string {
	return "SecurityDataSystemArmedDisarmed"
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (armCodeType)
	lengthInBits += m.ArmCodeType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SecurityDataSystemArmedDisarmed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataSystemArmedDisarmedParse(theBytes []byte) (SecurityDataSystemArmedDisarmed, error) {
	return SecurityDataSystemArmedDisarmedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataSystemArmedDisarmedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataSystemArmedDisarmed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataSystemArmedDisarmed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataSystemArmedDisarmed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (armCodeType)
	if pullErr := readBuffer.PullContext("armCodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for armCodeType")
	}
	_armCodeType, _armCodeTypeErr := SecurityArmCodeParseWithBuffer(ctx, readBuffer)
	if _armCodeTypeErr != nil {
		return nil, errors.Wrap(_armCodeTypeErr, "Error parsing 'armCodeType' field of SecurityDataSystemArmedDisarmed")
	}
	armCodeType := _armCodeType.(SecurityArmCode)
	if closeErr := readBuffer.CloseContext("armCodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for armCodeType")
	}

	if closeErr := readBuffer.CloseContext("SecurityDataSystemArmedDisarmed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataSystemArmedDisarmed")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataSystemArmedDisarmed{
		_SecurityData: &_SecurityData{},
		ArmCodeType:   armCodeType,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataSystemArmedDisarmed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataSystemArmedDisarmed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataSystemArmedDisarmed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataSystemArmedDisarmed")
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

		if popErr := writeBuffer.PopContext("SecurityDataSystemArmedDisarmed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataSystemArmedDisarmed")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataSystemArmedDisarmed) isSecurityDataSystemArmedDisarmed() bool {
	return true
}

func (m *_SecurityDataSystemArmedDisarmed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
