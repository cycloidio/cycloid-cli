{
  description = "A basic dev shell for nix/nixos users";

  inputs = {
    nixpkgs = { url = "github:NixOs/nixpkgs/nixos-25.05"; };
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
            libgcc

            go_1_25
            go-swagger
            # gci
            # golangci-lint
            # golangci-lint-langserver
            awscli
            docker
            pre-commit
          ]);
        };
      });
}
