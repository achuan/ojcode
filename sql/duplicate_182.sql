#https://leetcode.com/problems/duplicate-emails/description/

#hits:https://leetcode.com/problems/duplicate-emails/solution/


#mysql:having
select email from person group by email having count(email) > 1;

#mysql:where and temp table

select email from (select count(*)as num, email from person group by email )as a where num > 1;
