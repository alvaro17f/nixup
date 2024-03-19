#![allow(dead_code)]

use std::{io::{self, BufRead, Result}, process::Command};

use color_print::cprintln;

pub fn clear() -> Result<()> {
	Command::new("clear").status()?;
	Ok(())
}

pub fn clear_previous_line() {
	print!("\x1b[1A\x1b[2K");
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

pub fn get_hostname() -> String {
	let hostname = hostname::get();
	match hostname {
		Ok(h) => h.to_string_lossy().to_string(),
		Err(_) => panic!(
			"{:?}",
			"Failed to get hostname. Please consider passing it using `-n` flag".to_string()
		),
	}
}
pub fn get_home_dir() -> String {
	let home = dirs::home_dir();
	match home {
		Some(h) => h.to_string_lossy().to_string(),
		None => {
			panic!(
				"{:?}",
				"Failed to get home directory. Please consider passing it using `-r` flag".to_string()
			)
		}
	}
}
