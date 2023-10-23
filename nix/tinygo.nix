{ lib, stdenv, go }:

let
	version = "0.30.0";

	createURL = { GOOS, GOARCH, ... }:
		let
			base = "https://github.com/tinygo-org/tinygo/releases/download";
			name = "tinygo${version}.${GOOS}-${GOARCH}.tar.gz";
		in
			"${base}/v${version}/${name}";
in

fetchTarball (createURL go)
