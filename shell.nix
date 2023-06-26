with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    go
    golint
    gopls
    sqlite
    flyctl
    golangci-lint
    entr
    google-cloud-sdk
    nodejs
  ];
}
