## Hono.js Node.js with Redis OpenAPI REST Service (Source for Gateway API)

To install the app.

```shell
npm install
```

To run the app (non Docker version).

```shell
node run server.mjs
```


### App Configuration

The included `.env` and `config.mjs` work in cooperation to externalize the `server.mjs` configuration.

```javascript
import dotenv from 'dotenv';

// Load environment variables from a .env file if present
dotenv.config();

// Export configuration variables
export const config = {
  redisHost: process.env.REDIS_HOST || 'localhost',
  redisPort: process.env.REDIS_PORT || 6379,
  appPort: process.env.APP_PORT || 3000,
};
```

The corresponding `.env` files is as follows.

```
# .env file

# Redis configuration
REDIS_HOST=localhost
REDIS_PORT=6379

# Application configuration
APP_PORT=3000
```

This Node.js config could use the safer `dotenv-safe` module. To use this add the following to the package.json.

```shell
npm install dotenv-safe
```
