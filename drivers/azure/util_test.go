package azure

import (
	"testing"

	network "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/stretchr/testify/assert"
)

func TestParseSecurityRuleProtocol(t *testing.T) {
	tests := []struct {
		raw           string
		expectedProto network.SecurityRuleProtocol
		expectedErr   bool
	}{
		{"tcp", network.SecurityRuleProtocolTCP, false},
		{"udp", network.SecurityRuleProtocolUDP, false},
		{"*", network.SecurityRuleProtocolAsterisk, false},
		{"Invalid", "", true},
	}

	for _, tc := range tests {
		proto, err := parseSecurityRuleProtocol(tc.raw)
		assert.Equal(t, tc.expectedProto, proto)
		if tc.expectedErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
