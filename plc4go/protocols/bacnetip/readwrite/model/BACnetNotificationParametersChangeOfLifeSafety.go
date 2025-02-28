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

// BACnetNotificationParametersChangeOfLifeSafety is the corresponding interface of BACnetNotificationParametersChangeOfLifeSafety
type BACnetNotificationParametersChangeOfLifeSafety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewState returns NewState (property field)
	GetNewState() BACnetLifeSafetyStateTagged
	// GetNewMode returns NewMode (property field)
	GetNewMode() BACnetLifeSafetyModeTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetOperationExpected returns OperationExpected (property field)
	GetOperationExpected() BACnetLifeSafetyOperationTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersChangeOfLifeSafetyExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfLifeSafety.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfLifeSafetyExactly interface {
	BACnetNotificationParametersChangeOfLifeSafety
	isBACnetNotificationParametersChangeOfLifeSafety() bool
}

// _BACnetNotificationParametersChangeOfLifeSafety is the data-structure of this message
type _BACnetNotificationParametersChangeOfLifeSafety struct {
	*_BACnetNotificationParameters
	InnerOpeningTag   BACnetOpeningTag
	NewState          BACnetLifeSafetyStateTagged
	NewMode           BACnetLifeSafetyModeTagged
	StatusFlags       BACnetStatusFlagsTagged
	OperationExpected BACnetLifeSafetyOperationTagged
	InnerClosingTag   BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewState() BACnetLifeSafetyStateTagged {
	return m.NewState
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewMode() BACnetLifeSafetyModeTagged {
	return m.NewMode
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetOperationExpected() BACnetLifeSafetyOperationTagged {
	return m.OperationExpected
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfLifeSafety factory function for _BACnetNotificationParametersChangeOfLifeSafety
func NewBACnetNotificationParametersChangeOfLifeSafety(innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfLifeSafety {
	_result := &_BACnetNotificationParametersChangeOfLifeSafety{
		InnerOpeningTag:               innerOpeningTag,
		NewState:                      newState,
		NewMode:                       newMode,
		StatusFlags:                   statusFlags,
		OperationExpected:             operationExpected,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfLifeSafety(structType interface{}) BACnetNotificationParametersChangeOfLifeSafety {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfLifeSafety"
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (newState)
	lengthInBits += m.NewState.GetLengthInBits(ctx)

	// Simple field (newMode)
	lengthInBits += m.NewMode.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (operationExpected)
	lengthInBits += m.OperationExpected.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfLifeSafetyParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersChangeOfLifeSafety, error) {
	return BACnetNotificationParametersChangeOfLifeSafetyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersChangeOfLifeSafetyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersChangeOfLifeSafety, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfLifeSafety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (newState)
	if pullErr := readBuffer.PullContext("newState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for newState")
	}
	_newState, _newStateErr := BACnetLifeSafetyStateTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _newStateErr != nil {
		return nil, errors.Wrap(_newStateErr, "Error parsing 'newState' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	newState := _newState.(BACnetLifeSafetyStateTagged)
	if closeErr := readBuffer.CloseContext("newState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for newState")
	}

	// Simple Field (newMode)
	if pullErr := readBuffer.PullContext("newMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for newMode")
	}
	_newMode, _newModeErr := BACnetLifeSafetyModeTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _newModeErr != nil {
		return nil, errors.Wrap(_newModeErr, "Error parsing 'newMode' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	newMode := _newMode.(BACnetLifeSafetyModeTagged)
	if closeErr := readBuffer.CloseContext("newMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for newMode")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (operationExpected)
	if pullErr := readBuffer.PullContext("operationExpected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for operationExpected")
	}
	_operationExpected, _operationExpectedErr := BACnetLifeSafetyOperationTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _operationExpectedErr != nil {
		return nil, errors.Wrap(_operationExpectedErr, "Error parsing 'operationExpected' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	operationExpected := _operationExpected.(BACnetLifeSafetyOperationTagged)
	if closeErr := readBuffer.CloseContext("operationExpected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for operationExpected")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersChangeOfLifeSafety")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfLifeSafety")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfLifeSafety{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag:   innerOpeningTag,
		NewState:          newState,
		NewMode:           newMode,
		StatusFlags:       statusFlags,
		OperationExpected: operationExpected,
		InnerClosingTag:   innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfLifeSafety")
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

		// Simple Field (newState)
		if pushErr := writeBuffer.PushContext("newState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for newState")
		}
		_newStateErr := writeBuffer.WriteSerializable(ctx, m.GetNewState())
		if popErr := writeBuffer.PopContext("newState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for newState")
		}
		if _newStateErr != nil {
			return errors.Wrap(_newStateErr, "Error serializing 'newState' field")
		}

		// Simple Field (newMode)
		if pushErr := writeBuffer.PushContext("newMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for newMode")
		}
		_newModeErr := writeBuffer.WriteSerializable(ctx, m.GetNewMode())
		if popErr := writeBuffer.PopContext("newMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for newMode")
		}
		if _newModeErr != nil {
			return errors.Wrap(_newModeErr, "Error serializing 'newMode' field")
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

		// Simple Field (operationExpected)
		if pushErr := writeBuffer.PushContext("operationExpected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for operationExpected")
		}
		_operationExpectedErr := writeBuffer.WriteSerializable(ctx, m.GetOperationExpected())
		if popErr := writeBuffer.PopContext("operationExpected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for operationExpected")
		}
		if _operationExpectedErr != nil {
			return errors.Wrap(_operationExpectedErr, "Error serializing 'operationExpected' field")
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

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfLifeSafety")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) isBACnetNotificationParametersChangeOfLifeSafety() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
