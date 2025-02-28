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

// BACnetConstructedDataLastNotifyRecord is the corresponding interface of BACnetConstructedDataLastNotifyRecord
type BACnetConstructedDataLastNotifyRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastNotifyRecord returns LastNotifyRecord (property field)
	GetLastNotifyRecord() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLastNotifyRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastNotifyRecord.
// This is useful for switch cases.
type BACnetConstructedDataLastNotifyRecordExactly interface {
	BACnetConstructedDataLastNotifyRecord
	isBACnetConstructedDataLastNotifyRecord() bool
}

// _BACnetConstructedDataLastNotifyRecord is the data-structure of this message
type _BACnetConstructedDataLastNotifyRecord struct {
	*_BACnetConstructedData
	LastNotifyRecord BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_NOTIFY_RECORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetLastNotifyRecord() BACnetApplicationTagUnsignedInteger {
	return m.LastNotifyRecord
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastNotifyRecord) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLastNotifyRecord())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastNotifyRecord factory function for _BACnetConstructedDataLastNotifyRecord
func NewBACnetConstructedDataLastNotifyRecord(lastNotifyRecord BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastNotifyRecord {
	_result := &_BACnetConstructedDataLastNotifyRecord{
		LastNotifyRecord:       lastNotifyRecord,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastNotifyRecord(structType interface{}) BACnetConstructedDataLastNotifyRecord {
	if casted, ok := structType.(BACnetConstructedDataLastNotifyRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastNotifyRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetTypeName() string {
	return "BACnetConstructedDataLastNotifyRecord"
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastNotifyRecord)
	lengthInBits += m.LastNotifyRecord.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastNotifyRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastNotifyRecordParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastNotifyRecord, error) {
	return BACnetConstructedDataLastNotifyRecordParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastNotifyRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastNotifyRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastNotifyRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastNotifyRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastNotifyRecord)
	if pullErr := readBuffer.PullContext("lastNotifyRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastNotifyRecord")
	}
	_lastNotifyRecord, _lastNotifyRecordErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _lastNotifyRecordErr != nil {
		return nil, errors.Wrap(_lastNotifyRecordErr, "Error parsing 'lastNotifyRecord' field of BACnetConstructedDataLastNotifyRecord")
	}
	lastNotifyRecord := _lastNotifyRecord.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lastNotifyRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastNotifyRecord")
	}

	// Virtual field
	_actualValue := lastNotifyRecord
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastNotifyRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastNotifyRecord")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastNotifyRecord{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastNotifyRecord: lastNotifyRecord,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastNotifyRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastNotifyRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastNotifyRecord")
		}

		// Simple Field (lastNotifyRecord)
		if pushErr := writeBuffer.PushContext("lastNotifyRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastNotifyRecord")
		}
		_lastNotifyRecordErr := writeBuffer.WriteSerializable(ctx, m.GetLastNotifyRecord())
		if popErr := writeBuffer.PopContext("lastNotifyRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastNotifyRecord")
		}
		if _lastNotifyRecordErr != nil {
			return errors.Wrap(_lastNotifyRecordErr, "Error serializing 'lastNotifyRecord' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastNotifyRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastNotifyRecord")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastNotifyRecord) isBACnetConstructedDataLastNotifyRecord() bool {
	return true
}

func (m *_BACnetConstructedDataLastNotifyRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
