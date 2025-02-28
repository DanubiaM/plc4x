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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataUTCTimeSynchronizationRecipients is the corresponding interface of BACnetConstructedDataUTCTimeSynchronizationRecipients
type BACnetConstructedDataUTCTimeSynchronizationRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUtcTimeSynchronizationRecipients returns UtcTimeSynchronizationRecipients (property field)
	GetUtcTimeSynchronizationRecipients() []BACnetRecipient
}

// BACnetConstructedDataUTCTimeSynchronizationRecipientsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUTCTimeSynchronizationRecipients.
// This is useful for switch cases.
type BACnetConstructedDataUTCTimeSynchronizationRecipientsExactly interface {
	BACnetConstructedDataUTCTimeSynchronizationRecipients
	isBACnetConstructedDataUTCTimeSynchronizationRecipients() bool
}

// _BACnetConstructedDataUTCTimeSynchronizationRecipients is the data-structure of this message
type _BACnetConstructedDataUTCTimeSynchronizationRecipients struct {
	*_BACnetConstructedData
	UtcTimeSynchronizationRecipients []BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UTC_TIME_SYNCHRONIZATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetUtcTimeSynchronizationRecipients() []BACnetRecipient {
	return m.UtcTimeSynchronizationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUTCTimeSynchronizationRecipients factory function for _BACnetConstructedDataUTCTimeSynchronizationRecipients
func NewBACnetConstructedDataUTCTimeSynchronizationRecipients(utcTimeSynchronizationRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUTCTimeSynchronizationRecipients {
	_result := &_BACnetConstructedDataUTCTimeSynchronizationRecipients{
		UtcTimeSynchronizationRecipients: utcTimeSynchronizationRecipients,
		_BACnetConstructedData:           NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUTCTimeSynchronizationRecipients(structType interface{}) BACnetConstructedDataUTCTimeSynchronizationRecipients {
	if casted, ok := structType.(BACnetConstructedDataUTCTimeSynchronizationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUTCTimeSynchronizationRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetTypeName() string {
	return "BACnetConstructedDataUTCTimeSynchronizationRecipients"
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.UtcTimeSynchronizationRecipients) > 0 {
		for _, element := range m.UtcTimeSynchronizationRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataUTCTimeSynchronizationRecipientsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUTCTimeSynchronizationRecipients, error) {
	return BACnetConstructedDataUTCTimeSynchronizationRecipientsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataUTCTimeSynchronizationRecipientsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUTCTimeSynchronizationRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUTCTimeSynchronizationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (utcTimeSynchronizationRecipients)
	if pullErr := readBuffer.PullContext("utcTimeSynchronizationRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for utcTimeSynchronizationRecipients")
	}
	// Terminated array
	var utcTimeSynchronizationRecipients []BACnetRecipient
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'utcTimeSynchronizationRecipients' field of BACnetConstructedDataUTCTimeSynchronizationRecipients")
			}
			utcTimeSynchronizationRecipients = append(utcTimeSynchronizationRecipients, _item.(BACnetRecipient))
		}
	}
	if closeErr := readBuffer.CloseContext("utcTimeSynchronizationRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for utcTimeSynchronizationRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUTCTimeSynchronizationRecipients")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUTCTimeSynchronizationRecipients{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UtcTimeSynchronizationRecipients: utcTimeSynchronizationRecipients,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUTCTimeSynchronizationRecipients")
		}

		// Array Field (utcTimeSynchronizationRecipients)
		if pushErr := writeBuffer.PushContext("utcTimeSynchronizationRecipients", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for utcTimeSynchronizationRecipients")
		}
		for _curItem, _element := range m.GetUtcTimeSynchronizationRecipients() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetUtcTimeSynchronizationRecipients()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'utcTimeSynchronizationRecipients' field")
			}
		}
		if popErr := writeBuffer.PopContext("utcTimeSynchronizationRecipients", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for utcTimeSynchronizationRecipients")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUTCTimeSynchronizationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUTCTimeSynchronizationRecipients")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) isBACnetConstructedDataUTCTimeSynchronizationRecipients() bool {
	return true
}

func (m *_BACnetConstructedDataUTCTimeSynchronizationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
