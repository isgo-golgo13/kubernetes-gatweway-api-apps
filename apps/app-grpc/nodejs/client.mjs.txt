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

// Load server address from .env or default to 'localhost:50051'
const SERVER_ADDRESS = process.env.SERVER_ADDRESS || 'localhost:50051';

const client = new sendProto.SendService(SERVER_ADDRESS, grpc.credentials.createInsecure());

function sendData(id, data) {
  client.send({ id, data }, (error, response) => {
    if (!error) {
      console.log('Bytes sent:', response.bytes_sent);
    } else {
      console.error(error);
    }
  });
}

function sendDataWithTimeout(id, data, timeout) {
  client.sendWithTimeout({ id, data, timeout }, (error, response) => {
    if (!error) {
      console.log('Bytes sent:', response.bytes_sent);
    } else {
      console.error(error);
    }
  });
}

function sendAllData(dataArray) {
  const call = client.sendAll();

  call.on('data', (response) => {
    console.log('Bytes sent:', response.bytes_sent);
  });

  call.on('end', () => {
    console.log('Stream ended.');
  });

  dataArray.forEach(data => {
    call.write(data);
  });

  call.end();
}

// Sending data
sendData(1, Buffer.from('Data: 000000000000000000000000000000000000000001'));
sendDataWithTimeout(2, Buffer.from('Data: 000000000000000000000000000000000000000001 timeout'), 1000);
sendAllData([{ id: 3, data: Buffer.from('Stream 1') }, { id: 4, data: Buffer.from('Stream 2') }]);
