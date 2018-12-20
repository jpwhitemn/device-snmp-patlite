// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
//
package driver

import (
	"fmt"
	"time"

	ds_models "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/edgex-go/pkg/clients/logging"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

type SNMPDriver struct {
	lc           logger.LoggingClient
	asyncCh      chan<- *ds_models.AsyncValues
	switchButton bool
}

// DisconnectDevice handles protocol-specific cleanup when a device
// is removed.
func (s *SNMPDriver) DisconnectDevice(address *models.Addressable) error {
	return nil
}

// Initialize performs protocol-specific initialization for the device
// service.
func (s *SNMPDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *ds_models.AsyncValues) error {
	s.lc = lc
	s.asyncCh = asyncCh
	return nil
}

// HandleReadCommands triggers a protocol Read operation for the specified device.
func (s *SNMPDriver) HandleReadCommands(addr *models.Addressable, reqs []ds_models.CommandRequest) (res []*ds_models.CommandValue, err error) {
	//TODO get operation and value from command request
	myOp:= "GREEN"
	cmd := NewGetDeviceCommand(myOp)
	client := NewSNMPClient("192.168.0.20")

	if len(reqs) != 1 {
		err = fmt.Errorf("SNMPDriver.HandleReadCommands; too many command requests; only one supported")
		return
	}

	s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleReadCommands: device: %s operation: %v attributes: %v", addr.Name, reqs[0].RO.Operation, reqs[0].DeviceObject.Attributes))

	res = make([]*ds_models.CommandValue, 1)
	now := time.Now().UnixNano() / int64(time.Millisecond)
	val, err2 := client.GetValue(cmd)

	if err2 != nil {
		s.lc.Error(fmt.Sprintf("SNMPDriver.HandleReadCommands; %s", err2))
		return
	}
	// TODO param the operation
	s.lc.Debug(fmt.Sprintf("Value of GREEN light is: %d\n ------", val))
	// TODO change command type
	cv, _ := ds_models.NewInt32Value(&reqs[0].RO, now, int32(val))
	res[0] = cv

	return
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource (aka DeviceObject).
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (s *SNMPDriver) HandleWriteCommands(addr *models.Addressable, reqs []ds_models.CommandRequest,
	params []*ds_models.CommandValue) error {

	if len(reqs) != 1 {
		err := fmt.Errorf("SNMPDriver.HandleWriteCommands; too many command requests; only one supported")
		return err
	}
	if len(params) != 1 {
		err := fmt.Errorf("SNMPDriver.HandleWriteCommands; the number of parameter is not correct; only one supported")
		return err
	}

	s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleWriteCommands: device: %s, operation: %v, parameters: %v", addr.Name, reqs[0].RO.Operation, params))
	var err error
	if s.switchButton, err = params[0].BoolValue(); err != nil {
		err := fmt.Errorf("SNMPDriver.HandleWriteCommands; the data type of parameter should be Boolean, parameter: %s", params[0].String())
		return err
	}

	return nil
}

// Stop the protocol-specific DS code to shutdown gracefully, or
// if the force parameter is 'true', immediately. The driver is responsible
// for closing any in-use channels, including the channel used to send async
// readings (if supported).
func (s *SNMPDriver) Stop(force bool) error {
	s.lc.Debug(fmt.Sprintf("SNMPDriver.Stop called: force=%v", force))
	return nil
}
