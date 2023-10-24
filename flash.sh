#!/usr/bin/env bash
set -eo pipefail

main() {
	if [[ "$1" == "-h" || "$1" == "--help" ]]; then
		log "Usage: $0 <preset> - flash given preset"
		log "	    $0          - list available presets"
		return
	fi

	if [[ ! "$1" ]]; then
		echo "Available presets:"
		for preset in preset_*.go; do
			preset=${preset#preset_}
			preset=${preset%.go}
			echo "- ${preset}"
		done
		return
	fi

	local preset="$1"

	# Assert preset exists
	if [[ ! -f preset_$preset.go ]]; then
		log "Preset $preset does not exist"
		return 1
	fi

	tinygo flash \
		-tags preset_$preset \
		-target $TINYGO_TARGET
}

log() {
	echo "$@" >&2
}

main "$@"
