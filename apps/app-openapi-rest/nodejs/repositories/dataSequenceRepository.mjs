import { createClient } from 'redis';

class DataSequenceRepository {
  constructor() {
    this.redisClient = createClient({
      url: `redis://${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`,
    });
    this.redisClient.on('error', (err) => console.error('Redis Client Error', err));
    this.redisClient.connect();
  }

  async createDataSequence(dataSequence) {
    const { ID, Data, DataOffsetStart, DataOffsetEnd, ResequenceCount, OwnerID, Timestamp, PurgeExpiry } = dataSequence;

    // Store the DataSequence in Redis as a hash
    await this.redisClient.hSet(`data_sequence:${ID}`, {
      ID,
      Data: Buffer.from(Data).toString('base64'), // Encode the byte array as a base64 string
      DataOffsetStart,
      DataOffsetEnd,
      ResequenceCount,
      OwnerID,
      Timestamp,
      PurgeExpiry,
    });

    // Add the DataSequence ID to the owner's list of sequences
    await this.ownerRepo.addDataSequenceToOwner(OwnerID, ID);
  }

  async getDataSequenceById(id) {
    const dataSequence = await this.redisClient.hGetAll(`data_sequence:${id}`);

    if (dataSequence && Object.keys(dataSequence).length > 0) {
      dataSequence.Data = Buffer.from(dataSequence.Data, 'base64'); // Decode the base64 Data back to binary
      return dataSequence;
    }
    return null;
  }

  async getAllDataSequences() {
    const keys = await this.redisClient.keys('data_sequence:*');
    const sequences = [];

    for (const key of keys) {
      const dataSequence = await this.redisClient.hGetAll(key);
      if (dataSequence && Object.keys(dataSequence).length > 0) {
        dataSequence.Data = Buffer.from(dataSequence.Data, 'base64'); // Decode the base64 Data back to binary
        sequences.push(dataSequence);
      }
    }

    return sequences;
  }
}

export default DataSequenceRepository;
