#[macro_use]
extern crate diesel;
extern crate dotenv;

use regex::Regex;
use std::process::Command;

mod conn;
use conn::establish_connection;

mod model;
use model::*;

fn is_ipv4(text: &str) -> bool {
    let re = Regex::new(
        r"^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$",
    ).unwrap();
    return re.is_match(text);
}

fn main() {
    let curl_target = "inet-ip.info";
    let curl_result = Command::new("curl")
        .args([curl_target])
        .output()
        .expect("failed to call `curl`");

    let ipv4_address = String::from_utf8_lossy(&curl_result.stdout).replace("\n", "");
    assert!(is_ipv4(&ipv4_address));

    let conn = establish_connection();
    let effective_records = get_effective_records(&conn);
    assert!(effective_records.len() <= 1);

    if effective_records.len() == 0 {
        insert_record(&conn, ipv4_address);
        return;
    }

    if ipv4_address == effective_records[0].ipv4_address {
        update_last_checked_at(&conn, &effective_records[0].id);
        return;
    }
    disable_record(&conn, &effective_records[0].id);
    insert_record(&conn, ipv4_address);
}
