-- Combine Two Tables
-- 
-- Table: Person
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | PersonId    | int     |
-- | FirstName   | varchar |
-- | LastName    | varchar |
-- +-------------+---------+
-- PersonId is the primary key column for this table.
-- Table: Address
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | AddressId   | int     |
-- | PersonId    | int     |
-- | City        | varchar |
-- | State       | varchar |
-- +-------------+---------+
-- AddressId is the primary key column for this table.
-- 
-- Write a SQL query for a report that provides the following information for each
-- person in the Person table, regardless if there is an address for each of those people:
-- FirstName, LastName, City, State

-- ---------
-- init data
-- ---------

DROP TABLE IF EXISTS `Address`;
CREATE TABLE `Address` (
  `AddressId` int(11) NOT NULL AUTO_INCREMENT,
  `PersonId` int(11) DEFAULT NULL,
  `City` varchar(20) DEFAULT NULL,
  `State` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`AddressId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
LOCK TABLES `Address` WRITE;
INSERT INTO `Address` VALUES (1,1,'LA','Cal'),(2,2,'BJ','BJ');
UNLOCK TABLES;


DROP TABLE IF EXISTS `Person`;
CREATE TABLE `Person` (
  `PersonId` int(11) NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(10) DEFAULT NULL,
  `LastName` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`PersonId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
LOCK TABLES `Person` WRITE;
INSERT INTO `Person` VALUES (1,'kobe','bryant'),(2,'zz','huang'),(3,'hah','ha');
UNLOCK TABLES;

-- -----------
-- my solution
-- -----------

SELECT FirstName, LastName, City, State FROM Person
LEFT JOIN Address ON Person.PersonId = Address.PersonId;
