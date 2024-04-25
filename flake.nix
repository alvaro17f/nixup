{
  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixos-unstable";
    };
  };

  outputs = { self, nixpkgs, ... }:
  let
    pkgs = import nixpkgs {
      system = "x86_64-linux";
    };
    lib = pkgs.lib;
    version = "0.2.1";
  in
  {
    defaultPackage.x86_64-linux = pkgs.buildGoModule {
      pname = "nixup";
      version = version;
      src = self;

      vendorHash = "sha256-0iNKOvdFFlqMw/63p2qXv1j5GmYbYOniLP9QeSBeOfg=";

      CGO_ENABLED = 0;

      outputs = [ "out" ];

      ldflags = [
        "-s" "-w" "-X main.version=${version}"
      ];

      meta = with lib; {
        homepage = "https://github.com/alvaro17f/nixup";
        description = "nixup";
        license = licenses.mit;
        maintainers = with maintainers; [ alvaro17f ];
        platforms = platforms.unix;
        #changelog = "https://github.com/alvaro17f/nixup/blob/${version}/CHANGELOG.md";
        mainProgram = "nixup";
      };
    };
  };
}
