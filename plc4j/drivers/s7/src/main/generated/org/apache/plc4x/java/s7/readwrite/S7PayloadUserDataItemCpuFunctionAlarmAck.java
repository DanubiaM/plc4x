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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class S7PayloadUserDataItemCpuFunctionAlarmAck extends S7PayloadUserDataItem
    implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionType() {
    return (byte) 0x04;
  }

  public Short getCpuSubfunction() {
    return (short) 0x0b;
  }

  public Integer getDataLength() {
    return 0;
  }

  // Properties.
  protected final short functionId;
  protected final List<AlarmMessageObjectAckType> messageObjects;

  public S7PayloadUserDataItemCpuFunctionAlarmAck(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      short functionId,
      List<AlarmMessageObjectAckType> messageObjects) {
    super(returnCode, transportSize);
    this.functionId = functionId;
    this.messageObjects = messageObjects;
  }

  public short getFunctionId() {
    return functionId;
  }

  public List<AlarmMessageObjectAckType> getMessageObjects() {
    return messageObjects;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadUserDataItemCpuFunctionAlarmAck");

    // Simple Field (functionId)
    writeSimpleField("functionId", functionId, writeUnsignedShort(writeBuffer, 8));

    // Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short numberOfObjects = (short) (COUNT(getMessageObjects()));
    writeImplicitField("numberOfObjects", numberOfObjects, writeUnsignedShort(writeBuffer, 8));

    // Array Field (messageObjects)
    writeComplexTypeArrayField("messageObjects", messageObjects, writeBuffer);

    writeBuffer.popContext("S7PayloadUserDataItemCpuFunctionAlarmAck");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCpuFunctionAlarmAck _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (functionId)
    lengthInBits += 8;

    // Implicit Field (numberOfObjects)
    lengthInBits += 8;

    // Array field
    if (messageObjects != null) {
      int i = 0;
      for (AlarmMessageObjectAckType element : messageObjects) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= messageObjects.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionType, Short cpuSubfunction) throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCpuFunctionAlarmAck");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short functionId = readSimpleField("functionId", readUnsignedShort(readBuffer, 8));

    short numberOfObjects = readImplicitField("numberOfObjects", readUnsignedShort(readBuffer, 8));

    List<AlarmMessageObjectAckType> messageObjects =
        readCountArrayField(
            "messageObjects",
            new DataReaderComplexDefault<>(
                () -> AlarmMessageObjectAckType.staticParse(readBuffer), readBuffer),
            numberOfObjects);

    readBuffer.closeContext("S7PayloadUserDataItemCpuFunctionAlarmAck");
    // Create the instance
    return new S7PayloadUserDataItemCpuFunctionAlarmAckBuilderImpl(functionId, messageObjects);
  }

  public static class S7PayloadUserDataItemCpuFunctionAlarmAckBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final short functionId;
    private final List<AlarmMessageObjectAckType> messageObjects;

    public S7PayloadUserDataItemCpuFunctionAlarmAckBuilderImpl(
        short functionId, List<AlarmMessageObjectAckType> messageObjects) {
      this.functionId = functionId;
      this.messageObjects = messageObjects;
    }

    public S7PayloadUserDataItemCpuFunctionAlarmAck build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize) {
      S7PayloadUserDataItemCpuFunctionAlarmAck s7PayloadUserDataItemCpuFunctionAlarmAck =
          new S7PayloadUserDataItemCpuFunctionAlarmAck(
              returnCode, transportSize, functionId, messageObjects);
      return s7PayloadUserDataItemCpuFunctionAlarmAck;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCpuFunctionAlarmAck)) {
      return false;
    }
    S7PayloadUserDataItemCpuFunctionAlarmAck that = (S7PayloadUserDataItemCpuFunctionAlarmAck) o;
    return (getFunctionId() == that.getFunctionId())
        && (getMessageObjects() == that.getMessageObjects())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFunctionId(), getMessageObjects());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
