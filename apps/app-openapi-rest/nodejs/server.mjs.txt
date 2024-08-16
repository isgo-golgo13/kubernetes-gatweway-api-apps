import { Hono } from 'hono';
import DataSequenceRepository from './repositories/dataSequenceRepository.mjs';
import OwnerRepository from './repositories/ownerRepository.mjs';
import { setupRoutes } from './routes/routes.mjs';
import { config } from './config.mjs';

let server;  // Reference to the server instance

// Initialize the application
const app = new Hono();

// Initialize repositories
const dataSequenceRepo = new DataSequenceRepository();
const ownerRepo = new OwnerRepository();

// Setup routes with the initialized repositories
setupRoutes(app, dataSequenceRepo, ownerRepo);

// Add a healthcheck endpoint for Kubernetes liveness probe
app.get('/healthcheck', (c) => {
  return c.json({ status: 'ok' }, 200);
});

// Start the server on the configured port
server = app.listen(config.appPort, () => {
  console.log(`Server running on port ${config.appPort}`);
});

// Handle SIGINT (Ctrl-C) for graceful shutdown
process.on('SIGINT', () => {
  console.log('SIGINT signal received: closing HTTP server');
  server.close(() => {
    console.log('HTTP server closed');
    process.exit(0);
  });
});
