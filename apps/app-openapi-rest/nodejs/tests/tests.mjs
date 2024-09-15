import { describe, it } from 'mocha';
import chai from 'chai';
import chaiHttp from 'chai-http';
import { Hono } from 'hono';
import DataSequenceRepository from '../src/repositories/dataSequenceRepository.mjs';
import OwnerRepository from '../src/repositories/ownerRepository.mjs';
import { setupRoutes } from '../src/routes/routes.mjs';

chai.use(chaiHttp);
const { expect } = chai;

// Setup test app
const app = new Hono();
const dataSequenceRepo = new DataSequenceRepository();
const ownerRepo = new OwnerRepository();
setupRoutes(app, dataSequenceRepo, ownerRepo);

// Define test cases
describe('DataSequence API Tests', () => {
  it('should create a new DataSequence', async () => {
    const res = await chai.request(app.requestListener)
      .post('/api/sequences')
      .send({
        ID: '123',
        Data: Buffer.alloc(200).toString('base64'),
        DataOffsetStart: 0,
        DataOffsetEnd: 199,
        ResequenceCount: 0,
        OwnerID: 'owner123',
        Timestamp: new Date().toISOString(),
        PurgeExpiry: new Date(Date.now() + 86400000).toISOString(),
      });

    expect(res).to.have.status(201);
    expect(res.body).to.have.property('message', 'DataSequence created successfully');
  });

  it('should retrieve an existing DataSequence by ID', async () => {
    const res = await chai.request(app.requestListener)
      .get('/api/sequences/123');

    expect(res).to.have.status(200);
    expect(res.body).to.have.property('ID', '123');
  });
});
