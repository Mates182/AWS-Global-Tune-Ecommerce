import React from "react";
import AddToCartForm from "@/components/AddToCartForm";
import ProductImagesDisplay from "@/components/ProductImagesDisplay";

// TODOs: Consume the api for our products
const loadProducts = async (id) => {
  const res = await fetch(`https://dummyjson.com/products/${id}`); //using dummyjson products api for placeholders
  const data = await res.json();
  console.log(data);
  return data;
};

async function ProductPage({ params }) {
  const { productId } = await params;
  const product = await loadProducts(productId);
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
