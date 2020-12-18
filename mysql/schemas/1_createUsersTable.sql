drop table if exists `users`;

/*
users, which is a transaction table, keeps user information.
 - ID is PRIMARY KEY
   it is assigned uniquely on calling /user/create.
 - name is NOT NULL
   because it needs on calling /user/create.
   the max length is 64 because it's enough, I think.
 - token is NOT NULL
   because it is assigned on creating user record.
*/

create table `users` (
  `id` int AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `token` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
);