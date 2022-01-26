-- https://leetcode.com/problems/customers-who-never-order/
SELECT c.Name as Customers
FROM Customers c
LEFT JOIN Orders o ON o.CustomerId = c.Id
WHERE true
GROUP BY c.ID
HAVING COUNT(o.Id) = 0