### Environment Variables
The project uses a `.env` file for configuration. Below are the required environment variables:

| Variable | Description                                           | Default Value                                                           |
| --- |-------------------------------------------------------|-------------------------------------------------------------------------|
| `PORT` | The port that the application listens on              | `:8080`                                                                 |
| `REDIS` | Redis connection string (host:port)                   | `localhost:6379`                                                        |
| `LIMIT` | Default request limit per client (for rate limiting)  | `1`                                                                     |
| `INTERVAL` | Time window to track requests (in seconds)            | `5`                                                                     |
| `BLOCK_INTERVAL` | Duration to block a client after exceeding the limit (in seconds) | `60`                                                                    |
| `TOKEN_CONFIG` | JSON specifying custom rate limits for API tokens     | [{"TOKEN":"sub-token-1", "LIMIT":1,"INTERVAL":15, "BLOCK_INTERVAL":60}] |
