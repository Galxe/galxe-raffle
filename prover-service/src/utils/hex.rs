/// Parse a hex string into a 32-byte array
pub fn parse_hex_string(s: &str) -> Result<[u8; 32], String> {
    let s = s.strip_prefix("0x").unwrap_or(s);
    if s.len() != 64 {
        return Err("Randomness must be a 32-byte (64 character) hex string".to_string());
    }

    let bytes = hex::decode(s).map_err(|e| format!("Failed to decode hex string: {}", e))?;

    bytes
        .try_into()
        .map_err(|_| "Failed to convert to 32 byte array".to_string())
}
