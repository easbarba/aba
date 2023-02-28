# Aba

Like postman/insomnia but CLI.

# Usage

# Methods Configuration

JSON configuration with objects describing methods

```json
[
  {
    "method": "get",
    "path": "$base_url/$version/$referee/1"
  },
  {
    "method": "post",
    "path": "$base_url/$version/$referee"
    "resource": "referee"
  }
]
```

# Input

JSON input to send to HTTP receiver.

# Environment

Shareable environment variables

```json
{
  "base_url": "localhost:5000/api",
  "referee_url": "referees"
}
```

# Collections

TODO

## GNU Guix

Load all system locall dependencies with `guix shell --pure`

## LICENSE

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
