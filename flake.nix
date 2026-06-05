{
  description = "Navidrome lyrics scrape plugin dev environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            tinygo
            gopls
            gofumpt
            zip
            just
          ];

          shellHook = ''
            echo "navidrome plugin dev shell"
            echo "  go      $(go version | awk '{print $3}')"
            echo "  tinygo  $(tinygo version 2>/dev/null | awk '{print $3}')"
            echo ""
            echo "run 'just' to list available recipes"
          '';
        };
      });
}
