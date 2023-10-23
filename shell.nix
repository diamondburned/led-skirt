{ pkgs ? import <nixpkgs> {} }:

let
	fetchFromGitHub = pkgs.fetchFromGitHub;
in

let
	pkgs = import ./nix/pkgs.nix {
		inherit fetchFromGitHub;
		overlays = [
			(self: super: {
				go = super.go_1_21;
			})
		];
	};
in

with pkgs.lib;
with builtins;

let
	tinygo = pkgs.callPackage ./nix/tinygo.nix {};
	# tinygo = pkgs.tinygo;

	# Tinygo target for gopls to use.
	tinygoTarget = "xiao-rp2040";

	tinygoHook =
		with pkgs.lib;
		with builtins;
		''
			hookTinygoEnv() {
				vars=$(tinygo info -json -target ${tinygoTarget})

				export GOROOT=$(jq -r '.goroot' <<< "$vars")
				export GOARCH=$(jq -r '.goarch' <<< "$vars")
				export GOOS=$(jq -r '.goos' <<< "$vars")
				export GOFLAGS="-tags=$(jq -r '.build_tags | join(",")' <<< "$vars")"
				export GOWORK=
			}
		'';

	withTinygoHook = name: bin:
		pkgs.writeShellScriptBin name ''
			${tinygoHook}
			exec ${bin} "$@"
		'';

	gopls = withTinygoHook "gopls" "${pkgs.gopls}/bin/gopls";
	goimports = withTinygoHook "goimports" "${pkgs.gotools}/bin/goimports";

	staticcheck = pkgs.writeShellScriptBin "staticcheck" ''
		echo "Not running staticcheck for Tinygo" >&2
		exit 0
	'';
in

pkgs.mkShell {
	buildInputs = with pkgs; [
		niv
		go
		gopls
		gotools
		go-tools # staticcheck
		goimports
		staticcheck
	] ++ [ tinygo ];

	CGO_ENABLED = "1";
	TINYGO_TARGET = tinygoTarget;
}
