-- Second Highest Salary
-- 
-- Write a SQL query to get the second highest salary from the Employee table.
-- 
-- +----+--------+
-- | Id | Salary |
-- +----+--------+
-- | 1  | 100    |
-- | 2  | 200    |
-- | 3  | 300    |
-- +----+--------+
-- For example, given the above Employee table, the query should return 200 as the second
-- highest salary. If there is no second highest salary, then the query should return null.
-- 
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | 200                 |
-- +---------------------+

-- ---------
-- init data
-- ---------

DROP TABLE IF EXISTS `Employee`;
CREATE TABLE `Employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Salary` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
LOCK TABLES `Employee` WRITE;
INSERT INTO `Employee` VALUES (1,100),(2,100);
UNLOCK TABLES;

-- -----------
-- my solution
-- -----------

SELECT IFNULL (  
    (SELECT Salary   
        FROM Employee  
        GROUP BY Salary  
        ORDER BY Salary DESC  
        LIMIT 1,1
    ),NULL  
) AS SecondHighestSalary 
