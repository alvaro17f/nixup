use std::process::exit;

use color_print::cformat;
use dialoguer::{theme::ColorfulTheme, Confirm};

use crate::{utils::commands::Cmd, Cli};
pub fn init(cli: Cli) -> std::io::Result<()> {
	let mut cmd = Cmd::new(cli);

	cmd.config();
	if Confirm::with_theme(&ColorfulTheme::default())
		.with_prompt(cformat!("<y>Do you want to update your system? (Y/n): "))
		.default(true)
		.interact()
		.unwrap()
	{
		cmd.git_pull();

		if cmd.cli.update {
			cmd.nix_update();
		}

		if cmd.git_diff() {
			cmd.git_status();
			if Confirm::with_theme(&ColorfulTheme::default())
				.with_prompt(cformat!("<y>Do you want to add this changes to the stage? (Y/n): "))
				.default(true)
				.interact()
				.unwrap()
			{
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
