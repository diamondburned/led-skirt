{ fetchFromGitHub, ... }@args:

let
	src = fetchFromGitHub {
		owner = "NixOS";
		repo = "nixpkgs";
		rev = "834f66f1edc6b54d1feab0f99e33051c7d1b8e07";
		sha256 = "sha256:0y5jwrmkj274aniqvymz1mgsaf8pmkrsyg095z7zwd1w597mss68";
	};
in
	import src args
