package authtoken

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractToken(t *testing.T) {
	req := require.New(t)

	cases := map[string]struct {
		authHeader   string
		want         string
		isError      bool
		errorMessage string
	}{
		"success extract": {
			authHeader: "Bearer test_token",
			want:       "test_token",
		},
		"more space": {
			authHeader: "Bearer             test_token",
			want:       "test_token",
		},
		"space after token": {
			authHeader: "Bearer  test_token  ",
			want:       "test_token",
		},
		"empty header": {
			isError:      true,
			errorMessage: "auth header is empty",
		},
		"without bearer template": {
			authHeader:   "test_token",
			isError:      true,
			errorMessage: ErrInvalidToken.Error(),
		},
		"only bearer template": {
			authHeader:   "Bearer   ",
			isError:      true,
			errorMessage: ErrInvalidToken.Error(),
		},
		"without space": {
			authHeader:   "Bearertest_token",
			isError:      true,
			errorMessage: ErrInvalidToken.Error(),
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			token, err := ExtractToken(cs.authHeader)
			if cs.isError {
				req.Error(err)
				req.Contains(err.Error(), cs.errorMessage)
			} else {
				req.NoError(err)
			}
			req.Equal(cs.want, token)
		})
	}
}
