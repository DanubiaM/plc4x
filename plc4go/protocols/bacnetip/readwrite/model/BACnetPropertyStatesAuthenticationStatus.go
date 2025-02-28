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

// BACnetPropertyStatesAuthenticationStatus is the corresponding interface of BACnetPropertyStatesAuthenticationStatus
type BACnetPropertyStatesAuthenticationStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetAuthenticationStatus returns AuthenticationStatus (property field)
	GetAuthenticationStatus() BACnetAuthenticationStatusTagged
}

// BACnetPropertyStatesAuthenticationStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesAuthenticationStatus.
// This is useful for switch cases.
type BACnetPropertyStatesAuthenticationStatusExactly interface {
	BACnetPropertyStatesAuthenticationStatus
	isBACnetPropertyStatesAuthenticationStatus() bool
}

// _BACnetPropertyStatesAuthenticationStatus is the data-structure of this message
type _BACnetPropertyStatesAuthenticationStatus struct {
	*_BACnetPropertyStates
	AuthenticationStatus BACnetAuthenticationStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAuthenticationStatus) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAuthenticationStatus) GetAuthenticationStatus() BACnetAuthenticationStatusTagged {
	return m.AuthenticationStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAuthenticationStatus factory function for _BACnetPropertyStatesAuthenticationStatus
func NewBACnetPropertyStatesAuthenticationStatus(authenticationStatus BACnetAuthenticationStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesAuthenticationStatus {
	_result := &_BACnetPropertyStatesAuthenticationStatus{
		AuthenticationStatus:  authenticationStatus,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAuthenticationStatus(structType interface{}) BACnetPropertyStatesAuthenticationStatus {
	if casted, ok := structType.(BACnetPropertyStatesAuthenticationStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAuthenticationStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetTypeName() string {
	return "BACnetPropertyStatesAuthenticationStatus"
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (authenticationStatus)
	lengthInBits += m.AuthenticationStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesAuthenticationStatusParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesAuthenticationStatus, error) {
	return BACnetPropertyStatesAuthenticationStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesAuthenticationStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesAuthenticationStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAuthenticationStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAuthenticationStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (authenticationStatus)
	if pullErr := readBuffer.PullContext("authenticationStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationStatus")
	}
	_authenticationStatus, _authenticationStatusErr := BACnetAuthenticationStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _authenticationStatusErr != nil {
		return nil, errors.Wrap(_authenticationStatusErr, "Error parsing 'authenticationStatus' field of BACnetPropertyStatesAuthenticationStatus")
	}
	authenticationStatus := _authenticationStatus.(BACnetAuthenticationStatusTagged)
	if closeErr := readBuffer.CloseContext("authenticationStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAuthenticationStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAuthenticationStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesAuthenticationStatus{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		AuthenticationStatus:  authenticationStatus,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAuthenticationStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAuthenticationStatus")
		}

		// Simple Field (authenticationStatus)
		if pushErr := writeBuffer.PushContext("authenticationStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authenticationStatus")
		}
		_authenticationStatusErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationStatus())
		if popErr := writeBuffer.PopContext("authenticationStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authenticationStatus")
		}
		if _authenticationStatusErr != nil {
			return errors.Wrap(_authenticationStatusErr, "Error serializing 'authenticationStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAuthenticationStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAuthenticationStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAuthenticationStatus) isBACnetPropertyStatesAuthenticationStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesAuthenticationStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
