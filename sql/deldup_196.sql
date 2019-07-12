//https://leetcode.com/problems/delete-duplicate-emails/description/

//delete, group by min
delete from Person where Id not in (select Min(ID) from (select * from person)tmp group by Email);
