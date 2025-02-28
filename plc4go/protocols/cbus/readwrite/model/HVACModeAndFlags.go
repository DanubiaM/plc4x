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

// HVACModeAndFlags is the corresponding interface of HVACModeAndFlags
type HVACModeAndFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAuxiliaryLevel returns AuxiliaryLevel (property field)
	GetAuxiliaryLevel() bool
	// GetGuard returns Guard (property field)
	GetGuard() bool
	// GetSetback returns Setback (property field)
	GetSetback() bool
	// GetLevel returns Level (property field)
	GetLevel() bool
	// GetMode returns Mode (property field)
	GetMode() HVACModeAndFlagsMode
	// GetIsAuxLevelUnused returns IsAuxLevelUnused (virtual field)
	GetIsAuxLevelUnused() bool
	// GetIsAuxLevelUsed returns IsAuxLevelUsed (virtual field)
	GetIsAuxLevelUsed() bool
	// GetIsGuardDisabled returns IsGuardDisabled (virtual field)
	GetIsGuardDisabled() bool
	// GetIsGuardEnabled returns IsGuardEnabled (virtual field)
	GetIsGuardEnabled() bool
	// GetIsSetbackDisabled returns IsSetbackDisabled (virtual field)
	GetIsSetbackDisabled() bool
	// GetIsSetbackEnabled returns IsSetbackEnabled (virtual field)
	GetIsSetbackEnabled() bool
	// GetIsLevelTemperature returns IsLevelTemperature (virtual field)
	GetIsLevelTemperature() bool
	// GetIsLevelRaw returns IsLevelRaw (virtual field)
	GetIsLevelRaw() bool
}

// HVACModeAndFlagsExactly can be used when we want exactly this type and not a type which fulfills HVACModeAndFlags.
// This is useful for switch cases.
type HVACModeAndFlagsExactly interface {
	HVACModeAndFlags
	isHVACModeAndFlags() bool
}

// _HVACModeAndFlags is the data-structure of this message
type _HVACModeAndFlags struct {
	AuxiliaryLevel bool
	Guard          bool
	Setback        bool
	Level          bool
	Mode           HVACModeAndFlagsMode
	// Reserved Fields
	reservedField0 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACModeAndFlags) GetAuxiliaryLevel() bool {
	return m.AuxiliaryLevel
}

func (m *_HVACModeAndFlags) GetGuard() bool {
	return m.Guard
}

func (m *_HVACModeAndFlags) GetSetback() bool {
	return m.Setback
}

func (m *_HVACModeAndFlags) GetLevel() bool {
	return m.Level
}

func (m *_HVACModeAndFlags) GetMode() HVACModeAndFlagsMode {
	return m.Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACModeAndFlags) GetIsAuxLevelUnused() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetAuxiliaryLevel()))
}

func (m *_HVACModeAndFlags) GetIsAuxLevelUsed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetAuxiliaryLevel())
}

func (m *_HVACModeAndFlags) GetIsGuardDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetGuard()))
}

func (m *_HVACModeAndFlags) GetIsGuardEnabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetGuard())
}

func (m *_HVACModeAndFlags) GetIsSetbackDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetSetback()))
}

func (m *_HVACModeAndFlags) GetIsSetbackEnabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetSetback())
}

func (m *_HVACModeAndFlags) GetIsLevelTemperature() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetLevel()))
}

func (m *_HVACModeAndFlags) GetIsLevelRaw() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACModeAndFlags factory function for _HVACModeAndFlags
func NewHVACModeAndFlags(auxiliaryLevel bool, guard bool, setback bool, level bool, mode HVACModeAndFlagsMode) *_HVACModeAndFlags {
	return &_HVACModeAndFlags{AuxiliaryLevel: auxiliaryLevel, Guard: guard, Setback: setback, Level: level, Mode: mode}
}

// Deprecated: use the interface for direct cast
func CastHVACModeAndFlags(structType interface{}) HVACModeAndFlags {
	if casted, ok := structType.(HVACModeAndFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACModeAndFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACModeAndFlags) GetTypeName() string {
	return "HVACModeAndFlags"
}

func (m *_HVACModeAndFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (auxiliaryLevel)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (guard)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (setback)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (level)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (mode)
	lengthInBits += 3

	return lengthInBits
}

func (m *_HVACModeAndFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACModeAndFlagsParse(theBytes []byte) (HVACModeAndFlags, error) {
	return HVACModeAndFlagsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func HVACModeAndFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACModeAndFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACModeAndFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of HVACModeAndFlags")
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

	// Simple Field (auxiliaryLevel)
	_auxiliaryLevel, _auxiliaryLevelErr := readBuffer.ReadBit("auxiliaryLevel")
	if _auxiliaryLevelErr != nil {
		return nil, errors.Wrap(_auxiliaryLevelErr, "Error parsing 'auxiliaryLevel' field of HVACModeAndFlags")
	}
	auxiliaryLevel := _auxiliaryLevel

	// Virtual field
	_isAuxLevelUnused := !(auxiliaryLevel)
	isAuxLevelUnused := bool(_isAuxLevelUnused)
	_ = isAuxLevelUnused

	// Virtual field
	_isAuxLevelUsed := auxiliaryLevel
	isAuxLevelUsed := bool(_isAuxLevelUsed)
	_ = isAuxLevelUsed

	// Simple Field (guard)
	_guard, _guardErr := readBuffer.ReadBit("guard")
	if _guardErr != nil {
		return nil, errors.Wrap(_guardErr, "Error parsing 'guard' field of HVACModeAndFlags")
	}
	guard := _guard

	// Virtual field
	_isGuardDisabled := !(guard)
	isGuardDisabled := bool(_isGuardDisabled)
	_ = isGuardDisabled

	// Virtual field
	_isGuardEnabled := guard
	isGuardEnabled := bool(_isGuardEnabled)
	_ = isGuardEnabled

	// Simple Field (setback)
	_setback, _setbackErr := readBuffer.ReadBit("setback")
	if _setbackErr != nil {
		return nil, errors.Wrap(_setbackErr, "Error parsing 'setback' field of HVACModeAndFlags")
	}
	setback := _setback

	// Virtual field
	_isSetbackDisabled := !(setback)
	isSetbackDisabled := bool(_isSetbackDisabled)
	_ = isSetbackDisabled

	// Virtual field
	_isSetbackEnabled := setback
	isSetbackEnabled := bool(_isSetbackEnabled)
	_ = isSetbackEnabled

	// Simple Field (level)
	_level, _levelErr := readBuffer.ReadBit("level")
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of HVACModeAndFlags")
	}
	level := _level

	// Virtual field
	_isLevelTemperature := !(level)
	isLevelTemperature := bool(_isLevelTemperature)
	_ = isLevelTemperature

	// Virtual field
	_isLevelRaw := level
	isLevelRaw := bool(_isLevelRaw)
	_ = isLevelRaw

	// Simple Field (mode)
	if pullErr := readBuffer.PullContext("mode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for mode")
	}
	_mode, _modeErr := HVACModeAndFlagsModeParseWithBuffer(ctx, readBuffer)
	if _modeErr != nil {
		return nil, errors.Wrap(_modeErr, "Error parsing 'mode' field of HVACModeAndFlags")
	}
	mode := _mode
	if closeErr := readBuffer.CloseContext("mode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for mode")
	}

	if closeErr := readBuffer.CloseContext("HVACModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACModeAndFlags")
	}

	// Create the instance
	return &_HVACModeAndFlags{
		AuxiliaryLevel: auxiliaryLevel,
		Guard:          guard,
		Setback:        setback,
		Level:          level,
		Mode:           mode,
		reservedField0: reservedField0,
	}, nil
}

func (m *_HVACModeAndFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACModeAndFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("HVACModeAndFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACModeAndFlags")
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

	// Simple Field (auxiliaryLevel)
	auxiliaryLevel := bool(m.GetAuxiliaryLevel())
	_auxiliaryLevelErr := writeBuffer.WriteBit("auxiliaryLevel", (auxiliaryLevel))
	if _auxiliaryLevelErr != nil {
		return errors.Wrap(_auxiliaryLevelErr, "Error serializing 'auxiliaryLevel' field")
	}
	// Virtual field
	if _isAuxLevelUnusedErr := writeBuffer.WriteVirtual(ctx, "isAuxLevelUnused", m.GetIsAuxLevelUnused()); _isAuxLevelUnusedErr != nil {
		return errors.Wrap(_isAuxLevelUnusedErr, "Error serializing 'isAuxLevelUnused' field")
	}
	// Virtual field
	if _isAuxLevelUsedErr := writeBuffer.WriteVirtual(ctx, "isAuxLevelUsed", m.GetIsAuxLevelUsed()); _isAuxLevelUsedErr != nil {
		return errors.Wrap(_isAuxLevelUsedErr, "Error serializing 'isAuxLevelUsed' field")
	}

	// Simple Field (guard)
	guard := bool(m.GetGuard())
	_guardErr := writeBuffer.WriteBit("guard", (guard))
	if _guardErr != nil {
		return errors.Wrap(_guardErr, "Error serializing 'guard' field")
	}
	// Virtual field
	if _isGuardDisabledErr := writeBuffer.WriteVirtual(ctx, "isGuardDisabled", m.GetIsGuardDisabled()); _isGuardDisabledErr != nil {
		return errors.Wrap(_isGuardDisabledErr, "Error serializing 'isGuardDisabled' field")
	}
	// Virtual field
	if _isGuardEnabledErr := writeBuffer.WriteVirtual(ctx, "isGuardEnabled", m.GetIsGuardEnabled()); _isGuardEnabledErr != nil {
		return errors.Wrap(_isGuardEnabledErr, "Error serializing 'isGuardEnabled' field")
	}

	// Simple Field (setback)
	setback := bool(m.GetSetback())
	_setbackErr := writeBuffer.WriteBit("setback", (setback))
	if _setbackErr != nil {
		return errors.Wrap(_setbackErr, "Error serializing 'setback' field")
	}
	// Virtual field
	if _isSetbackDisabledErr := writeBuffer.WriteVirtual(ctx, "isSetbackDisabled", m.GetIsSetbackDisabled()); _isSetbackDisabledErr != nil {
		return errors.Wrap(_isSetbackDisabledErr, "Error serializing 'isSetbackDisabled' field")
	}
	// Virtual field
	if _isSetbackEnabledErr := writeBuffer.WriteVirtual(ctx, "isSetbackEnabled", m.GetIsSetbackEnabled()); _isSetbackEnabledErr != nil {
		return errors.Wrap(_isSetbackEnabledErr, "Error serializing 'isSetbackEnabled' field")
	}

	// Simple Field (level)
	level := bool(m.GetLevel())
	_levelErr := writeBuffer.WriteBit("level", (level))
	if _levelErr != nil {
		return errors.Wrap(_levelErr, "Error serializing 'level' field")
	}
	// Virtual field
	if _isLevelTemperatureErr := writeBuffer.WriteVirtual(ctx, "isLevelTemperature", m.GetIsLevelTemperature()); _isLevelTemperatureErr != nil {
		return errors.Wrap(_isLevelTemperatureErr, "Error serializing 'isLevelTemperature' field")
	}
	// Virtual field
	if _isLevelRawErr := writeBuffer.WriteVirtual(ctx, "isLevelRaw", m.GetIsLevelRaw()); _isLevelRawErr != nil {
		return errors.Wrap(_isLevelRawErr, "Error serializing 'isLevelRaw' field")
	}

	// Simple Field (mode)
	if pushErr := writeBuffer.PushContext("mode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for mode")
	}
	_modeErr := writeBuffer.WriteSerializable(ctx, m.GetMode())
	if popErr := writeBuffer.PopContext("mode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for mode")
	}
	if _modeErr != nil {
		return errors.Wrap(_modeErr, "Error serializing 'mode' field")
	}

	if popErr := writeBuffer.PopContext("HVACModeAndFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACModeAndFlags")
	}
	return nil
}

func (m *_HVACModeAndFlags) isHVACModeAndFlags() bool {
	return true
}

func (m *_HVACModeAndFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
