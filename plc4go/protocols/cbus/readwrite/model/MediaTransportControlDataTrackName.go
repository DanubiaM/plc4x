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

// MediaTransportControlDataTrackName is the corresponding interface of MediaTransportControlDataTrackName
type MediaTransportControlDataTrackName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetTrackName returns TrackName (property field)
	GetTrackName() string
}

// MediaTransportControlDataTrackNameExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataTrackName.
// This is useful for switch cases.
type MediaTransportControlDataTrackNameExactly interface {
	MediaTransportControlDataTrackName
	isMediaTransportControlDataTrackName() bool
}

// _MediaTransportControlDataTrackName is the data-structure of this message
type _MediaTransportControlDataTrackName struct {
	*_MediaTransportControlData
	TrackName string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataTrackName) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataTrackName) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataTrackName) GetTrackName() string {
	return m.TrackName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataTrackName factory function for _MediaTransportControlDataTrackName
func NewMediaTransportControlDataTrackName(trackName string, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataTrackName {
	_result := &_MediaTransportControlDataTrackName{
		TrackName:                  trackName,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataTrackName(structType interface{}) MediaTransportControlDataTrackName {
	if casted, ok := structType.(MediaTransportControlDataTrackName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataTrackName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataTrackName) GetTypeName() string {
	return "MediaTransportControlDataTrackName"
}

func (m *_MediaTransportControlDataTrackName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (trackName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataTrackName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataTrackNameParse(theBytes []byte, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataTrackName, error) {
	return MediaTransportControlDataTrackNameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func MediaTransportControlDataTrackNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataTrackName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataTrackName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataTrackName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (trackName)
	_trackName, _trackNameErr := readBuffer.ReadString("trackName", uint32(((commandTypeContainer.NumBytes())-(1))*(8)), "UTF-8")
	if _trackNameErr != nil {
		return nil, errors.Wrap(_trackNameErr, "Error parsing 'trackName' field of MediaTransportControlDataTrackName")
	}
	trackName := _trackName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataTrackName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataTrackName")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataTrackName{
		_MediaTransportControlData: &_MediaTransportControlData{},
		TrackName:                  trackName,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataTrackName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataTrackName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataTrackName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataTrackName")
		}

		// Simple Field (trackName)
		trackName := string(m.GetTrackName())
		_trackNameErr := writeBuffer.WriteString("trackName", uint32(((m.GetCommandTypeContainer().NumBytes())-(1))*(8)), "UTF-8", (trackName))
		if _trackNameErr != nil {
			return errors.Wrap(_trackNameErr, "Error serializing 'trackName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataTrackName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataTrackName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataTrackName) isMediaTransportControlDataTrackName() bool {
	return true
}

func (m *_MediaTransportControlDataTrackName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
