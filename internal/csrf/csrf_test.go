package csrf

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCSRF_Check(t *testing.T) {
	hashToken := NewHMACHashToken("secret")

	cookie := "cookie"
	tests := []struct {
		name        string
		inputToken  string
		expected    bool
		expectedErr error
	}{
		{
			name: "valid token",
			inputToken: func() string {
				token, _ := hashToken.Create(cookie, time.Now().Add(time.Hour).Unix())
				return token
			}(),
			expected:    true,
			expectedErr: nil,
		},
		{
			name:        "bad token data",
			inputToken:  "token",
			expected:    false,
			expectedErr: fmt.Errorf("bad token data"),
		},
		{
			name: "token expired",
			inputToken: func() string {
				token, _ := hashToken.Create(cookie, -1)
				return token
			}(),
			expected:    false,
			expectedErr: fmt.Errorf("token expired"),
		},
		{
			name:        "bad token time",
			inputToken:  "token:a111",
			expected:    false,
			expectedErr: fmt.Errorf("bad token time"),
		},
		{
			name:        "can't hex decode token",
			inputToken:  "awce1:30000000000",
			expected:    false,
			expectedErr: fmt.Errorf("can't hex decode token"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			isValid, err := hashToken.Check(cookie, th.inputToken)
			if th.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, th.expectedErr, err)
				assert.Equal(t, test.expected, isValid)
			} else {
				assert.Equal(t, test.expected, isValid)
				assert.NoError(t, err)
			}
		})
	}
}
