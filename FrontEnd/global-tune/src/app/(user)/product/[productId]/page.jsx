import React from "react";
import AddToCartForm from "@/components/AddToCartForm";
import ProductImagesDisplay from "@/components/ProductImagesDisplay";
import { loadProduct } from "@/controllers/products";

import { CartDAO } from "@/DAOs/CartDAO.js";
import { CartDTO } from "@/DTOs/CartDTO.js";

import Link from "next/link";

async function ProductPage({ params, searchParams }) {
  const { productId } = await params;
  const { quantity } = await searchParams;
  let modal = quantity !== undefined;
  if (modal) {
    try {
      const cart = new CartDTO({
        id: "1",
        items: [{ product_id: productId, quantity: parseInt(quantity) }],
      });
      const response = await CartDAO.createOrUpdateCart(cart);
    } catch (error) {
      console.log("Error adding to cart", error);
      modal = false;
    }
  }
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
        {/* TODO: and other info*/}
        {product.stock > 0 ? (
          <>
            <h3 className="text-success">{product.stock} IN STOCK!</h3>
            <AddToCartForm
              product={product}
              productId={productId}
            ></AddToCartForm>
          </>
        ) : (
          <h3 className="text-danger">OUT OF STOCK</h3>
        )}

        {modal && (
          <div
            className="modal-overlay"
            style={{
              position: "fixed",
              top: 0,
              left: 0,
              width: "100%",
              height: "100%",
              backgroundColor: "rgba(0, 0, 0, 0.5)",
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              zIndex: 1000,
            }}
          >
            <div
              className="modal-content"
              style={{
                backgroundColor: "#fff",
                padding: "20px",
                borderRadius: "10px",
                boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
                maxWidth: "500px",
                width: "100%",
                textAlign: "center",
              }}
            >
              <h2>Added to your Cart!</h2>
              <p>
                You can continue shopping or go to your cart and continue with
                the payment
              </p>
              <Link
                href={`/product/${productId}`}
                className="btn btn-outline-primary mb-1"
              >
                Continue Shopping
              </Link>
              {/* TODO: Validate if the user is logged in */}
              <Link href="/cart" className="btn btn-outline-success">
                View Cart
              </Link>
            </div>
          </div>
        )}
      </section>
    </div>
  );
}

export default ProductPage;
