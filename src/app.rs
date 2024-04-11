use std::process::exit;

use crate::{
	utils::{
		commands::Cmd,
		tools::{confirm, get_name},
	},
	Cli,
};

pub fn app(cli: Cli) -> std::io::Result<()> {
	let name = get_name();
	let mut cmd = Cmd::new(cli);

	cmd.config();

	if confirm(&format!("Hi {}, Do you want to update your system? (Y/n): ", name), true) {
		cmd.git_pull();

		if cmd.cli.update {
			cmd.nix_update();
		}

		if cmd.git_diff() {
			cmd.git_status();
			if confirm("Do you want to add this changes to the stage? (Y/n): ", true) {
				cmd.git_add();
			};
		};

		cmd.nix_reduild();
		cmd.nix_keep_last();

		if cmd.cli.diff {
			cmd.nix_diff();
		}
	} else {
		exit(0)
	}

	Ok(())
}
