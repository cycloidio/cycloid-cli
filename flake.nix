{
  description = "A basic dev shell for nix/nixos users";

  inputs = {
    nixpkgs = { url = "https://github.com/NixOS/nixpkgs/archive/9957cd48326fe8dbd52fdc50dd2502307f188b0d.tar.gz"; };
    flake-utils = { url = "github:numtide/flake-utils"; };
  };

  outputs = inputs@{ self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        lib = pkgs.lib;
        swaggerPython = (pkgs.python312.withPackages (p: with p; [
          pyyaml
        ]));
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = [ swaggerPython ]
            ++ (with pkgs; [
            # You packages here
            gnumake

            # Prod is built using go 1.18 rn
            go_1_18
          ]);
        };
      });
}
