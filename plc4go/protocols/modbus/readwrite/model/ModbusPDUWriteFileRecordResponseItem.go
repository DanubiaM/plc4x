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

// ModbusPDUWriteFileRecordResponseItem is the corresponding interface of ModbusPDUWriteFileRecordResponseItem
type ModbusPDUWriteFileRecordResponseItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordData returns RecordData (property field)
	GetRecordData() []byte
}

// ModbusPDUWriteFileRecordResponseItemExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteFileRecordResponseItem.
// This is useful for switch cases.
type ModbusPDUWriteFileRecordResponseItemExactly interface {
	ModbusPDUWriteFileRecordResponseItem
	isModbusPDUWriteFileRecordResponseItem() bool
}

// _ModbusPDUWriteFileRecordResponseItem is the data-structure of this message
type _ModbusPDUWriteFileRecordResponseItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordData    []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteFileRecordResponseItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetRecordData() []byte {
	return m.RecordData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteFileRecordResponseItem factory function for _ModbusPDUWriteFileRecordResponseItem
func NewModbusPDUWriteFileRecordResponseItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) *_ModbusPDUWriteFileRecordResponseItem {
	return &_ModbusPDUWriteFileRecordResponseItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordData: recordData}
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteFileRecordResponseItem(structType interface{}) ModbusPDUWriteFileRecordResponseItem {
	if casted, ok := structType.(ModbusPDUWriteFileRecordResponseItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordResponseItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponseItem"
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Implicit Field (recordLength)
	lengthInBits += 16

	// Array field
	if len(m.RecordData) > 0 {
		lengthInBits += 8 * uint16(len(m.RecordData))
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteFileRecordResponseItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteFileRecordResponseItemParse(theBytes []byte) (ModbusPDUWriteFileRecordResponseItem, error) {
	return ModbusPDUWriteFileRecordResponseItemParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUWriteFileRecordResponseItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUWriteFileRecordResponseItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponseItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteFileRecordResponseItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceType)
	_referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field of ModbusPDUWriteFileRecordResponseItem")
	}
	referenceType := _referenceType

	// Simple Field (fileNumber)
	_fileNumber, _fileNumberErr := readBuffer.ReadUint16("fileNumber", 16)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field of ModbusPDUWriteFileRecordResponseItem")
	}
	fileNumber := _fileNumber

	// Simple Field (recordNumber)
	_recordNumber, _recordNumberErr := readBuffer.ReadUint16("recordNumber", 16)
	if _recordNumberErr != nil {
		return nil, errors.Wrap(_recordNumberErr, "Error parsing 'recordNumber' field of ModbusPDUWriteFileRecordResponseItem")
	}
	recordNumber := _recordNumber

	// Implicit Field (recordLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	recordLength, _recordLengthErr := readBuffer.ReadUint16("recordLength", 16)
	_ = recordLength
	if _recordLengthErr != nil {
		return nil, errors.Wrap(_recordLengthErr, "Error parsing 'recordLength' field of ModbusPDUWriteFileRecordResponseItem")
	}
	// Byte Array field (recordData)
	numberOfBytesrecordData := int(recordLength)
	recordData, _readArrayErr := readBuffer.ReadByteArray("recordData", numberOfBytesrecordData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'recordData' field of ModbusPDUWriteFileRecordResponseItem")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponseItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteFileRecordResponseItem")
	}

	// Create the instance
	return &_ModbusPDUWriteFileRecordResponseItem{
		ReferenceType: referenceType,
		FileNumber:    fileNumber,
		RecordNumber:  recordNumber,
		RecordData:    recordData,
	}, nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponseItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteFileRecordResponseItem")
	}

	// Simple Field (referenceType)
	referenceType := uint8(m.GetReferenceType())
	_referenceTypeErr := writeBuffer.WriteUint8("referenceType", 8, (referenceType))
	if _referenceTypeErr != nil {
		return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
	}

	// Simple Field (fileNumber)
	fileNumber := uint16(m.GetFileNumber())
	_fileNumberErr := writeBuffer.WriteUint16("fileNumber", 16, (fileNumber))
	if _fileNumberErr != nil {
		return errors.Wrap(_fileNumberErr, "Error serializing 'fileNumber' field")
	}

	// Simple Field (recordNumber)
	recordNumber := uint16(m.GetRecordNumber())
	_recordNumberErr := writeBuffer.WriteUint16("recordNumber", 16, (recordNumber))
	if _recordNumberErr != nil {
		return errors.Wrap(_recordNumberErr, "Error serializing 'recordNumber' field")
	}

	// Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	recordLength := uint16(uint16(uint16(len(m.GetRecordData()))) / uint16(uint16(2)))
	_recordLengthErr := writeBuffer.WriteUint16("recordLength", 16, (recordLength))
	if _recordLengthErr != nil {
		return errors.Wrap(_recordLengthErr, "Error serializing 'recordLength' field")
	}

	// Array Field (recordData)
	// Byte Array field (recordData)
	if err := writeBuffer.WriteByteArray("recordData", m.GetRecordData()); err != nil {
		return errors.Wrap(err, "Error serializing 'recordData' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponseItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUWriteFileRecordResponseItem")
	}
	return nil
}

func (m *_ModbusPDUWriteFileRecordResponseItem) isModbusPDUWriteFileRecordResponseItem() bool {
	return true
}

func (m *_ModbusPDUWriteFileRecordResponseItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
