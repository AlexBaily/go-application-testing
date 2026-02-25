{
  description = "Go application nix packages";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;  # ← Add this line
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go
            go_1_25
            gopls
            gotools
            go-tools
            golangci-lint

            # Database tools
            sqlc
            postgresql

            # Infrastructure
            terraform
            
            # Docker & Container tools
            docker-compose

            # Development tools
            git
            gnumake
            curl
            jq
          ];

          shellHook = ''
            export LC_ALL=en_US.UTF-8
            export LANG=en_US.UTF-8

            echo "Development environment loaded!"
            echo "Run 'make' to see available commands"
          '';

          # Environment variables
          CGO_ENABLED = "0";
          GOPATH = "${builtins.toString ./.}/.go";
        };
      }
    );
}