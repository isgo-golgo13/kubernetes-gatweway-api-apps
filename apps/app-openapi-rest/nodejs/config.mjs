import dotenv from 'dotenv';

// Load environment variables from a .env file if present
dotenv.config();

// Export configuration variables
export const config = {
  redisHost: process.env.REDIS_HOST || 'localhost',
  redisPort: process.env.REDIS_PORT || 6379,
  appPort: process.env.APP_PORT || 3000,
};
