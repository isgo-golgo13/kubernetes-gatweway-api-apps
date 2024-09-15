import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import path from 'path';
import { fileURLToPath } from 'url';
import dotenv from 'dotenv';

dotenv.config();  // Load environment variables from .env file

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const packageDefinition = protoLoader.loadSync(path.join(__dirname, 'proto/send.proto'), {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});

const sendProto = grpc.loadPackageDefinition(packageDefinition).send;

function send(call, callback) {
  const { id, data } = call.request;
  callback(null, { bytes_sent: data.length });
}

function sendWithTimeout(call, callback) {
  const { id, data, timeout } = call.request;

  setTimeout(() => {
    if (timeout > 0) {
      callback(null, { bytes_sent: data.length });
    } else {
      callback({
        code: grpc.status.DEADLINE_EXCEEDED,
        message: 'Timeout exceeded'
      });
    }
  }, timeout);
}

function sendAll(call) {
  call.on('data', (request) => {
    const { id, data } = request;
    call.write({ bytes_sent: data.length });
  });

  call.on('end', () => {
    call.end();
  });
}

const server = new grpc.Server();
server.addService(sendProto.SendService.service, { send, sendWithTimeout, sendAll });

const PORT = process.env.PORT || '50051';  // Default to 50051 if PORT is not defined

function startServer() {
  server.bindAsync(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
      if (err.code === 'EADDRINUSE') {
        console.error(`Port ${PORT} is already in use.`);
        process.exit(1);
      } else {
        console.error(`Failed to bind server: ${err.message}`);
        process.exit(1);
      }
    }
    console.log(`Server running at http://0.0.0.0:${PORT}`);
    server.start();
  });
}

startServer();

process.on('SIGINT', () => {
  console.log('Received SIGINT. Shutting down gracefully...');
  server.tryShutdown((err) => {
    if (err) {
      console.error('Error during server shutdown:', err);
    } else {
      console.log('Server shut down gracefully.');
    }
    process.exit(err ? 1 : 0);
  });
});
