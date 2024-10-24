// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/omec-project/nas/logger"
	"github.com/omec-project/nas/nasMessage"
	"github.com/omec-project/nas/nasType"
	"github.com/stretchr/testify/assert"
)

type nasMessageAuthenticationRejectData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inAuthenticationRejectMessageIdentity uint8
	inEAPMessage                          nasType.EAPMessage
}

var nasMessageAuthenticationRejectTable = []nasMessageAuthenticationRejectData{
	{
		inExtendedProtocolDiscriminator:       0x01,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inAuthenticationRejectMessageIdentity: 0x01,
		inEAPMessage:                          nasType.EAPMessage{Iei: nasMessage.AuthenticationRejectEAPMessageType, Len: 2, Buffer: []byte{0x00, 0x00}},
	},
}

func TestNasTypeNewAuthenticationReject(t *testing.T) {
	a := nasMessage.NewAuthenticationReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationRejectMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationRejectMessage---")
	for i, table := range nasMessageAuthenticationRejectTable {
		t.Logf("Test Cnt:%d", i)

		a := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		b := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		assert.NotNil(t, a)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.AuthenticationRejectMessageIdentity.SetMessageType(table.inAuthenticationRejectMessageIdentity)

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationReject(buff)
		logger.NasMsgLog.Debugln(a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationReject(&data)
		logger.NasMsgLog.Debugln("decode:", data)
		logger.NasMsgLog.Infoln(b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
