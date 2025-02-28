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

public class COTPPacketDisconnectResponse extends COTPPacket implements Message {

  // Accessors for discriminator values.
  public Short getTpduCode() {
    return (short) 0xC0;
  }

  // Properties.
  protected final int destinationReference;
  protected final int sourceReference;

  public COTPPacketDisconnectResponse(
      List<COTPParameter> parameters,
      S7Message payload,
      int destinationReference,
      int sourceReference) {
    super(parameters, payload);
    this.destinationReference = destinationReference;
    this.sourceReference = sourceReference;
  }

  public int getDestinationReference() {
    return destinationReference;
  }

  public int getSourceReference() {
    return sourceReference;
  }

  @Override
  protected void serializeCOTPPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("COTPPacketDisconnectResponse");

    // Simple Field (destinationReference)
    writeSimpleField(
        "destinationReference", destinationReference, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (sourceReference)
    writeSimpleField("sourceReference", sourceReference, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("COTPPacketDisconnectResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    COTPPacketDisconnectResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (destinationReference)
    lengthInBits += 16;

    // Simple field (sourceReference)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static COTPPacketBuilder staticParseCOTPPacketBuilder(
      ReadBuffer readBuffer, Integer cotpLen) throws ParseException {
    readBuffer.pullContext("COTPPacketDisconnectResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int destinationReference =
        readSimpleField("destinationReference", readUnsignedInt(readBuffer, 16));

    int sourceReference = readSimpleField("sourceReference", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("COTPPacketDisconnectResponse");
    // Create the instance
    return new COTPPacketDisconnectResponseBuilderImpl(destinationReference, sourceReference);
  }

  public static class COTPPacketDisconnectResponseBuilderImpl
      implements COTPPacket.COTPPacketBuilder {
    private final int destinationReference;
    private final int sourceReference;

    public COTPPacketDisconnectResponseBuilderImpl(int destinationReference, int sourceReference) {
      this.destinationReference = destinationReference;
      this.sourceReference = sourceReference;
    }

    public COTPPacketDisconnectResponse build(List<COTPParameter> parameters, S7Message payload) {
      COTPPacketDisconnectResponse cOTPPacketDisconnectResponse =
          new COTPPacketDisconnectResponse(
              parameters, payload, destinationReference, sourceReference);
      return cOTPPacketDisconnectResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof COTPPacketDisconnectResponse)) {
      return false;
    }
    COTPPacketDisconnectResponse that = (COTPPacketDisconnectResponse) o;
    return (getDestinationReference() == that.getDestinationReference())
        && (getSourceReference() == that.getSourceReference())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDestinationReference(), getSourceReference());
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
