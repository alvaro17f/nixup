mod app;
mod utils;

use app::app;
use clap::{command, CommandFactory, Parser, Subcommand};
use clap_complete::Shell;

use crate::utils::tools::{get_home_dir, get_hostname};

#[derive(Parser, Debug, PartialEq)]
#[command(author, version, about, long_about = None)]
struct Cli {
	/// List of available commands
	#[command(subcommand)]
	commands: Option<Commands>,
	/// Path to the dotfiles repository
	#[arg(short = 'r', long = "repo", default_value_t = format!("{}/.dotfiles", get_home_dir()))]
	repo: String,
	/// Hostname
	#[arg(short = 'n', long = "hostname", default_value_t = get_hostname() )]
	hostname: String,
	/// Number of generations to keep
	#[arg(short = 'k', long = "keep", default_value = "10")]
	keep: i32,
	/// Update flake lock files
	#[arg(short = 'u', long = "update", default_value = "false")]
	update: bool,
	/// Show the difference between the current and the last generation
	#[arg(short = 'd', long = "diff", default_value = "false")]
	diff: bool,
}
#[derive(Subcommand, Debug, PartialEq)]
enum Commands {
	/// Generate tab-completion scripts for your shell
	Completions {
		#[clap(value_enum)]
		shell: Shell,
	},
}

fn main() -> std::io::Result<()> {
	let cli = Cli::parse();
	handle_commands(cli)?;
	Ok(())
}

fn handle_commands(cli: Cli) -> std::io::Result<()> {
	match cli.commands {
		Some(Commands::Completions { shell }) => {
			clap_complete::generate(shell, &mut Cli::command(), "nixup", &mut std::io::stdout().lock());
		}
		None => {
			app(cli).unwrap();
		}
	}
	Ok(())
}
