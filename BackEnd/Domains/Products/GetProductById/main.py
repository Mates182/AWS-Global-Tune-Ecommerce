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
    def product_by_sku(self, sku: str) -> Product:
        # Query MongoDB for a product with a specific SKU
        product_data = products_collection.find_one({"sku": sku})  # Find the product by SKU
        if product_data:
            # Convert MongoDB document to Product object
            return Product(
                id=str(product_data["_id"]),
                title=product_data["title"],
                description=product_data["description"],
                category=product_data["category"],
                price=product_data["price"],
                stock=product_data["stock"],
                tags=product_data["tags"],
                brand=product_data["brand"],
                sku=product_data["sku"],
                weight=product_data["weight"],
                warranty=product_data["warranty"],
                thumbnail=product_data["thumbnail"],
                images=product_data["images"]
            )
        return None  # Return None if no product found

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
