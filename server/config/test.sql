-- SELECT * from ek_server_users;

-- select DATE_FORMAT(created_at,"%Y-%m-%d") date,count(*) count from (
--     select * from ek_server_articles 
-- ) as table1 GROUP by date;


-- SELECT @s :=@s + 1 as `index`, DATE(DATE_SUB(CURRENT_DATE, INTERVAL @s WEEK)) AS `date`
-- FROM mysql.help_topic,(SELECT @s := -1) temp
-- WHERE @s < 16
-- ORDER BY `date` ASC 

select date,count,title,id from
(
    SELECT DATE_FORMAT(created_at,'%Y-%m-%d') date,count(*) count
    FROM `ek_server_articles`  table1 
    WHERE (`created_at` BETWEEN '2020-07-02 16:40:47 +0800 CST' AND '2020-07-08 09:31:54 +0800 CST') 
    GROUP BY date ORDER BY date desc
) t1
right join 
(
    select title,DATE_FORMAT(created_at,'%Y-%m-%d') date2,id from `ek_server_articles` 
)t2
ON t2.date2 = t1.date
limit 5