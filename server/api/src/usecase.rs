use crate::db;

pub fn get_effective_ipv4_record() -> Result<db::Ipv4Record, String> {
  let conn = db::establish_connection();
  let effective_records = db::get_effective_records(&conn);

  if effective_records.len() == 0 {
    return Err("IPv4 record not found".into());
  }

  if effective_records.len() > 1 {
    return Err("Too many IPv4 records have been found".into());
  }

  return Ok(effective_records[0].clone());
}

use chrono::{offset::TimeZone, Utc};
use diesel::pg::PgConnection;
use std::path::Path;

#[test]
fn test_get_ip() {
  struct Testcase {
    name: String,
    sql_file_path: String,
    expect: Result<db::Ipv4Record, String>,
  }

  let testcases = vec![
    Testcase {
      name: "case_001".to_string(),
      sql_file_path: "./src/testcase/get_ip/case_001.sql".to_string(),
      expect: Ok(db::Ipv4Record {
        id: 2,
        ipv4_address: "112.112.112.112".to_string(),
        effective_flg: true,
        last_checked_at: match Utc.datetime_from_str("2022/01/02 00:00:00", "%Y/%m/%d %H:%M:%S") {
          Err(why) => panic!("datetime parse error:{}", why),
          Ok(datetime) => datetime,
        },
      }),
    },
    Testcase {
      name: "case_002".to_string(),
      sql_file_path: "./src/testcase/get_ip/case_002.sql".to_string(),
      expect: Err("IPv4 record not found".to_string()),
    },
    Testcase {
      name: "case_002".to_string(),
      sql_file_path: "./src/testcase/get_ip/case_002.sql".to_string(),
      expect: Err("IPv4 record not found".to_string()),
    },
    Testcase {
      name: "case_003".to_string(),
      sql_file_path: "./src/testcase/get_ip/case_003.sql".to_string(),
      expect: Err("Too many IPv4 records have been found".to_string()),
    },
  ];

  let conn = db::establish_connection();
  fn reset_db(conn: &PgConnection) {
    let reset = Path::new("./src/testcase/get_ip/reset.sql");
    db::exec_sql(reset, &conn).expect("unexpected error on reset");
  }

  for testcase in testcases {
    reset_db(&conn);

    let path = Path::new(&testcase.sql_file_path);
    db::exec_sql(path, &conn).expect(&format!("error on {}", testcase.name));

    let result = get_effective_ipv4_record();
    assert_eq!(result, testcase.expect);
  }
}
