# API Reference & Error Handling

## Endpoints

### Get All Books
- **URL**: `/api/books`
- **Method**: `GET`
- **Response**: `200 OK` (JSON Array)

### Get Book By ID
- **URL**: `/api/books/{id}`
- **Method**: `GET`
- **Response**:
    - `200 OK` (JSON Object) if found.
    - `404 Not Found` (JSON Error) if not found.

## Error Handling Visual Anchor

> [!TIP]
> ⚓ **Visual Anchor:** The Request Flow Decision Tree

When a client requests a specific resource, the server makes a decision:

```mermaid
graph TD
    Client[Client Request] -->|GET /api/books/1| Router
    Router --> Handler[GetBookByID Handler]
    Handler -->|Parse ID| Store[In-Memory Store]
    Store -->|Found| Check{Exists?}
    Check -->|Yes| JSON[Return JSON Object]
    Check -->|No| Error[Return 404 Error]
    
    JSON -->|200 OK| Client
    Error -->|{"error": "not found"}| Client
```
