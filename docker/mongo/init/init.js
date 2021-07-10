var users = [
  {
    user: "user_management_root",
    pwd: "admin",
    roles: [
      {
        role: "dbOwner",
        db: "user_management"
      }
    ]
  },
  {
    user: "user_management_client",
    pwd: "client",
    roles: [
      {
        role: "readWrite",
        db: "user_management"
      }
    ]
  }
];

for (var i = 0, length = users.length; i < length; ++i) {
  db.createUser(users[i]);
}

db = new Mongo().getDB("user_management");

db.createCollection('user');

db.user.insert({
  'e_mail': 'test@domain.com',
  'first_name': 'test',
  'last_name': 'test'
});