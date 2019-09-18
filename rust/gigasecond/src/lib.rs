use chrono::{DateTime, Duration, Utc};
use std::ops::Add;

pub fn after(time: DateTime<Utc>) -> DateTime<Utc> {
    let gigasecond = (10 as i64).pow(9);
    time.add(Duration::seconds(gigasecond))
}
