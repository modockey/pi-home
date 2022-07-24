use chrono::{DateTime, Utc};

use diesel::pg::PgConnection;
use diesel::prelude::*;
use diesel::sql_query;

use std::fs::File;
use std::io::prelude::*;

use std::path::Path;

mod schema;
use schema::ipv4_history::dsl::*;

use dotenv::dotenv;
use std::env;

pub fn establish_connection() -> PgConnection {
  if cfg!(test) | cfg!(debug_assertions) {
    dotenv().ok();
  }

  let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
  PgConnection::establish(&database_url).expect(&format!("Error connecting to {}", database_url))
}

#[derive(Clone, Debug, PartialEq, Queryable)]
pub struct Ipv4Record {
  pub id: i32,
  pub ipv4_address: String,
  pub effective_flg: bool,
  // created_at: DateTime<Utc>,
  // updated_at: DateTime<Utc>,
  pub last_checked_at: DateTime<Utc>,
}

pub fn get_effective_records(conn: &PgConnection) -> Vec<Ipv4Record> {
  let ipv4_effective = ipv4_history
    .filter(effective_flg.eq(true))
    .load::<Ipv4Record>(conn)
    .expect("Error loading ipv4_history");
  return ipv4_effective;
}

#[allow(dead_code)]
pub fn exec_sql(file_path: &Path, conn: &PgConnection) -> Result<bool, String> {
  let display = file_path.display();
  let mut file = match File::open(&file_path) {
    Err(why) => panic!("couldn't open {}: {}", display, why),
    Ok(file) => file,
  };

  let mut s = String::new();
  match file.read_to_string(&mut s) {
    Err(why) => panic!("couldn't read {}: {}", display, why),
    Ok(_) => (),
  };

  match sql_query(s).execute(conn) {
    Err(why) => return Err(format!("execute query: \n{}\nfailed\n{}", display, why).into()),
    Ok(_) => return Ok(true),
  };
}
