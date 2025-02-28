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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsTableSizes is the corresponding interface of AdsTableSizes
type AdsTableSizes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSymbolCount returns SymbolCount (property field)
	GetSymbolCount() uint32
	// GetSymbolLength returns SymbolLength (property field)
	GetSymbolLength() uint32
	// GetDataTypeCount returns DataTypeCount (property field)
	GetDataTypeCount() uint32
	// GetDataTypeLength returns DataTypeLength (property field)
	GetDataTypeLength() uint32
	// GetExtraCount returns ExtraCount (property field)
	GetExtraCount() uint32
	// GetExtraLength returns ExtraLength (property field)
	GetExtraLength() uint32
}

// AdsTableSizesExactly can be used when we want exactly this type and not a type which fulfills AdsTableSizes.
// This is useful for switch cases.
type AdsTableSizesExactly interface {
	AdsTableSizes
	isAdsTableSizes() bool
}

// _AdsTableSizes is the data-structure of this message
type _AdsTableSizes struct {
	SymbolCount    uint32
	SymbolLength   uint32
	DataTypeCount  uint32
	DataTypeLength uint32
	ExtraCount     uint32
	ExtraLength    uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsTableSizes) GetSymbolCount() uint32 {
	return m.SymbolCount
}

func (m *_AdsTableSizes) GetSymbolLength() uint32 {
	return m.SymbolLength
}

func (m *_AdsTableSizes) GetDataTypeCount() uint32 {
	return m.DataTypeCount
}

func (m *_AdsTableSizes) GetDataTypeLength() uint32 {
	return m.DataTypeLength
}

func (m *_AdsTableSizes) GetExtraCount() uint32 {
	return m.ExtraCount
}

func (m *_AdsTableSizes) GetExtraLength() uint32 {
	return m.ExtraLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsTableSizes factory function for _AdsTableSizes
func NewAdsTableSizes(symbolCount uint32, symbolLength uint32, dataTypeCount uint32, dataTypeLength uint32, extraCount uint32, extraLength uint32) *_AdsTableSizes {
	return &_AdsTableSizes{SymbolCount: symbolCount, SymbolLength: symbolLength, DataTypeCount: dataTypeCount, DataTypeLength: dataTypeLength, ExtraCount: extraCount, ExtraLength: extraLength}
}

// Deprecated: use the interface for direct cast
func CastAdsTableSizes(structType interface{}) AdsTableSizes {
	if casted, ok := structType.(AdsTableSizes); ok {
		return casted
	}
	if casted, ok := structType.(*AdsTableSizes); ok {
		return *casted
	}
	return nil
}

func (m *_AdsTableSizes) GetTypeName() string {
	return "AdsTableSizes"
}

func (m *_AdsTableSizes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (symbolCount)
	lengthInBits += 32

	// Simple field (symbolLength)
	lengthInBits += 32

	// Simple field (dataTypeCount)
	lengthInBits += 32

	// Simple field (dataTypeLength)
	lengthInBits += 32

	// Simple field (extraCount)
	lengthInBits += 32

	// Simple field (extraLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsTableSizes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsTableSizesParse(theBytes []byte) (AdsTableSizes, error) {
	return AdsTableSizesParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AdsTableSizesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsTableSizes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsTableSizes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsTableSizes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (symbolCount)
	_symbolCount, _symbolCountErr := readBuffer.ReadUint32("symbolCount", 32)
	if _symbolCountErr != nil {
		return nil, errors.Wrap(_symbolCountErr, "Error parsing 'symbolCount' field of AdsTableSizes")
	}
	symbolCount := _symbolCount

	// Simple Field (symbolLength)
	_symbolLength, _symbolLengthErr := readBuffer.ReadUint32("symbolLength", 32)
	if _symbolLengthErr != nil {
		return nil, errors.Wrap(_symbolLengthErr, "Error parsing 'symbolLength' field of AdsTableSizes")
	}
	symbolLength := _symbolLength

	// Simple Field (dataTypeCount)
	_dataTypeCount, _dataTypeCountErr := readBuffer.ReadUint32("dataTypeCount", 32)
	if _dataTypeCountErr != nil {
		return nil, errors.Wrap(_dataTypeCountErr, "Error parsing 'dataTypeCount' field of AdsTableSizes")
	}
	dataTypeCount := _dataTypeCount

	// Simple Field (dataTypeLength)
	_dataTypeLength, _dataTypeLengthErr := readBuffer.ReadUint32("dataTypeLength", 32)
	if _dataTypeLengthErr != nil {
		return nil, errors.Wrap(_dataTypeLengthErr, "Error parsing 'dataTypeLength' field of AdsTableSizes")
	}
	dataTypeLength := _dataTypeLength

	// Simple Field (extraCount)
	_extraCount, _extraCountErr := readBuffer.ReadUint32("extraCount", 32)
	if _extraCountErr != nil {
		return nil, errors.Wrap(_extraCountErr, "Error parsing 'extraCount' field of AdsTableSizes")
	}
	extraCount := _extraCount

	// Simple Field (extraLength)
	_extraLength, _extraLengthErr := readBuffer.ReadUint32("extraLength", 32)
	if _extraLengthErr != nil {
		return nil, errors.Wrap(_extraLengthErr, "Error parsing 'extraLength' field of AdsTableSizes")
	}
	extraLength := _extraLength

	if closeErr := readBuffer.CloseContext("AdsTableSizes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsTableSizes")
	}

	// Create the instance
	return &_AdsTableSizes{
		SymbolCount:    symbolCount,
		SymbolLength:   symbolLength,
		DataTypeCount:  dataTypeCount,
		DataTypeLength: dataTypeLength,
		ExtraCount:     extraCount,
		ExtraLength:    extraLength,
	}, nil
}

func (m *_AdsTableSizes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsTableSizes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsTableSizes"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsTableSizes")
	}

	// Simple Field (symbolCount)
	symbolCount := uint32(m.GetSymbolCount())
	_symbolCountErr := writeBuffer.WriteUint32("symbolCount", 32, (symbolCount))
	if _symbolCountErr != nil {
		return errors.Wrap(_symbolCountErr, "Error serializing 'symbolCount' field")
	}

	// Simple Field (symbolLength)
	symbolLength := uint32(m.GetSymbolLength())
	_symbolLengthErr := writeBuffer.WriteUint32("symbolLength", 32, (symbolLength))
	if _symbolLengthErr != nil {
		return errors.Wrap(_symbolLengthErr, "Error serializing 'symbolLength' field")
	}

	// Simple Field (dataTypeCount)
	dataTypeCount := uint32(m.GetDataTypeCount())
	_dataTypeCountErr := writeBuffer.WriteUint32("dataTypeCount", 32, (dataTypeCount))
	if _dataTypeCountErr != nil {
		return errors.Wrap(_dataTypeCountErr, "Error serializing 'dataTypeCount' field")
	}

	// Simple Field (dataTypeLength)
	dataTypeLength := uint32(m.GetDataTypeLength())
	_dataTypeLengthErr := writeBuffer.WriteUint32("dataTypeLength", 32, (dataTypeLength))
	if _dataTypeLengthErr != nil {
		return errors.Wrap(_dataTypeLengthErr, "Error serializing 'dataTypeLength' field")
	}

	// Simple Field (extraCount)
	extraCount := uint32(m.GetExtraCount())
	_extraCountErr := writeBuffer.WriteUint32("extraCount", 32, (extraCount))
	if _extraCountErr != nil {
		return errors.Wrap(_extraCountErr, "Error serializing 'extraCount' field")
	}

	// Simple Field (extraLength)
	extraLength := uint32(m.GetExtraLength())
	_extraLengthErr := writeBuffer.WriteUint32("extraLength", 32, (extraLength))
	if _extraLengthErr != nil {
		return errors.Wrap(_extraLengthErr, "Error serializing 'extraLength' field")
	}

	if popErr := writeBuffer.PopContext("AdsTableSizes"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsTableSizes")
	}
	return nil
}

func (m *_AdsTableSizes) isAdsTableSizes() bool {
	return true
}

func (m *_AdsTableSizes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
