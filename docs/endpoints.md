## API endpoints

This document provides an overview of the available API endpoints, their methods, and usage.

### 1. Coffee Roasters

```bash
curl -X POST 'http://localhost:4000/api/v1/coffee-roasters' \
-H 'Content-Type: application/json' \
-d '{
  "name": "Example Roaster",
  "online_shop_url": "https://example.com"
}' | jq
```
