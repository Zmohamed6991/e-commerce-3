# Project Title: Sole Select

**Description**: Contributed and collaborated with project managers and developers on an e-commerce project called Sole Select - a sneaker trading website.

**Backend**:
  - Go 
  - Gin (framework)
    
**Database**: 
 - PostgreSQL (GORM for interaction)
   
**Backend Deployed on**:
 - Render

# Tools & Project Management:

**Version Control**: 
 - GitHub
   
**Project Management**:
 - Jira: We used Jira to implement an agile methodology to assign and manage tasks with the product manager and other developers. We used Jira to plan a 4-week sprint to complete the e-commerce project and updated Jira daily after stand-ups.
 - Slack: We used Slack as a communication platform, by communicating in huddles and channels to share findings, progress and ideas.

**Key features**:
 -  User authentication using JWT-based authentication for buyer and seller.
 -  Users can place orders, view order history, update their cart and delete a product
 -  CRUD APIs.

# My role and contribution:
**Sprint 1**:
At the start of the project, my team and I set up our development environments, including connecting to the PostgreSQL database. My main responsibility was to develop the backend functionality for the shopping cart feature using GORM for communication with the PostgreSQL database and Gin to efficiently handle HTTP requests and routes. 
In the first sprint, we focused on add to cart method, this was done in collaboration with a colleague. We utilised our boot camp training and external documentation to implement this feature.
During the Sprint 1 retrospective, our team discussed areas for improvement, identified issues like unclear project direction, and proposed solutions for a more structured workflow.
![slack communication](https://github.com/user-attachments/assets/6f619a68-a03d-4d33-9ee3-9d3b053399df)
![sprint 1 retrospective](https://github.com/user-attachments/assets/d9d3ce0e-2b86-4ed6-851c-a5801582fc5d)

**Sprint 2**:
Having completed the "add-to-cart" method, I was assigned to work on the "edit cart" method, which allows users to modify product quantities in their shopping cart. I submitted a pull request for the feature, which was reviewed and successfully merged by the project manager.
In the sprint retrospective, our team shared the challenges we encountered, such as handling dynamic product quantities, and how each member resolved their respective blockers. I also contributed updates in Jira and planned for upcoming tasks.
![Sprint 2](https://github.com/user-attachments/assets/ae80b7f5-911c-40d8-aeb0-bf8f5d102cb5)


**Sprint 3**:
For Sprint 3, I was responsible for developing the "view order history" functionality, enabling users to view their past orders. This involved retrieving data from the database and displaying it in the backend, which presented a new learning opportunity for me in data retrieval and from Postgres database interaction.
During the retrospective, I raised blockers related to complex queries and sought guidance from the project manager, who provided valuable input on database calls. The sprint ended with most team members submitting pull requests that were reviewed and merged.
![Sprint 2 - add to cart](https://github.com/user-attachments/assets/8da39987-2f7b-46b9-b1bd-500e96d172af)

**Sprint 4**:
The final sprint focused on wrapping up outstanding tasks. I was tasked with deploying the backend project and database on Render (link to deployment: https://e-commerce-3-5xee.onrender.com). 
We also conducted API testing using Postman to ensure the proper functionality of our endpoints.

 
# Challenges and solutions:
**Challenge 1**:
In the initial stages, the internship's structure was unclear, and this confusion affected our workflow. To resolve this, I worked closely with the product manager, who reassigned tasks and provided clearer directions for the project. Regular communication helped streamline the development process.
![change of plans 2](https://github.com/user-attachments/assets/44be29be-4bf9-4e75-9450-4a8e67b2d6d9)
![change of plans 1](https://github.com/user-attachments/assets/4f4069b5-4954-4fd5-8361-9fbd5e955817)

**Challenge 2**:
As a beginner fresh from a three-month boot camp, the fast-paced environment of completing a project in just four weeks felt overwhelming. I had to quickly get up to speed with using PostgreSQL and GORM for database interactions. To overcome this, I sought 1-on-1 guidance from the project manager, who provided mentorship and support in debugging errors and resolving technical challenges.

# Links:
 - [Live Demo](https://drive.google.com/file/d/1glPezV347OwBZpPMLP0LneaZBgpKRN6c/view?usp=drive_link)
 - [GitHub Repository](https://github.com/Zmohamed6991/e-commerce-3)


# Full Project Documentation

# Setup Steps:
  - Go installed (version 1.16 or higher): https://golang.org/dl/
  - Gin framework: Install using the terminal:

  - go get -u github.com/gin-gonic/gin

  - GORM ORM (for database interaction): Install using the terminal
  - 	go get -u gorm.io/gorm
    
  - PostgreSQL: Download and install PostgreSQL to manage the database. You can also install pgAdmin 4 to manage the database via GUI - https://www.postgresql.org/download/
    
- Postman: Download Postman to test API endpoints: https://www.postman.com/downloads/
Clone the repository: git clone Mariana-consultancy/e-commerce-3


# Running the Application:
  - To start the application use the following:
    - Go run ./cmd/main.go

**The application will start on the web server at https://localhost:8080**


# Testing Endpoints using Postman:

**Create Seller**
- Create a seller to manage the e-commerce
- Method: POST
- URL: https://localhost:8080/seller/create
- Request body example:
  
		  {
		  "first_name": "Sam",
		  "last_name": "Sammy",
		  "password": "sammy96",
		  "date_of_birth": "30/12/96",
		  "email": "sam96@gmail.com",
		  "phone": "079234567812",
		  "address": "sam's road",
		  "store_name": "Select Sole",
		  "store_category": "sneaker store"
		}

**Success response with status code 200 OK.**

      {
          "data": null,
          "errors": null,
          "message": "Seller created",
          "status": "OK",
          "timestamp": "2024-09-26 18:06:23"
      }

**Error responses**:
**500 Internal Server Error**:
- If the seller is not created in the database.
- If the password is not hashed.

---------

**Sign In Seller**
- Seller logins in and the JWT access token is generated.
- Method: POST
- URL: https://localhost:8080/seller/login
- Request Body example:
  
		{
		    "email": "sam96@gmail.com",
		    "password": "sammy96"
		}

**Success response with status code 200 OK.**

		{
		    "data": {
		        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTgwNTMsInVzZXJfZW1haWwiOiJzYW05NkBnbWFpbC5jb20ifQ.3BBAwUk2_yHG7r7JTQ9tebvjDDUUQC0xvOoWhzoUR3M",
		        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTgwNTMsInN1YiI6MX0.-gFJrw-gbhb2l4HGcVJKktHTncGvubb91AWXPjb7UIM",
		        "seller": {
		            "ID": 1,
		            "CreatedAt": "2024-09-26T18:06:23.145474+01:00",
		            "UpdatedAt": "2024-09-26T18:06:23.145474+01:00",
		            "DeletedAt": null,
		            "first_name": "Sam",
		            "last_name": "Sammy",
		            "password": "$2a$14$dw1FVk99wRZ2WXlol3FnZOWhwmQR0giObtdSnivZEQNH4Q5YuJ8Yq",
		            "date_of_birth": "30/12/96",
		            "email": "sam96@gmail.com",
		            "phone": "079234567812",
		            "address": "sam's road",
		            "store_name": "Select Sole",
		            "store_category": "sneaker store"
		        }
		    },
		    "errors": null,
		    "message": "Login successful",
		    "status": "OK",
		    "timestamp": "2024-09-26 18:27:33"
		}

**Error responses**:
**404 Not Found**: 
- If the email does not exist in the database.

**400 Bad Request**: 
- If there is an invalid request
- If the email and password are left empty.
- If the email or password is invalid.

**500 Internal Server Error**:
- If there is an error generating the access token.
- If there is an error generating the refresh token.

---------

**Add a Product (Seller)**:

- The seller can add product
- Method: POST
- URL: https://localhost:8080/seller/product/create
- Use Access Token to authorise adding a product.
 
- Request Body example: 

		{
		    "name": "Jordan 1",
		    "price": 159.99,
		    "image_url": "https://unsplash.com/photos/a-box-with-a-pair-of-shoes-inside-of-it-XFV51iCfdrw",
		    "quantity": 5,
		    "description": "Air Jordan 1, Black and Pink"
		}

**Success response with status code 200 OK and message product has been created.**

		{
		    "data": null,
		    "errors": null,
		    "message": "product created successfully",
		    "status": "Created",
		    "timestamp": "2024-09-26 18:50:37"
		}

**Error response**:

**401 Unauthorised**: 

- If the token is valid or not inputted
- If there is an invalid request
  
**500 Internal Server Error**:
  
- If the product is not created in the database.
 
--------

**Sign Up User**

- Create a user account 
- Method: PUT
- URL: https://localhost:8080/user/create
- Request Body example:
  
		{
		    "first_name": "Zagel",
		    "last_name": "Mohamed",
		    "password": "zmohamed96",
		    "date_of_birth": "21/11/96",
		    "email": "zmohamed96@gmail.com",
		    "phone": "07912312312",
		    "address": "zee road"
		}

**Success response with status code 200 OK**

		{
		    "data": null,
		    "errors": null,
		    "message": "User created",
		    "status": "OK",
		    "timestamp": "2024-09-26 19:07:48"
		}




**Error response*:
**400 Bad Request**: 

- If there is an invalid request in creating the user.
  
**500 Internal Server Error**:

-  If there is an error in hashing the password in the database.
  
- If the user is not created in the database.

------------

**User Login**

- The user can log in and a JWT access token is generated.
- Method: POST
- URL: https://localhost:8080/user/login
- Request Body Example:
  
		{
		    "email": "zmohamed96@gmail.com",
		    "password": "zmohamed96"
		}

  
**Success response with status code 200 OK**

		{
		    "data": {
		        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NjEwMDgsInVzZXJfZW1haWwiOiJ6bW9oYW1lZDk2QGdtYWlsLmNvbSJ9.oSAWCrrTVpNq7sE8w0QZm1YiR5xAo2U5mHD1SRAggeM",
		        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NjEwMDgsInN1YiI6MX0.epRC6RTgvVZqaWowmQtXt3Y3yvpsR7evBIHSt1oKAsI",
		        "user": {
		            "ID": 1,
		            "CreatedAt": "2024-09-26T19:07:48.909972+01:00",
		            "UpdatedAt": "2024-09-26T19:07:48.909972+01:00",
		            "DeletedAt": null,
		            "first_name": "Zagel",
		            "last_name": "Mohamed",
		            "password": "$2a$14$G9kDjXAhvr/XP/jFuB5fHeH60oOhe5SQiDKEpTfL.eMRgCpEMgfYS",
		            "date_of_birth": "21/11/96",
		            "email": "zmohamed96@gmail.com",
		            "phone": "07912312312",
		            "address": "zee road"
		        }
		    },
		    "errors": null,
		    "message": "Login successful",
		    "status": "OK",
		    "timestamp": "2024-09-26 19:16:48"
		}


**Error response**:
**400 Bad Request**:

- Invalid request when logging in.
- If the email and password are left blank.
- If the email or password is invalid.
  
**404 Not Found**:

- If the email does not exist in the database.
  
**500 Internal Server Error**:

- If there is an error generating the access token.
- If there is an error generating the refresh token.

--------------

**Add to cart**

- Users can add product(s) to the cart.
- Method: POST
- URL: https://localhost:8080/user/cart/add
- Use the access token to authorise adding to the cart
- Request Body example:
  
		{
		    "product_id": 1,
		    "quantity": 2
		}


**Success response with status code 200 OK.** 


	{
	    "data": {
	        "ID": 1,
	        "CreatedAt": "2024-09-26T18:50:37.83469+01:00",
	        "UpdatedAt": "2024-09-26T18:50:37.83469+01:00",
	        "DeletedAt": null,
	        "seller_id": 1,
	        "name": "Jordan 1",
	        "price": 159.99,
	        "image_url": "https://unsplash.com/photos/a-box-with-a-pair-of-shoes-inside-of-it-XFV51iCfdrw",
	        "quantity": 5,
	        "description": "Air Jordan 1, Black and Pink",
	        "status": false,
	        "orders": null
	    },
	    "errors": null,
	    "message": "product added to cart",
	    "status": "OK",
	    "timestamp": "2024-09-26 19:31:46"
	}


**Error response**:
**401 Unauthorised**:
	- If the token is invalid.
	- If the request is invalid.
**500 Internal Server Error**:
	- If the product is not found in the database.
	- If there is an error adding a product to the cart.
**400 Bad Request**:
	- If the product is out of stock.

----------

**Edit Cart (User)**

- The user can edit product quantity
- Method: PUT 
- URL: https://localhost:8080/user/cart/edit
- Use the access token to authorise 
- Request Body Example:
  
 
		{
		    "product_id": 1,
		    "quantity": 1
		}


**Success response with status code 200 OK.**

	{
	    "data": null,
	    "errors": null,
	    "message": "product successfully added",
	    "status": "OK",
	    "timestamp": "2024-09-26 20:12:20"
	}

 
**Error response**:
**401 Unauthorised**:

- If the token is invalid, the user is unauthorised. 
- If the request is invalid.
  
**404 Status Not Found**:
  
- If the cart is not found in the database.
  
**500 Internal Server Error**:
  
- If the product is not found in the database by ID. 
- If there is an error updating the product quantity in the cart table in the database.
  
**400 Bad Request**:

- If the requested quantity exceeds the quantity available. 


