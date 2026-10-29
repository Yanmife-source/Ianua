# Ianua

Ianua (Latin for "door") is a reverse proxy and web application firewall written in Go. It sits in front of a web application, inspects each HTTP request, and blocks or logs those that match known attack patterns, such as SQL injection, cross-site scripting and path traversal.

It was built as a learning project to understand how WAFs work from both sides: writing the rules, then trying to bypass them. It is not a replacement for production WAFs like Cloudflare or ModSecurity. The README documents the bypasses I found and the limits of regex-based detection.

**Features (v0.1.0):** configurable rules (YAML), block and log-only modes, JSON logs, per-IP rate limiting, tested against DVWA.
