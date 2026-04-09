# JWT Decoder

A simple CLI tool for decoding JSON Web Tokens (JWT).

## Features

- Decode JWTs and view header and payload
- Optional signature verification for HMAC-signed tokens (for example HS256)
- Decodes tokens with any algorithm in the header/payload display path
- Uses `github.com/golang-jwt/jwt/v5` for signature parsing and verification

## Usage

```bash
git clone https://github.com/tbeyer567/jwt-decoder.git
cd jwt-decoder
go build -o jwt-decoder
```

## Example

```bash
./jwt-decoder -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJzcnNiaXoiLCJhdWQiOiJodHRwczovL2FwaS5leGFtcGxlLmNvbSIsImV4cCI6MzI1MDM2ODAwMDAsImlhdCI6MTc1ODkwNjMwMn0.WfaBbm1mUEao6Li9JFjIVhim7Tc3RnmcT9yDO5N1eAc -s yZ6xG7FPG+LY943gjz9SLW4gGhoelfaExe2xRQEgV+c=
Header:
{
  "alg": "HS256",
  "typ": "JWT"
}
Payload:
{
  "aud": "https://api.example.com",
  "exp": 32503680000,
  "iat": 1758906302,
  "sub": "srsbiz"
}
Signature checks out
```

## Notes

- Signature verification currently supports HMAC methods only. Tokens that use asymmetric algorithms such as `RS256` are decoded for inspection, but not verified by this CLI yet.

## License

MIT