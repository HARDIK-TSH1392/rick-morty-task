
# Rick and Morty API Project

This project is a full-stack application that integrates with the **Rick and Morty API**. It includes the following components:

- **Backend**: A REST API built using **Go** that fetches character details from the Rick and Morty API and stores data in **MongoDB**.
- **Frontend**: A **React (TypeScript)** application that allows users to search and view Rick and Morty character details.
- **Cloud Infrastructure**: **Terraform** scripts to provision MongoDB on **Google Cloud Platform (GCP)**.

### Table of Contents
- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)
- [Architecture Overview](#architecture-overview)

---

### Setup Instructions

#### 1. **Clone the Repository**

To get started, first clone the repository to your local machine:

```bash
git clone https://github.com/your-username/rick-and-morty-api-project.git
cd rick-and-morty-api-project
```

#### 2. **Setting Up the Backend**

1. **Navigate to the Backend Directory**:

   ```bash
   cd rick_and_morty_api
   ```

2. **Ensure Go is Installed**:

   You need **Go** installed on your machine. You can check if Go is installed by running:
   ```bash
   go version
   ```

   If Go is not installed, follow the [Go installation guide](https://golang.org/doc/install).

3. **Install Dependencies**:

   Install any dependencies for the Go backend:
   ```bash
   go mod tidy
   ```

4. **Run the Backend API**:

   Run the backend server using:
   ```bash
   go run main.go
   ```

   The backend API will be available at `http://localhost:8080`.

#### 3. **Setting Up the Frontend**

1. **Navigate to the Frontend Directory**:

   ```bash
   cd rick-and-morty-frontend
   ```

2. **Install Node.js Dependencies**:

   Ensure you have **Node.js** and **npm** installed. You can check the installation by running:
   ```bash
   node --version
   npm --version
   ```

   If not installed, you can install Node.js from [here](https://nodejs.org/).

   After that, install the required frontend dependencies:
   ```bash
   npm install
   ```

3. **Run the React Application**:

   Start the frontend application using:
   ```bash
   npm start
   ```

   The React app will be available at `http://localhost:3000`.

#### 4. **Setting Up MongoDB with Terraform on GCP**

1. **Navigate to the Terraform Directory**:

   ```bash
   cd terraform
   ```

2. **Initialize Terraform**:

   Run the following command to initialize Terraform and download necessary providers:

   ```bash
   terraform init
   ```

3. **Configure Your GCP Credentials**:

   Ensure that your GCP credentials are set up. You can follow the [GCP Authentication guide](https://cloud.google.com/docs/authentication/getting-started) to authenticate with GCP.

4. **Apply Terraform Configuration**:

   Apply the Terraform configuration to provision MongoDB on GCP:

   ```bash
   terraform apply
   ```

   Terraform will ask for confirmation. Type `yes` to continue. Once the infrastructure is created, you’ll get the **external IP** of the MongoDB instance.

5. **Update MongoDB Connection String**:

   In the backend code, update the MongoDB connection string to use the external IP of your MongoDB instance:

   ```go
   mongoURI := "mongodb://<username>:<password>@<GCP_MongoDB_IP>:27017"
   ```

---

### API Documentation

#### Endpoints:

1. **GET `/search?name=<character_name>`**
   - **Description**: Search for a character by name in the Rick and Morty API.
   - **Response**: Returns the character details (name, species, status, etc.) and episodes and saves it to database.

   **Example Request**:
   ```bash
   GET http://localhost:8080/search?name=Rick
   ```

2. **GET `/database/character?name=<character_name>`**
   - **Description**: Retrieve character data from the MongoDB database.
   - **Response**: Returns character details stored in the database (e.g., name, species, episodes).

   **Example Request**:
   ```bash
   GET http://localhost:8080/database/character?name=Rick Sanchez
   ```

---

### Architecture Overview

- **Frontend**: Built with React (TypeScript), responsible for displaying character data and interacting with the backend API for search and episode information.
- **Backend**: Built with Go, serving as a REST API that interacts with MongoDB and fetches data from the Rick and Morty API.
- **MongoDB**: Hosted on Google Cloud Platform (GCP), provisioned with Terraform, and stores character data.
- **Terraform**: Used for provisioning the MongoDB instance on GCP and configuring networking/firewall rules.

---



---

### License

MIT License

---

### Final Notes

- Make sure to **update environment variables** such as MongoDB connection details before running the backend.
- The frontend and backend should both be running locally for full functionality.
