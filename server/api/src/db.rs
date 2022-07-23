use chrono::{DateTime, Utc};

use diesel::pg::PgConnection;
use diesel::prelude::*;

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

#[derive(Queryable)]
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
