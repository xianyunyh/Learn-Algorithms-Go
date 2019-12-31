-- 175. 左查询
select  a.FirstName,a.LastName,b.City,b.State from Person as a  left join Address as b on a.PersonId = b.PersonId
-- 176 薪水第二的
select   (select distinct  salary from Employee order by salary desc limit 1,1 ) as SecondHighestSalary
-- 181.超过经理的用户 自连接
select a.Name as Employee  from Employee as a ,Employee as b where a.ManagerId =b.Id  and a.Salary > b.Salary
-- 182 重复的电子邮件
select email from Person group by email having count(email)  > 1
-- 183 从不订购的客户
select Name as   Customers from  Customers where Id not in (select CustomerId from Orders)