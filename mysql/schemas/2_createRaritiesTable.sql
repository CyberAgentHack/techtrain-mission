drop table if exists `rarities`;

/*
rarities, which is a master table, keeps rarity information for characters.
 - ID is PRIMARY KEY
 - name is NOT NULL
   it's prepared for expansion for future work
   if no needs, this column will be deleted
 - weight is NOT NULL
   this column defines ratio
*/

create table `rarities` (
  `id` int,
  `name` int NOT NULL,
  `weight` int NOT NULL,
  PRIMARY KEY (`id`)
);