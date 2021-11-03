package shell

import _ "embed" // yes

//go:embed files/sh_common_hooks.sh
var commonHooks string

const (
	hookStartMarker = "# Generated by Hermit; START; DO NOT EDIT."
	hookEndMarker   = "# Generated by Hermit; END; DO NOT EDIT."
)