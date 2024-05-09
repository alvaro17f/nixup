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
    version = "0.2.2";
  in
  {
    defaultPackage.x86_64-linux = pkgs.buildGoModule {
      pname = "nixup";
      version = version;
      src = self;

      vendorHash = "sha256-ls73S+ZOnMJjhg5zZPo44sOFfh8Khh7rGDKWgXB15R4=";

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
