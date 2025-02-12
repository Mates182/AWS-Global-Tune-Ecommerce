import os
from fastapi import FastAPI
import strawberry
from strawberry.asgi import GraphQL
from typing import List
from fastapi.middleware.cors import CORSMiddleware
from pymongo import MongoClient
from dotenv import load_dotenv, dotenv_values

load_dotenv()

client = MongoClient(os.getenv("products_URI"))
db = client["products"]
products_collection = db["products"]

@strawberry.type
class Product:
    id: str
    title: str
    description: str
    category: str
    price: float
    stock: int
    tags: List[str]
    brand: str
    sku: str
    weight: float
    warranty: str
    thumbnail: str
    images: List[str]


@strawberry.type
class Query:
    @strawberry.field
    def products(self) -> List[Product]:
        # Fetch products from MongoDB
        products_data = list(products_collection.find())  # Retrieve all products from the collection
        
        # Convert MongoDB documents to Product objects
        products = [
            Product(
                id=str(product["_id"]),
                title=product["title"],
                description=product["description"],
                category=product["category"],
                price=product["price"],
                stock=product["stock"],
                tags=product["tags"],
                brand=product["brand"],
                sku=product["sku"],
                weight=product["weight"],
                warranty=product["warranty"],
                thumbnail=product["thumbnail"],
                images=product["images"]
            )
            for product in products_data
        ]
        return products

schema = strawberry.Schema(query=Query)
app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"], 
    allow_credentials=True,
    allow_methods=["*"],  
    allow_headers=["*"], 
)

@app.get("/")
def index():
    return {"message": "Welcome to the GraphQL API"}

app.add_route("/graphql", GraphQL(schema, debug=True))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=80)
