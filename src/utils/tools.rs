#![allow(dead_code)]

use color_print::{cformat, cprintln};
use dialoguer::{theme::ColorfulTheme, Confirm};
use std::{
	io::{self, BufRead, Result},
	process::Command,
};
use whoami;

pub fn clear() -> Result<()> {
	Command::new("clear").status()?;
	Ok(())
}

pub fn clear_previous_line() {
	print!("\x1b[1A\x1b[2K");
}

pub fn confirm(prompt: &str, default_value: bool) -> bool {
	Confirm::with_theme(&ColorfulTheme::default())
		.with_prompt(cformat!("<y>{}</y>", prompt))
		.default(default_value)
		.interact()
		.unwrap_or(false)
}

pub fn pause() -> Result<()> {
	cprintln!("Press <c>'Enter'</c> to continue...");
	io::stdin().lock().lines().next();
	clear()?;
	Ok(())
}

pub fn title_maker(text: &str) {
	let text_len = text.chars().count();
	let borders = "*".repeat(text_len + 4);
	cprintln!("\n<b>{}\n* <r>{}</r> *\n{}</b>", borders, text, borders);
}

pub fn get_name() -> String {
	whoami::realname()
}

pub fn get_username() -> String {
	whoami::username()
}

pub fn get_hostname() -> String {
	whoami::fallible::hostname().expect("Failed to get hostname")
}

pub fn get_home_dir() -> String {
	format!("/home/{}", get_username())
}
