{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  packages = with pkgs; [
    go
    golangci-lint
    godef
    go-outline
    gopls
    cassandra
  ];
}
