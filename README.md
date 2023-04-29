# Synapsis BE Test
An online shop REST API built with Go, using JWT for authentication, PostgreSQL for the database, using bycrypt for hashing password, Cloudinary for image handling, and Midtrans for payment processing.

# Docker Web Image Link
https://hub.docker.com/r/nagetpresto/synapsis_test_web

# Docker Db Image Link
https://hub.docker.com/r/nagetpresto/synapsis_test_db

# Features
- Customer can view product list by product category.
- Customer can add products to the shopping cart.
- Customers can see a list of products in their shopping cart.
- Customer can delete products from the shopping cart.
- Customers can checkout and make payment transactions.
- User cannot register with same email.
- Email confirmation is sent after registration and successful orders.
- Customer cannot add product if they have not confirmed their email.
- Admin cannot add category and product if they have not confirmed their email.
- Stock is automatically deducted upon successful orders.

# Database Design
## User Table:
- Columns: 
	- ID (int, PK),
	- Name (string),
	- Email (string),
	- Password (string),
	- Image (string),
	- Status (string),
	- IsConfirmed (boolean),
	- ConfirmCode (string)
- Relationships:
	- One-to-Many with Transaction: A user can have multiple transactions.

## Category Table:
- Columns: 
	- ID (int, PK),
	- Name (string),
	- Image (string)
- Relationships:
	- One-to-Many with Product: A category can have multiple products.

## Product Table:
- Columns: 
	- ID (int, PK),
	- CategoryID (int, FK),
	- Name (string),
	- Stock (int),
	- Price (int),
	- Description (string),
	- Image (string)
- Relationships:
	- Many-to-One with Category: A product belongs to a category.
	- One-to-Many with Cart: A product can be present in multiple carts.

## Cart Table:
- Columns:
	- ID (int, PK),
	- UserID (int, FK),
	- ProductID (int, FK),
	- Qty (int),
	- Amount (int),
	- TransactionID (int, FK)
- Relationships:
	- Many-to-One with User: A cart belongs to a user.
	- Many-to-One with Product: A cart contains a product.
	- Many-to-One with Transaction: A cart is associated with a transaction.

## Transaction Table:
- Columns:
	- ID (int, PK),
	- UserID (int, FK),
	- Name (string),
	- Address (string),
	- PostalCode (string),
	- Phone (string),
	- Day (string),
	- Date (string),
	- Status (string),
	- TotalAmount (int)
- Relationships:
	- Many-to-One with User: A transaction belongs to a user.
	- One-to-Many with Cart: A transaction can have multiple carts.

# API Documentation
## User
https://documenter.getpostman.com/view/26087314/2s93eSYF6F
## Category
https://documenter.getpostman.com/view/26087314/2s93eSYFF7
## Product
https://documenter.getpostman.com/view/26087314/2s93eSYa3H
## Cart
https://documenter.getpostman.com/view/26087314/2s93eSYa7m
## Transaction
https://documenter.getpostman.com/view/26087314/2s93eSYaGe

# Environment
- SERVER_KEY=SB-Mid-server-R5p1DsRrkwGtvUm1trR2_yOb
- CLIENT_KEY=SB-Mid-client-07Igqhe3u-n4OVVO

- CLOUD_NAME=dpvugaeq1
- CLOUD_FOLDER=Synapsis Test
- API_KEY=823725867286355
- API_SECRET=fSWh1G7esyFPamIVVJU9dJfM3vU

- EMAIL_SYSTEM=bilqist1234@gmail.com
- PASSWORD_SYSTEM=obugcqpocntgtbub
- CONFIRM_URL=http://localhost:3000/cofirm-email-status/

- DB_HOST= localhost
- DB_USER=postgres
- DB_PASSWORD=098765
- DB_NAME=synapsis_be_test
- DB_PORT=5432

- PORT=3030

# How to Access Docker Image
# 1. Pull Web and Database Image
$ docker pull nagetpresto/synapsis_test_web
$ docker pull nagetpresto/synapsis_test_db

# 2. Run Database Container first
$ docker run -p 5432:5432 -e POSTGRES_USER="postgres" -e POSTGRES_PASSWORD="098765" -e POSTGRES_DB="synapsis_test" -e DB_PORT="5432"  nagetpresto/synapsis_test_db

## Check Database Created in the container
$ docker exec -it <db-container-name> psql -U postgres -l
output: <db-name>

## Check Database container host
$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <db-container-name>
output: <db-host>

# 3. Run Web Container
$ docker run -p 3030:3030 -e SERVER_KEY="SB-Mid-server-R5p1DsRrkwGtvUm1trR2_yOb" -e CLIENT_KEY="SB-Mid-client-07Igqhe3u-n4OVVO" -e CLOUD_NAME="dpvugaeq1" -e CLOUD_FOLDER="Synapsis Test" -e API_KEY="823725867286355" -e API_SECRET="fSWh1G7esyFPamIVVJU9dJfM3vU" -e EMAIL_SYSTEM="bilqist1234@gmail.com" -e PASSWORD_SYSTEM="obugcqpocntgtbub" -e CONFIRM_URL="http://localhost:3000/cofirm-email-status/" -e DB_HOST="<db-host>" -e DB_USER="postgres" -e DB_PASSWORD="098765" -e DB_NAME="<db-name>" -e DB_PORT="5432" -e PORT="3030" nagetpresto/synapsis_test_web

# 4. Test API using PORT: 3030



<!-- # docker pull nagetpresto/synapsis_test_web
# docker pull nagetpresto/synapsis_test_db
# docker exec -it heuristic_diffie psql -U postgres -l
# docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' heuristic_diffie
# host.docker.internal -->