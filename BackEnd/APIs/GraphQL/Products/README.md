# Products GraphQL API - Python

A GraphQL API built with Python using FastAPI and Strawberry to fetch the products.  

## 🚀 Getting Started (To run locally)

### 🌟 Create a Virtual Environment  
- **For Windows**:  
  ```bash
  python -m venv tenv
  ```

- **For Mac/Linux**:  
  ```bash
  python3 -m venv tenv
  ```

### 🔑 Activate the Virtual Environment  

  ```bash
  source tenv/Scripts/activate
  ```

### 📦 Install Dependencies  
```bash
pip install -r requirements.txt
```

### ▶️ Run the Application  
```bash
python main.py
```

The API will be available at:  
- **Homepage**: [http://localhost:27017](http://localhost:27017)  
- **GraphQL Endpoint**: [http://localhost:27017/graphql](http://localhost:27017/graphql)  

## 🔍 Example Usage  

### 📝 Example Query:

To fetch the list of products from the API, you can use a query like this in a GraphQL client or a tool like Postman or Insomnia:

```graphql
query {
  products {
    name
    price
    description
  }
}
```

This query will retrieve the `name`, `price`, and `description` of all available products.

### ⛔ Deactivate the Virtual Environment  
When done, you can deactivate the virtual environment:  
```bash
deactivate
```

---

# TODOs
- Implement other necessary queries when needed