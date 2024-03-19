use std::{io::{BufRead, BufReader}, process::{Command, Stdio}};

use color_print::cprintln;
use nix::unistd::geteuid;

use crate::{utils::tools::title_maker, Cli};

/// Application.
#[derive(Debug)]
pub struct Cmd {
	pub cli: Cli,
}

impl Cmd {
	pub fn new(cli: Cli) -> Self { Cmd { cli } }

	fn exec_command(&mut self, command: &str) {
		let mut cmd = Command::new("sh")
			.arg("-c")
			.arg(command)
			.current_dir(&self.cli.repo)
			.stdout(Stdio::piped())
			.spawn()
			.expect("Failed to execute command");

		let stdout = cmd.stdout.take().unwrap();
		let reader = BufReader::new(stdout);

		for line in reader.lines() {
			println!("{}", line.unwrap());
		}

		cmd.wait().expect("Failed to wait on command");
	}

	fn exec_sudo_command(&mut self, command: &str) {
		if geteuid().is_root() {
			self.exec_command(command)
		} else {
			let mut cmd = Command::new("sudo")
				.arg("sh")
				.arg("-c")
				.arg(command)
				.stdin(Stdio::piped())
				.stdout(Stdio::piped())
				.spawn()
				.expect("Failed to execute command");

			let stdout = cmd.stdout.take().unwrap();
			let reader = BufReader::new(stdout);

			for line in reader.lines() {
				println!("{}", line.unwrap());
			}

			cmd.wait().expect("Failed to wait on command");
		}
	}

	pub fn config(&mut self) {
		let repo = &self.cli.repo;
		let hostname = &self.cli.hostname;
		let update = &self.cli.update;
		let keep = &self.cli.keep;
		let diff = &self.cli.diff;

		title_maker("Configuration");
		cprintln!(
			r#"<b>
• Repo: <r>{repo}</r>
• Hostname: <r>{hostname}</r>
• Update: <r>{update}</r>
• keep: <r>{keep}</r>
• Diff: <r>{diff}</r>
			</b>"#
		)
	}

	pub fn git_pull(&mut self) {
		title_maker("Git Pull");
		self.exec_command("git pull")
	}

	pub fn git_diff(&mut self) -> bool {
		let command = "git diff --exit-code";

		let output = Command::new("sh")
			.arg("-c")
			.arg(command)
			.current_dir(&self.cli.repo)
			.output()
			.expect("Failed to execute command");

		!output.status.success()
	}

	pub fn git_status(&mut self) {
		title_maker("Git Status");
		self.exec_command("git status --porcelain")
	}

	pub fn git_add(&mut self) { self.exec_command("git add .") }

	pub fn nix_update(&mut self) {
		title_maker("Nix Update");
		self.exec_command("nix flake update")
	}

	pub fn nix_reduild(&mut self) {
		title_maker("Nix Rebuild");
		self.exec_sudo_command(&format!(
			"nixos-rebuild switch --flake {}#{} --show-trace",
			&self.cli.repo, &self.cli.hostname,
		))
	}

	pub fn nix_keep_last(&mut self) {
		self.exec_sudo_command(&format!(
			"nix-env --profile /nix/var/nix/profiles/system --delete-generations +{}",
			&self.cli.keep,
		))
	}

	pub fn nix_diff(&mut self) {
		title_maker("Nix Diff");
		self.exec_command(
			"nix profile diff-closures --profile /nix/var/nix/profiles/system | tac | awk '/Version/{print; exit} 1' | tac",
		)
	}
}
