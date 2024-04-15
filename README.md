# NIXUP

![](vhs/nixup.gif)

Nixup is a command line tool for managing NixOS configuration.

> :warning: **Work in Progress**: This project is currently under development. Some features may not be complete and may change in the future.
## Installation

To install Nixup, you can clone the repository and compile the source code:

```sh
git clone https://github.com/alvaro17f/nixup.git
cd nixup
go build -o nixup
```

## Usage
To use Nixup, run the nixup command with the desired options. Here are some examples:

```sh
# Update the system
nixup --update

# Keep the last 5 generations
nixup --keep 5

# Show the differences between generations
nixup --diff
```

## License
Nixup is distributed under the MIT license. See the LICENSE file for more information.
