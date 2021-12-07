-- Import File to Mysql
LOAD DATA INFILE '/var/lib/mysql-files/Test_Case_DataUser.csv' INTO TABLE data_belanja FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n' IGNORE 1 ROWS;
LOAD DATA INFILE '/var/lib/mysql-files/Test_Case_DataBelanja.csv' INTO TABLE data_belanja FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n' IGNORE 1 ROWS;

-- SQL Concate Two Table 
SELECT * FROM data_country dc ,data_belanja db WHERE dc.id = db.id_user 
 


-- SQL SELECT FROM MAX VALUE ONE ROW RESULT
SELECT * FROM data_belanja WHERE total_buy = (SELECT MAX(data_belanja.total_buy) FROM data_belanja) 

-- SQL SELECT FROM MAX VALUE ALL ROW RESULT
SELECT 
    s1.id, s1.id_user , s1.total_buy
FROM
    data_belanja s1
        LEFT JOIN
    data_belanja s2 ON s1.total_buy < s2.total_buy
WHERE
    s2.id IS NULL;


-- SQL Concate Two Table AND  SELECT FROM MAX VALUE ONE ROW RESULT
SELECT * FROM data_country dc ,data_belanja db WHERE dc.id = db.id_user AND total_buy = (SELECT MAX(data_belanja.total_buy) FROM data_belanja)

SELECT
  country,
  COUNT(country) AS `value_occurrence` 
FROM
  data_country
GROUP BY 
  country
ORDER BY 
  `value_occurrence` DESC
LIMIT 1;

SELECT * FROM data_country GROUP BY country ORDER BY  count(country) DESC;