#https://leetcode.com/problems/customers-who-never-order/description/

#left join

select C.Name as Customers from Customers as C left join Orders as O on C.Id = O.CustomerId where O.CustomerId IS NULL;
