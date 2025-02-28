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

// BACnetConstructedDataFileRecordCount is the corresponding interface of BACnetConstructedDataFileRecordCount
type BACnetConstructedDataFileRecordCount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRecordCount returns RecordCount (property field)
	GetRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataFileRecordCountExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFileRecordCount.
// This is useful for switch cases.
type BACnetConstructedDataFileRecordCountExactly interface {
	BACnetConstructedDataFileRecordCount
	isBACnetConstructedDataFileRecordCount() bool
}

// _BACnetConstructedDataFileRecordCount is the data-structure of this message
type _BACnetConstructedDataFileRecordCount struct {
	*_BACnetConstructedData
	RecordCount BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_FILE
}

func (m *_BACnetConstructedDataFileRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFileRecordCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFileRecordCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFileRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFileRecordCount factory function for _BACnetConstructedDataFileRecordCount
func NewBACnetConstructedDataFileRecordCount(recordCount BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFileRecordCount {
	_result := &_BACnetConstructedDataFileRecordCount{
		RecordCount:            recordCount,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFileRecordCount(structType interface{}) BACnetConstructedDataFileRecordCount {
	if casted, ok := structType.(BACnetConstructedDataFileRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFileRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFileRecordCount) GetTypeName() string {
	return "BACnetConstructedDataFileRecordCount"
}

func (m *_BACnetConstructedDataFileRecordCount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (recordCount)
	lengthInBits += m.RecordCount.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFileRecordCount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFileRecordCountParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFileRecordCount, error) {
	return BACnetConstructedDataFileRecordCountParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFileRecordCountParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFileRecordCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFileRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFileRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recordCount)
	if pullErr := readBuffer.PullContext("recordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recordCount")
	}
	_recordCount, _recordCountErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _recordCountErr != nil {
		return nil, errors.Wrap(_recordCountErr, "Error parsing 'recordCount' field of BACnetConstructedDataFileRecordCount")
	}
	recordCount := _recordCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("recordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recordCount")
	}

	// Virtual field
	_actualValue := recordCount
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFileRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFileRecordCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFileRecordCount{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RecordCount: recordCount,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFileRecordCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFileRecordCount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFileRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFileRecordCount")
		}

		// Simple Field (recordCount)
		if pushErr := writeBuffer.PushContext("recordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for recordCount")
		}
		_recordCountErr := writeBuffer.WriteSerializable(ctx, m.GetRecordCount())
		if popErr := writeBuffer.PopContext("recordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for recordCount")
		}
		if _recordCountErr != nil {
			return errors.Wrap(_recordCountErr, "Error serializing 'recordCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFileRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFileRecordCount")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFileRecordCount) isBACnetConstructedDataFileRecordCount() bool {
	return true
}

func (m *_BACnetConstructedDataFileRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
