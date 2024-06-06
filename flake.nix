{
  description = "A basic dev shell for nix/nixos users";

  inputs = {
    nixpkgs = { url = "github:NixOs/nixpkgs/nixos-unstable"; };
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
            go_1_21

            go-swagger
          ]);
        };
      });
}
