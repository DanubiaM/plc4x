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
package org.apache.plc4x.java.cbus.readwrite;

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

public class SALDataRoomControlSystem extends SALData implements Message {

  // Accessors for discriminator values.
  public ApplicationId getApplicationId() {
    return ApplicationId.ROOM_CONTROL_SYSTEM;
  }

  public SALDataRoomControlSystem(SALData salData) {
    super(salData);
  }

  @Override
  protected void serializeSALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SALDataRoomControlSystem");

    writeBuffer.popContext("SALDataRoomControlSystem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SALDataRoomControlSystem _value = this;

    return lengthInBits;
  }

  public static SALDataRoomControlSystemBuilder staticParseBuilder(
      ReadBuffer readBuffer, ApplicationId applicationId) throws ParseException {
    readBuffer.pullContext("SALDataRoomControlSystem");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    // Validation
    if (!((1) == (2))) {
      throw new ParseValidationException("ROOM_CONTROL_SYSTEM Not yet implemented");
    }

    readBuffer.closeContext("SALDataRoomControlSystem");
    // Create the instance
    return new SALDataRoomControlSystemBuilder();
  }

  public static class SALDataRoomControlSystemBuilder implements SALData.SALDataBuilder {

    public SALDataRoomControlSystemBuilder() {}

    public SALDataRoomControlSystem build(SALData salData) {
      SALDataRoomControlSystem sALDataRoomControlSystem = new SALDataRoomControlSystem(salData);
      return sALDataRoomControlSystem;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SALDataRoomControlSystem)) {
      return false;
    }
    SALDataRoomControlSystem that = (SALDataRoomControlSystem) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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