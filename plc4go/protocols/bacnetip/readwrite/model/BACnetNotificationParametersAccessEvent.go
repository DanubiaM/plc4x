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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersAccessEvent is the corresponding interface of BACnetNotificationParametersAccessEvent
type BACnetNotificationParametersAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() BACnetAccessEventTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetAccessEventTag returns AccessEventTag (property field)
	GetAccessEventTag() BACnetContextTagUnsignedInteger
	// GetAccessEventTime returns AccessEventTime (property field)
	GetAccessEventTime() BACnetTimeStampEnclosed
	// GetAccessCredential returns AccessCredential (property field)
	GetAccessCredential() BACnetDeviceObjectReferenceEnclosed
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactorTypeTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersAccessEventExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersAccessEvent.
// This is useful for switch cases.
type BACnetNotificationParametersAccessEventExactly interface {
	BACnetNotificationParametersAccessEvent
	isBACnetNotificationParametersAccessEvent() bool
}

// _BACnetNotificationParametersAccessEvent is the data-structure of this message
type _BACnetNotificationParametersAccessEvent struct {
	*_BACnetNotificationParameters
	InnerOpeningTag      BACnetOpeningTag
	AccessEvent          BACnetAccessEventTagged
	StatusFlags          BACnetStatusFlagsTagged
	AccessEventTag       BACnetContextTagUnsignedInteger
	AccessEventTime      BACnetTimeStampEnclosed
	AccessCredential     BACnetDeviceObjectReferenceEnclosed
	AuthenticationFactor BACnetAuthenticationFactorTypeTagged
	InnerClosingTag      BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersAccessEvent) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersAccessEvent) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersAccessEvent) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersAccessEvent) GetAccessEvent() BACnetAccessEventTagged {
	return m.AccessEvent
}

func (m *_BACnetNotificationParametersAccessEvent) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersAccessEvent) GetAccessEventTag() BACnetContextTagUnsignedInteger {
	return m.AccessEventTag
}

func (m *_BACnetNotificationParametersAccessEvent) GetAccessEventTime() BACnetTimeStampEnclosed {
	return m.AccessEventTime
}

func (m *_BACnetNotificationParametersAccessEvent) GetAccessCredential() BACnetDeviceObjectReferenceEnclosed {
	return m.AccessCredential
}

func (m *_BACnetNotificationParametersAccessEvent) GetAuthenticationFactor() BACnetAuthenticationFactorTypeTagged {
	return m.AuthenticationFactor
}

func (m *_BACnetNotificationParametersAccessEvent) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersAccessEvent factory function for _BACnetNotificationParametersAccessEvent
func NewBACnetNotificationParametersAccessEvent(innerOpeningTag BACnetOpeningTag, accessEvent BACnetAccessEventTagged, statusFlags BACnetStatusFlagsTagged, accessEventTag BACnetContextTagUnsignedInteger, accessEventTime BACnetTimeStampEnclosed, accessCredential BACnetDeviceObjectReferenceEnclosed, authenticationFactor BACnetAuthenticationFactorTypeTagged, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersAccessEvent {
	_result := &_BACnetNotificationParametersAccessEvent{
		InnerOpeningTag:               innerOpeningTag,
		AccessEvent:                   accessEvent,
		StatusFlags:                   statusFlags,
		AccessEventTag:                accessEventTag,
		AccessEventTime:               accessEventTime,
		AccessCredential:              accessCredential,
		AuthenticationFactor:          authenticationFactor,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersAccessEvent(structType interface{}) BACnetNotificationParametersAccessEvent {
	if casted, ok := structType.(BACnetNotificationParametersAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersAccessEvent) GetTypeName() string {
	return "BACnetNotificationParametersAccessEvent"
}

func (m *_BACnetNotificationParametersAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (accessEventTag)
	lengthInBits += m.AccessEventTag.GetLengthInBits(ctx)

	// Simple field (accessEventTime)
	lengthInBits += m.AccessEventTime.GetLengthInBits(ctx)

	// Simple field (accessCredential)
	lengthInBits += m.AccessCredential.GetLengthInBits(ctx)

	// Optional Field (authenticationFactor)
	if m.AuthenticationFactor != nil {
		lengthInBits += m.AuthenticationFactor.GetLengthInBits(ctx)
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersAccessEventParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersAccessEvent, error) {
	return BACnetNotificationParametersAccessEventParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersAccessEventParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersAccessEvent")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (accessEvent)
	if pullErr := readBuffer.PullContext("accessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEvent")
	}
	_accessEvent, _accessEventErr := BACnetAccessEventTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _accessEventErr != nil {
		return nil, errors.Wrap(_accessEventErr, "Error parsing 'accessEvent' field of BACnetNotificationParametersAccessEvent")
	}
	accessEvent := _accessEvent.(BACnetAccessEventTagged)
	if closeErr := readBuffer.CloseContext("accessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEvent")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field of BACnetNotificationParametersAccessEvent")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (accessEventTag)
	if pullErr := readBuffer.PullContext("accessEventTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventTag")
	}
	_accessEventTag, _accessEventTagErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _accessEventTagErr != nil {
		return nil, errors.Wrap(_accessEventTagErr, "Error parsing 'accessEventTag' field of BACnetNotificationParametersAccessEvent")
	}
	accessEventTag := _accessEventTag.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("accessEventTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventTag")
	}

	// Simple Field (accessEventTime)
	if pullErr := readBuffer.PullContext("accessEventTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventTime")
	}
	_accessEventTime, _accessEventTimeErr := BACnetTimeStampEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(3)))
	if _accessEventTimeErr != nil {
		return nil, errors.Wrap(_accessEventTimeErr, "Error parsing 'accessEventTime' field of BACnetNotificationParametersAccessEvent")
	}
	accessEventTime := _accessEventTime.(BACnetTimeStampEnclosed)
	if closeErr := readBuffer.CloseContext("accessEventTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventTime")
	}

	// Simple Field (accessCredential)
	if pullErr := readBuffer.PullContext("accessCredential"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessCredential")
	}
	_accessCredential, _accessCredentialErr := BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(4)))
	if _accessCredentialErr != nil {
		return nil, errors.Wrap(_accessCredentialErr, "Error parsing 'accessCredential' field of BACnetNotificationParametersAccessEvent")
	}
	accessCredential := _accessCredential.(BACnetDeviceObjectReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("accessCredential"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessCredential")
	}

	// Optional Field (authenticationFactor) (Can be skipped, if a given expression evaluates to false)
	var authenticationFactor BACnetAuthenticationFactorTypeTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("authenticationFactor"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactor")
		}
		_val, _err := BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx, readBuffer, uint8(5), TagClass_CONTEXT_SPECIFIC_TAGS)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'authenticationFactor' field of BACnetNotificationParametersAccessEvent")
		default:
			authenticationFactor = _val.(BACnetAuthenticationFactorTypeTagged)
			if closeErr := readBuffer.CloseContext("authenticationFactor"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for authenticationFactor")
			}
		}
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersAccessEvent")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersAccessEvent{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag:      innerOpeningTag,
		AccessEvent:          accessEvent,
		StatusFlags:          statusFlags,
		AccessEventTag:       accessEventTag,
		AccessEventTime:      accessEventTime,
		AccessCredential:     accessCredential,
		AuthenticationFactor: authenticationFactor,
		InnerClosingTag:      innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersAccessEvent")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (accessEvent)
		if pushErr := writeBuffer.PushContext("accessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEvent")
		}
		_accessEventErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEvent())
		if popErr := writeBuffer.PopContext("accessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEvent")
		}
		if _accessEventErr != nil {
			return errors.Wrap(_accessEventErr, "Error serializing 'accessEvent' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (accessEventTag)
		if pushErr := writeBuffer.PushContext("accessEventTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventTag")
		}
		_accessEventTagErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEventTag())
		if popErr := writeBuffer.PopContext("accessEventTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventTag")
		}
		if _accessEventTagErr != nil {
			return errors.Wrap(_accessEventTagErr, "Error serializing 'accessEventTag' field")
		}

		// Simple Field (accessEventTime)
		if pushErr := writeBuffer.PushContext("accessEventTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventTime")
		}
		_accessEventTimeErr := writeBuffer.WriteSerializable(ctx, m.GetAccessEventTime())
		if popErr := writeBuffer.PopContext("accessEventTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventTime")
		}
		if _accessEventTimeErr != nil {
			return errors.Wrap(_accessEventTimeErr, "Error serializing 'accessEventTime' field")
		}

		// Simple Field (accessCredential)
		if pushErr := writeBuffer.PushContext("accessCredential"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessCredential")
		}
		_accessCredentialErr := writeBuffer.WriteSerializable(ctx, m.GetAccessCredential())
		if popErr := writeBuffer.PopContext("accessCredential"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessCredential")
		}
		if _accessCredentialErr != nil {
			return errors.Wrap(_accessCredentialErr, "Error serializing 'accessCredential' field")
		}

		// Optional Field (authenticationFactor) (Can be skipped, if the value is null)
		var authenticationFactor BACnetAuthenticationFactorTypeTagged = nil
		if m.GetAuthenticationFactor() != nil {
			if pushErr := writeBuffer.PushContext("authenticationFactor"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for authenticationFactor")
			}
			authenticationFactor = m.GetAuthenticationFactor()
			_authenticationFactorErr := writeBuffer.WriteSerializable(ctx, authenticationFactor)
			if popErr := writeBuffer.PopContext("authenticationFactor"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for authenticationFactor")
			}
			if _authenticationFactorErr != nil {
				return errors.Wrap(_authenticationFactorErr, "Error serializing 'authenticationFactor' field")
			}
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersAccessEvent) isBACnetNotificationParametersAccessEvent() bool {
	return true
}

func (m *_BACnetNotificationParametersAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
