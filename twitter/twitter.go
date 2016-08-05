// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package twitter provides constants for using OAuth2 to access Github.
package twitter // import "golang.org/x/oauth2/twitter"

import (
	"golang.org/x/oauth2"
)

// Endpoint is Github's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	TokenURL: "https://api.twitter.com/oauth2/token",
}
