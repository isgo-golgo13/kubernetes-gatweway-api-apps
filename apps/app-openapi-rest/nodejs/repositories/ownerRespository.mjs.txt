import { createClient } from 'redis';

class OwnerRepository {
  constructor() {
    this.redisClient = createClient({
      url: `redis://${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`,
    });
    this.redisClient.on('error', (err) => console.error('Redis Client Error', err));
    this.redisClient.connect();
  }

  async createOwner(ownerID) {
    // Create an owner hash
    await this.redisClient.hSet(`owner:${ownerID}`, { ID: ownerID });
  }

  async addDataSequenceToOwner(ownerID, dataSequenceID) {
    // Ensure the owner exists
    const ownerExists = await this.redisClient.exists(`owner:${ownerID}`);
    if (!ownerExists) {
      await this.createOwner(ownerID);
    }

    // Add the DataSequence ID to the owner's list of sequences
    await this.redisClient.rPush(`owner:${ownerID}:sequences`, dataSequenceID);
  }

  async getOwnerById(ownerID) {
    const owner = await this.redisClient.hGetAll(`owner:${ownerID}`);
    if (!owner || Object.keys(owner).length === 0) {
      return null;
    }

    const dataSequences = await this.redisClient.lRange(`owner:${ownerID}:sequences`, 0, -1);
    owner.DataSequences = dataSequences;
    return owner;
  }

  async getAllOwners() {
    const keys = await this.redisClient.keys('owner:*:sequences');
    const owners = [];

    for (const key of keys) {
      const ownerID = key.split(':')[1]; // Extract the owner ID from the key
      const owner = await this.getOwnerById(ownerID);
      if (owner) {
        owners.push(owner);
      }
    }

    return owners;
  }
}

export default OwnerRepository;
