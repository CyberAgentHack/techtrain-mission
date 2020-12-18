drop table if exists usercharacters;

/*
usercharacters, which is master table, keeps information which each user gets.
 - id is PRIMARY KEY
 - user_id is foreign key and NOT NULL
   To delete and update the related data is NOT ALLOWED.
 - character_id is foreign key and NOT NULL
   it's same as above
 - ts is NOT NULL
   It's logged when the user gets the character.
*/

create table `usercharacters` (
  `id` int,
  `user_id` int NOT NULL,
  `character_id` int NOT NULL,
  `ts` timestamp NOT NULL, 
  PRIMARY KEY (`id`),
  foreign key (`user_id`) references `users(id)`
    on delete RESTRICT
    on update RESTRICT,
  foreign key (`character_id`) references `character(id)`
    on delete RESTRICT
    on update RESTRICT
);