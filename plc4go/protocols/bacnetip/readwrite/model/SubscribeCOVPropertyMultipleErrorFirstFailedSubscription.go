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

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the corresponding interface of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionExactly can be used when we want exactly this type and not a type which fulfills SubscribeCOVPropertyMultipleErrorFirstFailedSubscription.
// This is useful for switch cases.
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionExactly interface {
	SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
	isSubscribeCOVPropertyMultipleErrorFirstFailedSubscription() bool
}

// _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the data-structure of this message
type _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription struct {
	OpeningTag                 BACnetOpeningTag
	MonitoredObjectIdentifier  BACnetContextTagObjectIdentifier
	MonitoredPropertyReference BACnetPropertyReferenceEnclosed
	ErrorType                  ErrorEnclosed
	ClosingTag                 BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription factory function for _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
func NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(openingTag BACnetOpeningTag, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, monitoredPropertyReference BACnetPropertyReferenceEnclosed, errorType ErrorEnclosed, closingTag BACnetClosingTag, tagNumber uint8) *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{OpeningTag: openingTag, MonitoredObjectIdentifier: monitoredObjectIdentifier, MonitoredPropertyReference: monitoredPropertyReference, ErrorType: errorType, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(structType interface{}) SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits(ctx)

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(theBytes []byte, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	return SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	monitoredObjectIdentifier := _monitoredObjectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredObjectIdentifier")
	}

	// Simple Field (monitoredPropertyReference)
	if pullErr := readBuffer.PullContext("monitoredPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredPropertyReference")
	}
	_monitoredPropertyReference, _monitoredPropertyReferenceErr := BACnetPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)))
	if _monitoredPropertyReferenceErr != nil {
		return nil, errors.Wrap(_monitoredPropertyReferenceErr, "Error parsing 'monitoredPropertyReference' field of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	monitoredPropertyReference := _monitoredPropertyReference.(BACnetPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("monitoredPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredPropertyReference")
	}

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
	_errorType, _errorTypeErr := ErrorEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	// Create the instance
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{
		TagNumber:                  tagNumber,
		OpeningTag:                 openingTag,
		MonitoredObjectIdentifier:  monitoredObjectIdentifier,
		MonitoredPropertyReference: monitoredPropertyReference,
		ErrorType:                  errorType,
		ClosingTag:                 closingTag,
	}, nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoredObjectIdentifier())
	if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
	}
	if _monitoredObjectIdentifierErr != nil {
		return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
	}

	// Simple Field (monitoredPropertyReference)
	if pushErr := writeBuffer.PushContext("monitoredPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredPropertyReference")
	}
	_monitoredPropertyReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoredPropertyReference())
	if popErr := writeBuffer.PopContext("monitoredPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredPropertyReference")
	}
	if _monitoredPropertyReferenceErr != nil {
		return errors.Wrap(_monitoredPropertyReferenceErr, "Error serializing 'monitoredPropertyReference' field")
	}

	// Simple Field (errorType)
	if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for errorType")
	}
	_errorTypeErr := writeBuffer.WriteSerializable(ctx, m.GetErrorType())
	if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for errorType")
	}
	if _errorTypeErr != nil {
		return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	return nil
}

////
// Arguments Getter

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) isSubscribeCOVPropertyMultipleErrorFirstFailedSubscription() bool {
	return true
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
