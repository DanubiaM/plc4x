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

// MeteringDataMeasureGas is the corresponding interface of MeteringDataMeasureGas
type MeteringDataMeasureGas interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MeteringData
}

// MeteringDataMeasureGasExactly can be used when we want exactly this type and not a type which fulfills MeteringDataMeasureGas.
// This is useful for switch cases.
type MeteringDataMeasureGasExactly interface {
	MeteringDataMeasureGas
	isMeteringDataMeasureGas() bool
}

// _MeteringDataMeasureGas is the data-structure of this message
type _MeteringDataMeasureGas struct {
	*_MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataMeasureGas) InitializeParent(parent MeteringData, commandTypeContainer MeteringCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataMeasureGas) GetParent() MeteringData {
	return m._MeteringData
}

// NewMeteringDataMeasureGas factory function for _MeteringDataMeasureGas
func NewMeteringDataMeasureGas(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringDataMeasureGas {
	_result := &_MeteringDataMeasureGas{
		_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureGas(structType interface{}) MeteringDataMeasureGas {
	if casted, ok := structType.(MeteringDataMeasureGas); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureGas); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureGas) GetTypeName() string {
	return "MeteringDataMeasureGas"
}

func (m *_MeteringDataMeasureGas) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MeteringDataMeasureGas) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataMeasureGasParse(theBytes []byte) (MeteringDataMeasureGas, error) {
	return MeteringDataMeasureGasParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataMeasureGasParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureGas, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureGas"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureGas")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureGas"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureGas")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataMeasureGas{
		_MeteringData: &_MeteringData{},
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataMeasureGas) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureGas) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureGas"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureGas")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureGas"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureGas")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MeteringDataMeasureGas) isMeteringDataMeasureGas() bool {
	return true
}

func (m *_MeteringDataMeasureGas) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
