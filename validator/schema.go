package validator

// Schema : JSONSchema validation schema
var Schema = `{
  "properties": {
    "id": {"type": "string", "minLength": 1},
    "first_name": {"type": "string", "minLength": 1},
    "last_name": {"type": "string", "minLength": 1},
    "tax_id": {"type": "string", "pattern": "^[0-9]{3}[ -]*[0-9]{2}[ -]*[0-9]{4}$"},
    "address": {
      "type": "object",
      "properties": {
        "street": {"type": "string", "minLength": 1},
        "city": {"type": "string", "minLength": 1},
        "state": {"type": "string"},
        "postal_code": {"type": "string", "pattern": "^[A-Za-z0-9][A-Za-z0-9 ]{3,8}[A-Za-z0-9]$"},
        "country": {"type": "string", "enum": ["US", "GB", "CA", "AT"]}
      },
      "required": ["street", "city", "postal_code", "country"],
      "if": {
        "properties": {
          "country": { "enum": ["US", "CA"] }
        }
      },
      "then": {
        "required": ["state"]
      }
    }
  },
  "required": ["id", "first_name", "last_name", "tax_id", "address"]
}`
