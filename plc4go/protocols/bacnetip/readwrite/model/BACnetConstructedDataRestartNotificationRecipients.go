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

// BACnetConstructedDataRestartNotificationRecipients is the corresponding interface of BACnetConstructedDataRestartNotificationRecipients
type BACnetConstructedDataRestartNotificationRecipients interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRestartNotificationRecipients returns RestartNotificationRecipients (property field)
	GetRestartNotificationRecipients() []BACnetRecipient
}

// BACnetConstructedDataRestartNotificationRecipientsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRestartNotificationRecipients.
// This is useful for switch cases.
type BACnetConstructedDataRestartNotificationRecipientsExactly interface {
	BACnetConstructedDataRestartNotificationRecipients
	isBACnetConstructedDataRestartNotificationRecipients() bool
}

// _BACnetConstructedDataRestartNotificationRecipients is the data-structure of this message
type _BACnetConstructedDataRestartNotificationRecipients struct {
	*_BACnetConstructedData
	RestartNotificationRecipients []BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTART_NOTIFICATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRestartNotificationRecipients) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetRestartNotificationRecipients() []BACnetRecipient {
	return m.RestartNotificationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRestartNotificationRecipients factory function for _BACnetConstructedDataRestartNotificationRecipients
func NewBACnetConstructedDataRestartNotificationRecipients(restartNotificationRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRestartNotificationRecipients {
	_result := &_BACnetConstructedDataRestartNotificationRecipients{
		RestartNotificationRecipients: restartNotificationRecipients,
		_BACnetConstructedData:        NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRestartNotificationRecipients(structType interface{}) BACnetConstructedDataRestartNotificationRecipients {
	if casted, ok := structType.(BACnetConstructedDataRestartNotificationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestartNotificationRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetTypeName() string {
	return "BACnetConstructedDataRestartNotificationRecipients"
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.RestartNotificationRecipients) > 0 {
		for _, element := range m.RestartNotificationRecipients {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataRestartNotificationRecipientsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRestartNotificationRecipients, error) {
	return BACnetConstructedDataRestartNotificationRecipientsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataRestartNotificationRecipientsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRestartNotificationRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestartNotificationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRestartNotificationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (restartNotificationRecipients)
	if pullErr := readBuffer.PullContext("restartNotificationRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for restartNotificationRecipients")
	}
	// Terminated array
	var restartNotificationRecipients []BACnetRecipient
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'restartNotificationRecipients' field of BACnetConstructedDataRestartNotificationRecipients")
			}
			restartNotificationRecipients = append(restartNotificationRecipients, _item.(BACnetRecipient))
		}
	}
	if closeErr := readBuffer.CloseContext("restartNotificationRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for restartNotificationRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestartNotificationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRestartNotificationRecipients")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRestartNotificationRecipients{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RestartNotificationRecipients: restartNotificationRecipients,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestartNotificationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRestartNotificationRecipients")
		}

		// Array Field (restartNotificationRecipients)
		if pushErr := writeBuffer.PushContext("restartNotificationRecipients", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for restartNotificationRecipients")
		}
		for _curItem, _element := range m.GetRestartNotificationRecipients() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetRestartNotificationRecipients()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'restartNotificationRecipients' field")
			}
		}
		if popErr := writeBuffer.PopContext("restartNotificationRecipients", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for restartNotificationRecipients")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestartNotificationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRestartNotificationRecipients")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) isBACnetConstructedDataRestartNotificationRecipients() bool {
	return true
}

func (m *_BACnetConstructedDataRestartNotificationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
