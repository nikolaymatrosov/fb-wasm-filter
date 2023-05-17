use std::os::raw::c_char;
use std::slice;

// Import pure and fast JSON library written in Rust
use serde_json::json;
use serde_json::Value;

#[no_mangle]
pub extern "C" fn level_mapper(tag: *const c_char, tag_len: u32, time_sec: u32, time_nsec: u32, record: *const c_char, record_len: u32) -> *const u8 {
    let slice_record: &[u8] = unsafe { slice::from_raw_parts(record as *const u8, record_len as usize) };

    let mut v: Value = serde_json::from_slice(slice_record).unwrap();

    let level = &v["level"];
    let mapped_level: &str;
    match level.as_str() {
        Some("0") => mapped_level = "FATAL",
        Some("1") | Some("2") | Some("3") => mapped_level = "ERROR",
        Some("4") | Some("5") => mapped_level = "WARN",
        Some("6") => mapped_level = "INFO",
        Some("7") => mapped_level = "DEBUG",
        _ => mapped_level = "LEVEL_UNSPECIFIED"
    }

    v["level"] = serde_json::Value::String(mapped_level.to_string());

    let buf: String = v.to_string();
    buf.as_ptr()
}