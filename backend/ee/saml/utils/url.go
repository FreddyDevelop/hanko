package utils

import (
	"github.com/FreddyDevelop/hanko/backend/v2/ee/saml/config"
	"strings"
)

func IsAllowedRedirect(config config.Saml, redirectTo string) bool {
	if redirectTo == "" {
		return false
	}

	redirectTo = strings.TrimSuffix(redirectTo, "/")

	for _, pattern := range config.AllowedRedirectURLMap {
		if pattern.Match(redirectTo) {
			return true
		}
	}

	return false
}
