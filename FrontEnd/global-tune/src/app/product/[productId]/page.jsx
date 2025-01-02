import React from "react";
import AddToCartForm from "@/components/AddToCartForm";
import ProductImagesDisplay from "@/components/ProductImagesDisplay";

// TODOs: Consume the api for our products, here we're using a local graphql api
// Display only the products with the category
const loadProduct = async (productId) => {
  const res = await fetch("http://localhost:27017/graphql/", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      query: `{productBySku(sku: "${productId}") {
    title
    images
    stock
    price
    description
    category
    brand
    thumbnail
  }}`,
    }),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json();
    })
    .catch((error) => {
      alert("Error: " + error.message);
    });
  const { data } = res;
  const {productBySku} = data
  return productBySku;
};

async function ProductPage({ params }) {
  const { productId } = await params;
  const product = await loadProduct(productId);
  return (
    <div className="row">
      <section className="col-4">
        <ProductImagesDisplay
          images={product.images}
          title={product.title}
        ></ProductImagesDisplay>
      </section>
      <section className="col-8">
        <h1>{product.title}</h1>
        <h3>${product.price}</h3>

        <p>{product.description}</p>
        <p>
          <span style={{ fontWeight: "700" }}>Brand: </span>
          {product.brand}
        </p>
        <p>
          <span style={{ fontWeight: "700" }}>Category: </span>
          {product.category}
        </p>
        {/* TODO: and other info and button functionality */}
        {product.stock > 0 ? (
          <>
            <h3 className="text-success">{product.stock} IN STOCK!</h3>
            <AddToCartForm product={product}></AddToCartForm>
          </>
        ) : (
          <h3 className="text-danger">OUT OF STOCK</h3>
        )}
      </section>
    </div>
  );
}

export default ProductPage;
