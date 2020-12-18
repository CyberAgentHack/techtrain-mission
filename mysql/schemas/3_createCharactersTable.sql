drop table if exists `characters`;

/*
characters, which is a transaction table, keeps character information.
 - ID is PRIMARY KEY
 - name is NOT NULL
 - rarity_id is foreign key and NOT NULL
   because rarity is centrally managed by rarity.
*/

create table `characters` (
    `id` int AUTO_INCREMENT,
    `name` varchar(64) NOT NULL,
    `rarity_id` int NOT NULL,
    PRIMARY KEY (`id`),
    foreign key (`rarity_id`) references rarities(id)
      on delete RESTRICT
      on update CASCADE
) ;