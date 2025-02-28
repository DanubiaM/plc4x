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

// LightingLabelOptions is the corresponding interface of LightingLabelOptions
type LightingLabelOptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLabelFlavour returns LabelFlavour (property field)
	GetLabelFlavour() LightingLabelFlavour
	// GetLabelType returns LabelType (property field)
	GetLabelType() LightingLabelType
}

// LightingLabelOptionsExactly can be used when we want exactly this type and not a type which fulfills LightingLabelOptions.
// This is useful for switch cases.
type LightingLabelOptionsExactly interface {
	LightingLabelOptions
	isLightingLabelOptions() bool
}

// _LightingLabelOptions is the data-structure of this message
type _LightingLabelOptions struct {
	LabelFlavour LightingLabelFlavour
	LabelType    LightingLabelType
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingLabelOptions) GetLabelFlavour() LightingLabelFlavour {
	return m.LabelFlavour
}

func (m *_LightingLabelOptions) GetLabelType() LightingLabelType {
	return m.LabelType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingLabelOptions factory function for _LightingLabelOptions
func NewLightingLabelOptions(labelFlavour LightingLabelFlavour, labelType LightingLabelType) *_LightingLabelOptions {
	return &_LightingLabelOptions{LabelFlavour: labelFlavour, LabelType: labelType}
}

// Deprecated: use the interface for direct cast
func CastLightingLabelOptions(structType interface{}) LightingLabelOptions {
	if casted, ok := structType.(LightingLabelOptions); ok {
		return casted
	}
	if casted, ok := structType.(*LightingLabelOptions); ok {
		return *casted
	}
	return nil
}

func (m *_LightingLabelOptions) GetTypeName() string {
	return "LightingLabelOptions"
}

func (m *_LightingLabelOptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelFlavour)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_LightingLabelOptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingLabelOptionsParse(theBytes []byte) (LightingLabelOptions, error) {
	return LightingLabelOptionsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func LightingLabelOptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingLabelOptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingLabelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingLabelOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LightingLabelOptions")
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

	// Simple Field (labelFlavour)
	if pullErr := readBuffer.PullContext("labelFlavour"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for labelFlavour")
	}
	_labelFlavour, _labelFlavourErr := LightingLabelFlavourParseWithBuffer(ctx, readBuffer)
	if _labelFlavourErr != nil {
		return nil, errors.Wrap(_labelFlavourErr, "Error parsing 'labelFlavour' field of LightingLabelOptions")
	}
	labelFlavour := _labelFlavour
	if closeErr := readBuffer.CloseContext("labelFlavour"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for labelFlavour")
	}

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LightingLabelOptions")
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

	var reservedField2 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LightingLabelOptions")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField2 = &reserved
		}
	}

	// Simple Field (labelType)
	if pullErr := readBuffer.PullContext("labelType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for labelType")
	}
	_labelType, _labelTypeErr := LightingLabelTypeParseWithBuffer(ctx, readBuffer)
	if _labelTypeErr != nil {
		return nil, errors.Wrap(_labelTypeErr, "Error parsing 'labelType' field of LightingLabelOptions")
	}
	labelType := _labelType
	if closeErr := readBuffer.CloseContext("labelType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for labelType")
	}

	var reservedField3 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of LightingLabelOptions")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField3 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("LightingLabelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingLabelOptions")
	}

	// Create the instance
	return &_LightingLabelOptions{
		LabelFlavour:   labelFlavour,
		LabelType:      labelType,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
		reservedField2: reservedField2,
		reservedField3: reservedField3,
	}, nil
}

func (m *_LightingLabelOptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingLabelOptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("LightingLabelOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LightingLabelOptions")
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

	// Simple Field (labelFlavour)
	if pushErr := writeBuffer.PushContext("labelFlavour"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelFlavour")
	}
	_labelFlavourErr := writeBuffer.WriteSerializable(ctx, m.GetLabelFlavour())
	if popErr := writeBuffer.PopContext("labelFlavour"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelFlavour")
	}
	if _labelFlavourErr != nil {
		return errors.Wrap(_labelFlavourErr, "Error serializing 'labelFlavour' field")
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

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField2 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField2
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (labelType)
	if pushErr := writeBuffer.PushContext("labelType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelType")
	}
	_labelTypeErr := writeBuffer.WriteSerializable(ctx, m.GetLabelType())
	if popErr := writeBuffer.PopContext("labelType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelType")
	}
	if _labelTypeErr != nil {
		return errors.Wrap(_labelTypeErr, "Error serializing 'labelType' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField3 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField3
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("LightingLabelOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LightingLabelOptions")
	}
	return nil
}

func (m *_LightingLabelOptions) isLightingLabelOptions() bool {
	return true
}

func (m *_LightingLabelOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
