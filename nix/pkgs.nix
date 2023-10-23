{ ... }@args:

let
	src = import ./sources.nix {};
in
	import src.nixpkgs args
