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

// SecurityDataRaiseAlarm is the corresponding interface of SecurityDataRaiseAlarm
type SecurityDataRaiseAlarm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataRaiseAlarmExactly can be used when we want exactly this type and not a type which fulfills SecurityDataRaiseAlarm.
// This is useful for switch cases.
type SecurityDataRaiseAlarmExactly interface {
	SecurityDataRaiseAlarm
	isSecurityDataRaiseAlarm() bool
}

// _SecurityDataRaiseAlarm is the data-structure of this message
type _SecurityDataRaiseAlarm struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataRaiseAlarm) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataRaiseAlarm) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataRaiseAlarm factory function for _SecurityDataRaiseAlarm
func NewSecurityDataRaiseAlarm(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataRaiseAlarm {
	_result := &_SecurityDataRaiseAlarm{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataRaiseAlarm(structType interface{}) SecurityDataRaiseAlarm {
	if casted, ok := structType.(SecurityDataRaiseAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataRaiseAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataRaiseAlarm) GetTypeName() string {
	return "SecurityDataRaiseAlarm"
}

func (m *_SecurityDataRaiseAlarm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataRaiseAlarm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataRaiseAlarmParse(theBytes []byte) (SecurityDataRaiseAlarm, error) {
	return SecurityDataRaiseAlarmParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataRaiseAlarmParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataRaiseAlarm, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataRaiseAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataRaiseAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataRaiseAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataRaiseAlarm")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataRaiseAlarm{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataRaiseAlarm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataRaiseAlarm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataRaiseAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataRaiseAlarm")
		}

		if popErr := writeBuffer.PopContext("SecurityDataRaiseAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataRaiseAlarm")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataRaiseAlarm) isSecurityDataRaiseAlarm() bool {
	return true
}

func (m *_SecurityDataRaiseAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
