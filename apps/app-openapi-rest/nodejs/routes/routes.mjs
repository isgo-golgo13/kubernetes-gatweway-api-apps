export function setupRoutes(app, dataSequenceRepo, ownerRepo) {
    app.post('/api/sequences', async (c) => {
      const body = await c.req.json();
      const { ID, Data, DataOffsetStart, DataOffsetEnd, ResequenceCount, OwnerID, Timestamp, PurgeExpiry } = body;
  
      // Input validation can be done here, such as checking for the byte length of Data
      if (!ID || !Data || Data.length < 200 || !DataOffsetStart || !DataOffsetEnd || !ResequenceCount || !OwnerID || !Timestamp || !PurgeExpiry) {
        return c.json({ error: 'Invalid input data' }, 400);
      }
  
      await dataSequenceRepo.createDataSequence(body);
      return c.json({ message: 'DataSequence created successfully' }, 201);
    });
  
    app.get('/api/sequences/:id', async (c) => {
      const id = c.req.param('id');
      const dataSequence = await dataSequenceRepo.getDataSequenceById(id);
  
      if (!dataSequence) {
        return c.json({ error: 'DataSequence not found' }, 404);
      }
  
      return c.json(dataSequence);
    });
  
    app.get('/api/sequences', async (c) => {
      const sequences = await dataSequenceRepo.getAllDataSequences();
      return c.json(sequences);
    });
  
    app.get('/api/owners/:id', async (c) => {
      const id = c.req.param('id');
      const owner = await ownerRepo.getOwnerById(id);
  
      if (!owner) {
        return c.json({ error: 'Owner not found' }, 404);
      }
  
      return c.json(owner);
    });
  
    app.get('/api/owners', async (c) => {
      const owners = await ownerRepo.getAllOwners();
      return c.json(owners);
    });
  }
  