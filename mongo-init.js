db = db.getSiblingDB('admin');
db.auth(
    process.env.MONGO_INITDB_ROOT_USERNAME,
    process.env.MONGO_INITDB_ROOT_PASSWORD,
);

db = db.getSiblingDB('inventorydb');
db.createUser({
  user: process.env.MONGO_DB_UN,
  pwd: process.env.MONGO_DB_PW,
  roles: [{
      role: 'readWrite',
      db:process.env.MONGO_INITDB_DATABASE,
  }]
});

db.createCollection('inventory');
db.createCollection('items');
db.createCollection('item_type');

