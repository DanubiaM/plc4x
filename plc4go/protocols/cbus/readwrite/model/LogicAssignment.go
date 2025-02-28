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

// LogicAssignment is the corresponding interface of LogicAssignment
type LogicAssignment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetGreaterOfOrLogic returns GreaterOfOrLogic (property field)
	GetGreaterOfOrLogic() bool
	// GetReStrikeDelay returns ReStrikeDelay (property field)
	GetReStrikeDelay() bool
	// GetAssignedToGav16 returns AssignedToGav16 (property field)
	GetAssignedToGav16() bool
	// GetAssignedToGav15 returns AssignedToGav15 (property field)
	GetAssignedToGav15() bool
	// GetAssignedToGav14 returns AssignedToGav14 (property field)
	GetAssignedToGav14() bool
	// GetAssignedToGav13 returns AssignedToGav13 (property field)
	GetAssignedToGav13() bool
}

// LogicAssignmentExactly can be used when we want exactly this type and not a type which fulfills LogicAssignment.
// This is useful for switch cases.
type LogicAssignmentExactly interface {
	LogicAssignment
	isLogicAssignment() bool
}

// _LogicAssignment is the data-structure of this message
type _LogicAssignment struct {
	GreaterOfOrLogic bool
	ReStrikeDelay    bool
	AssignedToGav16  bool
	AssignedToGav15  bool
	AssignedToGav14  bool
	AssignedToGav13  bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LogicAssignment) GetGreaterOfOrLogic() bool {
	return m.GreaterOfOrLogic
}

func (m *_LogicAssignment) GetReStrikeDelay() bool {
	return m.ReStrikeDelay
}

func (m *_LogicAssignment) GetAssignedToGav16() bool {
	return m.AssignedToGav16
}

func (m *_LogicAssignment) GetAssignedToGav15() bool {
	return m.AssignedToGav15
}

func (m *_LogicAssignment) GetAssignedToGav14() bool {
	return m.AssignedToGav14
}

func (m *_LogicAssignment) GetAssignedToGav13() bool {
	return m.AssignedToGav13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLogicAssignment factory function for _LogicAssignment
func NewLogicAssignment(greaterOfOrLogic bool, reStrikeDelay bool, assignedToGav16 bool, assignedToGav15 bool, assignedToGav14 bool, assignedToGav13 bool) *_LogicAssignment {
	return &_LogicAssignment{GreaterOfOrLogic: greaterOfOrLogic, ReStrikeDelay: reStrikeDelay, AssignedToGav16: assignedToGav16, AssignedToGav15: assignedToGav15, AssignedToGav14: assignedToGav14, AssignedToGav13: assignedToGav13}
}

// Deprecated: use the interface for direct cast
func CastLogicAssignment(structType interface{}) LogicAssignment {
	if casted, ok := structType.(LogicAssignment); ok {
		return casted
	}
	if casted, ok := structType.(*LogicAssignment); ok {
		return *casted
	}
	return nil
}

func (m *_LogicAssignment) GetTypeName() string {
	return "LogicAssignment"
}

func (m *_LogicAssignment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (greaterOfOrLogic)
	lengthInBits += 1

	// Simple field (reStrikeDelay)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (assignedToGav16)
	lengthInBits += 1

	// Simple field (assignedToGav15)
	lengthInBits += 1

	// Simple field (assignedToGav14)
	lengthInBits += 1

	// Simple field (assignedToGav13)
	lengthInBits += 1

	return lengthInBits
}

func (m *_LogicAssignment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LogicAssignmentParse(theBytes []byte) (LogicAssignment, error) {
	return LogicAssignmentParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func LogicAssignmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LogicAssignment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LogicAssignment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LogicAssignment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (greaterOfOrLogic)
	_greaterOfOrLogic, _greaterOfOrLogicErr := readBuffer.ReadBit("greaterOfOrLogic")
	if _greaterOfOrLogicErr != nil {
		return nil, errors.Wrap(_greaterOfOrLogicErr, "Error parsing 'greaterOfOrLogic' field of LogicAssignment")
	}
	greaterOfOrLogic := _greaterOfOrLogic

	// Simple Field (reStrikeDelay)
	_reStrikeDelay, _reStrikeDelayErr := readBuffer.ReadBit("reStrikeDelay")
	if _reStrikeDelayErr != nil {
		return nil, errors.Wrap(_reStrikeDelayErr, "Error parsing 'reStrikeDelay' field of LogicAssignment")
	}
	reStrikeDelay := _reStrikeDelay

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LogicAssignment")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LogicAssignment")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (assignedToGav16)
	_assignedToGav16, _assignedToGav16Err := readBuffer.ReadBit("assignedToGav16")
	if _assignedToGav16Err != nil {
		return nil, errors.Wrap(_assignedToGav16Err, "Error parsing 'assignedToGav16' field of LogicAssignment")
	}
	assignedToGav16 := _assignedToGav16

	// Simple Field (assignedToGav15)
	_assignedToGav15, _assignedToGav15Err := readBuffer.ReadBit("assignedToGav15")
	if _assignedToGav15Err != nil {
		return nil, errors.Wrap(_assignedToGav15Err, "Error parsing 'assignedToGav15' field of LogicAssignment")
	}
	assignedToGav15 := _assignedToGav15

	// Simple Field (assignedToGav14)
	_assignedToGav14, _assignedToGav14Err := readBuffer.ReadBit("assignedToGav14")
	if _assignedToGav14Err != nil {
		return nil, errors.Wrap(_assignedToGav14Err, "Error parsing 'assignedToGav14' field of LogicAssignment")
	}
	assignedToGav14 := _assignedToGav14

	// Simple Field (assignedToGav13)
	_assignedToGav13, _assignedToGav13Err := readBuffer.ReadBit("assignedToGav13")
	if _assignedToGav13Err != nil {
		return nil, errors.Wrap(_assignedToGav13Err, "Error parsing 'assignedToGav13' field of LogicAssignment")
	}
	assignedToGav13 := _assignedToGav13

	if closeErr := readBuffer.CloseContext("LogicAssignment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LogicAssignment")
	}

	// Create the instance
	return &_LogicAssignment{
		GreaterOfOrLogic: greaterOfOrLogic,
		ReStrikeDelay:    reStrikeDelay,
		AssignedToGav16:  assignedToGav16,
		AssignedToGav15:  assignedToGav15,
		AssignedToGav14:  assignedToGav14,
		AssignedToGav13:  assignedToGav13,
		reservedField0:   reservedField0,
		reservedField1:   reservedField1,
	}, nil
}

func (m *_LogicAssignment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LogicAssignment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("LogicAssignment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LogicAssignment")
	}

	// Simple Field (greaterOfOrLogic)
	greaterOfOrLogic := bool(m.GetGreaterOfOrLogic())
	_greaterOfOrLogicErr := writeBuffer.WriteBit("greaterOfOrLogic", (greaterOfOrLogic))
	if _greaterOfOrLogicErr != nil {
		return errors.Wrap(_greaterOfOrLogicErr, "Error serializing 'greaterOfOrLogic' field")
	}

	// Simple Field (reStrikeDelay)
	reStrikeDelay := bool(m.GetReStrikeDelay())
	_reStrikeDelayErr := writeBuffer.WriteBit("reStrikeDelay", (reStrikeDelay))
	if _reStrikeDelayErr != nil {
		return errors.Wrap(_reStrikeDelayErr, "Error serializing 'reStrikeDelay' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField1 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (assignedToGav16)
	assignedToGav16 := bool(m.GetAssignedToGav16())
	_assignedToGav16Err := writeBuffer.WriteBit("assignedToGav16", (assignedToGav16))
	if _assignedToGav16Err != nil {
		return errors.Wrap(_assignedToGav16Err, "Error serializing 'assignedToGav16' field")
	}

	// Simple Field (assignedToGav15)
	assignedToGav15 := bool(m.GetAssignedToGav15())
	_assignedToGav15Err := writeBuffer.WriteBit("assignedToGav15", (assignedToGav15))
	if _assignedToGav15Err != nil {
		return errors.Wrap(_assignedToGav15Err, "Error serializing 'assignedToGav15' field")
	}

	// Simple Field (assignedToGav14)
	assignedToGav14 := bool(m.GetAssignedToGav14())
	_assignedToGav14Err := writeBuffer.WriteBit("assignedToGav14", (assignedToGav14))
	if _assignedToGav14Err != nil {
		return errors.Wrap(_assignedToGav14Err, "Error serializing 'assignedToGav14' field")
	}

	// Simple Field (assignedToGav13)
	assignedToGav13 := bool(m.GetAssignedToGav13())
	_assignedToGav13Err := writeBuffer.WriteBit("assignedToGav13", (assignedToGav13))
	if _assignedToGav13Err != nil {
		return errors.Wrap(_assignedToGav13Err, "Error serializing 'assignedToGav13' field")
	}

	if popErr := writeBuffer.PopContext("LogicAssignment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LogicAssignment")
	}
	return nil
}

func (m *_LogicAssignment) isLogicAssignment() bool {
	return true
}

func (m *_LogicAssignment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
