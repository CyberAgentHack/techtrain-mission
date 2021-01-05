drop table if exists `usercharacters`;

/*
usercharacters, which is transaction table, keeps information which each user has.
 - id is PRIMARY KEY
 - user_id is foreign key and NOT NULL
   To delete and update the related data is NOT ALLOWED.
 - character_id is foreign key and NOT NULL
   it's same as above
 - possessions is NOT NULL
   It shows how many its characters the user has.
*/

create table `usercharacters` (
  `id` int,
  `user_id` int NOT NULL,
  `character_id` int NOT NULL,
  `possessions` int NOT NULL, 
  PRIMARY KEY (`id`),
  foreign key (`user_id`) references `users` (`id`)
    on delete RESTRICT
    on update CASCADE,
  foreign key (`character_id`) references `characters` (`id`)
    on delete RESTRICT
    on update CASCADE
);