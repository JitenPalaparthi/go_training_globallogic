# Task

- Product --> Name,Number,Category, Desc
- Customer --> Name, Number, Email, Address
- Orders --> OrderId, ProductID,CustomerID

- Create a restful service
- product/create
- product/get/:id
- product/update/:id
- product/delete/:id

- customer/create
- customer/get/:id
- customer/update/:id
- customer/delete/:id
- customer/orders/get/:customerid
- customer/orders/get/:customerid/:orderid/
- customer/orders/product/get/:customerid/:productid
  
- Order/create
- Order/get/:id
- Order/update/:id
- Order/delete/:id

Http based RESTFul service 
Use Postgres or Mysql as database
Write basic unit test for models

Use Kafa as a message broker upon adding products and purchase a product.
