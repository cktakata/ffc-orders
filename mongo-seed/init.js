db = db.getSiblingDB('ffc_database'); // Replace 'your_database' with your database name
db.createCollection('orders'); // Replace 'your_collection' with your collection name
db.your_collection.insertMany([
  { store: 'test', name: 'Alice', date: new Date('2024-08-19T20:00:00Z'), value: 35, prevHash: 0, Hash: 0 }
]);
