# Multi-Currency Exchange API

### Description

The Multi-Currency Exchange API is a web-based application that enables users to exchange USD, EUR, BTC and ETH cryptocurrencies.

### Technologies Used

- Golang
- GORM
- Gin Gonic
- Docker
- Postgres

### API endpoints

The following API endpoints are avaiable in the project:

- [http://localhost:8080/signup](http://localhost:8080/signup) - Create a new user
- [http://localhost:8080/login](http://localhost:8080/login) - Login user and retrieve api token secret in cookies
- [http://localhost:8080/user](http://localhost:8080/user) - List all users or create a new user
- [http://localhost:8080/user/1](http://localhost:8080/<pk>) - Retrieve, update or delete a user by ID
- [http://localhost:8080/exchange](http://localhost:8080/exchange) - Return the exchange value

### Models

| User             | type   |
| :--------------- | :----- |
| Username         | string |
| Password         | string |
| UserRequestCount | string |
| LastResetTime    | time   |

| Exchange | type    |
| :------- | :------ |
| From     | string  |
| To       | string  |
| Amount   | float64 |

| Log       | type   |
| :-------- | :----- |
| ID        | uint   |
| UserID    | uint   |
| Timestamp | time   |
| Method    | string |
| Path      | string |
| Status    | string |

### Avaiable Currencys

- USD
- EUR
- BTC
- ETH

<br>

# Challenge

### Objective

Using Go and any framework, your task is to build a currency conversion service that includes FIAT and cryptocurrencies.

### Brief

In this challenge, your assignment is to build a service that makes conversions between different currencies. You will connect to an external API to request currency data, log & store requests of your users, and rate limit requests based on specific criteria. Your service must support at least the following currency pairs:

USD
EUR
BTC
ETH

### Tasks

- Implement assignment using:

  - Language: **Go**
  - Framework: **any framework**

- We recommend using the Coinbase API for exchange rates:

  https://developers.coinbase.com/api/v2#get-exchange-rates

- Your service should be able to identify users. You may use any form of authentication that you think is suitable for the task (e.g., API keys, Username/Password)
- Your service needs to store each request, the date/time it was performed, its parameters and the response body
- Each user may perform 100 requests per workday (Monday-Friday) and 200 requests per day on weekends. After the quota is used up, you need to return an error message
- The service must accept the following parameters:
  - The source currency, the amount to be converted, and the final currency
  - e.g. `?from=BTC&to=USD&amount=999.20`
- Your service must return JSON in a structure you deem fit for the task
- BONUS: find a clever strategy to cache responses from the external currency API

### Evaluation Criteria

- **Go** best practices
- Show us your work through your commit history
- We're looking for you to produce working code, with enough room to demonstrate how to structure components in a small program
- Completeness: Did you complete the features?
- Correctness: Does the functionality act in sensible, thought-out ways?
- Maintainability: Is it written in a clean, maintainable way?
- Testing: Is the system adequately tested?

### CodeSubmit

Please organize, design, test, and document your code as if it were going into production - then push your changes to the master branch. After you have pushed your code, you may submit the assignment on the assignment page.

All the best and happy coding,

The Kajae Team
