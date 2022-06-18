use chrono::{DateTime, Utc};

use diesel::prelude::*;
use diesel::{insert_into, update};

mod schema;
use schema::ipv4_history;
use schema::ipv4_history::dsl::*;

#[derive(Queryable)]
pub struct Ipv4Record {
  pub id: i32,
  pub ipv4_address: String,
  effective_flg: bool,
  created_at: DateTime<Utc>,
  updated_at: DateTime<Utc>,
}

#[derive(Insertable, Debug)]
#[table_name = "ipv4_history"]
pub struct NewIpV4Record {
  ipv4_address: String,
  effective_flg: bool,
  created_at: DateTime<Utc>,
  updated_at: DateTime<Utc>,
}

pub fn get_effective_records(conn: &diesel::PgConnection) -> Vec<Ipv4Record> {
  let ipv4_effective = ipv4_history
    .filter(effective_flg.eq(true))
    .load::<Ipv4Record>(conn)
    .expect("Error loading ipv4_history");
  return ipv4_effective;
}

pub fn insert_record(conn: &diesel::PgConnection, address: impl Into<String>) -> Ipv4Record {
  let now = Utc::now();
  let new_ipv4_record = NewIpV4Record {
    ipv4_address: address.into(),
    effective_flg: true,
    created_at: now,
    updated_at: now,
  };
  insert_into(ipv4_history)
    .values(new_ipv4_record)
    .get_result(conn)
    .expect("Error saving record")
}

pub fn disable_record(conn: &diesel::PgConnection, target_id: &i32) -> Ipv4Record {
  update(ipv4_history.find(target_id))
    .set((effective_flg.eq(false), updated_at.eq(Utc::now())))
    .get_result::<Ipv4Record>(conn)
    .expect(&format!("Error Update Record {}", target_id))
}
