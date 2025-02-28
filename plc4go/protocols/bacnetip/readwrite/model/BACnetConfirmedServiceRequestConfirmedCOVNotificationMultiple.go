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

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() ListOfCovNotificationsList
}

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleExactly interface {
	BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
	isBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple() bool
}

// _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple struct {
	*_BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	TimeRemaining               BACnetContextTagUnsignedInteger
	Timestamp                   BACnetTimeStampEnclosed
	ListOfCovNotifications      ListOfCovNotificationsList
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetListOfCovNotifications() ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple factory function for _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, timestamp BACnetTimeStampEnclosed, listOfCovNotifications ListOfCovNotificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	_result := &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		SubscriberProcessIdentifier:    subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:     initiatingDeviceIdentifier,
		TimeRemaining:                  timeRemaining,
		Timestamp:                      timestamp,
		ListOfCovNotifications:         listOfCovNotifications,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(structType interface{}) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += m.Timestamp.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, error) {
	return BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriberProcessIdentifier")
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	subscriberProcessIdentifier := _subscriberProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriberProcessIdentifier")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initiatingDeviceIdentifier")
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	initiatingDeviceIdentifier := _initiatingDeviceIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initiatingDeviceIdentifier")
	}

	// Simple Field (timeRemaining)
	if pullErr := readBuffer.PullContext("timeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeRemaining")
	}
	_timeRemaining, _timeRemainingErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeRemainingErr != nil {
		return nil, errors.Wrap(_timeRemainingErr, "Error parsing 'timeRemaining' field of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	timeRemaining := _timeRemaining.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRemaining")
	}

	// Optional Field (timestamp) (Can be skipped, if a given expression evaluates to false)
	var timestamp BACnetTimeStampEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
		}
		_val, _err := BACnetTimeStampEnclosedParseWithBuffer(ctx, readBuffer, uint8(3))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timestamp' field of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		default:
			timestamp = _val.(BACnetTimeStampEnclosed)
			if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for timestamp")
			}
		}
	}

	// Simple Field (listOfCovNotifications)
	if pullErr := readBuffer.PullContext("listOfCovNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovNotifications")
	}
	_listOfCovNotifications, _listOfCovNotificationsErr := ListOfCovNotificationsListParseWithBuffer(ctx, readBuffer, uint8(uint8(4)))
	if _listOfCovNotificationsErr != nil {
		return nil, errors.Wrap(_listOfCovNotificationsErr, "Error parsing 'listOfCovNotifications' field of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	listOfCovNotifications := _listOfCovNotifications.(ListOfCovNotificationsList)
	if closeErr := readBuffer.CloseContext("listOfCovNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovNotifications")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SubscriberProcessIdentifier: subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:  initiatingDeviceIdentifier,
		TimeRemaining:               timeRemaining,
		Timestamp:                   timestamp,
		ListOfCovNotifications:      listOfCovNotifications,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
		}
		_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetSubscriberProcessIdentifier())
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for initiatingDeviceIdentifier")
		}
		_initiatingDeviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetInitiatingDeviceIdentifier())
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for initiatingDeviceIdentifier")
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (timeRemaining)
		if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeRemaining")
		}
		_timeRemainingErr := writeBuffer.WriteSerializable(ctx, m.GetTimeRemaining())
		if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeRemaining")
		}
		if _timeRemainingErr != nil {
			return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
		}

		// Optional Field (timestamp) (Can be skipped, if the value is null)
		var timestamp BACnetTimeStampEnclosed = nil
		if m.GetTimestamp() != nil {
			if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for timestamp")
			}
			timestamp = m.GetTimestamp()
			_timestampErr := writeBuffer.WriteSerializable(ctx, timestamp)
			if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for timestamp")
			}
			if _timestampErr != nil {
				return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
			}
		}

		// Simple Field (listOfCovNotifications)
		if pushErr := writeBuffer.PushContext("listOfCovNotifications"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfCovNotifications")
		}
		_listOfCovNotificationsErr := writeBuffer.WriteSerializable(ctx, m.GetListOfCovNotifications())
		if popErr := writeBuffer.PopContext("listOfCovNotifications"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfCovNotifications")
		}
		if _listOfCovNotificationsErr != nil {
			return errors.Wrap(_listOfCovNotificationsErr, "Error serializing 'listOfCovNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) isBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
