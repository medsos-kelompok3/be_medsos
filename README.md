# be_medsos API

Medsos App is a simple social media application that allows users to create posts, comment on posts, edit comments, edit user data, and delete accounts. This repository contains the frontend code for the Medsos App.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [API](#api)
 
## Getting Started

### Prerequisites

- Go installed on your machine
- GORM
- Echo
- Postman
- Database (e.g., Mysql) installed and accessible

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/medsos-kelompok3/be_medsos.git
2. Change into the project directory:

   ```bash
    cd be_medsos
3. Install dependencies:

   ```bash
    go mod tidy

### API

### Overview

Circles provides a RESTful API for managing and retrieving media sosial. Below are the available endpoints:

### Authentication

- **Endpoint**: `/register`
  - **Method**: `POST`
  - **Description**: Register a new user.
  - **Request Body**:
    ```json
    {
      "username": "your_name",
      "email":"your_email",
      "address":"your_address",
      "password": "your_password"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User registered successfully"
    }
    ```

- **Endpoint**: `/login`
  - **Method**: `POST`
  - **Description**: Log in an existing user.
  - **Request Body**:
    ```json
    {
      "username": "your_name",
      "password": "your_password"
    }
    ```
  - **Response**:
    ```json
    {
      "token": "your_access_token"
    }
    ```

### User

- **Endpoint**: `/user/:id`
  - **Method**: `PUT`
  - **Description**: Update data from User.
  - **Request Body**:
    ```json
    {
      "bio": "your_bio",
      "avatar":"your_avatar",
      "address":"your_address"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User registered successfully"
    }
    ```
- **Endpoint**: `/user/:id`
  - **Method**: `DELETE`
  - **Description**: Delete data from User.
  - **Response**:
    ```json
    {
      "message": "User deleted successfully"
    }
    ```

### Posting

- **Endpoint**: `/posting`
  - **Method**: `POST`
  - **Description**: Posting Successfully created.
  - **Request Body**:
    ```json
    {
      "caption": "your_caption",
      "gambar_posting":"your_posting"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Posting successfully created"
    }
    ```
- **Endpoint**: `/posting/:id`
  - **Method**: `PUT`
  - **Description**: Posting Sucessfully updated.
  - **Request Body**:
    ```json
    {
      "caption": "your_caption",
      "gambar_posting":"your_posting"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Posting updated successfully"
    }
    ```
- **Endpoint**: `/posting`
  - **Method**: `GET`
  - **Description**: Posting Successfully read.
  - **Response**:
    ```json
    {
      "caption": "your_caption",
      "gambar_posting":"your_posting"
    }
    ```
- **Endpoint**: `/posting/:id`
  - **Method**: `DELETED`
  - **Description**: Posting Sucessfully deleted.
  - **Response**:
    ```json
    {
      "message": "Posting deleted successfully"
    }
    ```

### Comment

- **Endpoint**: `/comment`
  - **Method**: `POST`
  - **Description**: Comment Successfully created.
  - **Request Body**:
    ```json
    {
      "posting_id": 1,
      "isi_comment":"your_comment"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Comment successfully created"
    }
    ```
- **Endpoint**: `/posting/:id`
  - **Method**: `PUT`
  - **Description**: Comment Sucessfully updated.
  - **Request Body**:
    ```json
    {
      "isi_comment":"your_comment"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Posting updated successfully"
    }
    ```
- **Endpoint**: `/comment/:id`
  - **Method**: `DELETED`
  - **Description**: Comment Sucessfully deleted.
  - **Response**:
    ```json
    {
      "message": "Comment deleted successfully"
    }
    ```
