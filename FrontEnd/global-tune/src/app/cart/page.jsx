import React from "react";
import Content from "@/components/CartPageContent";
import { getCartById, createOrUpdateCart } from "@/services/Rest/cart";
import { loadProduct } from "@/services/GraphQL/products";

async function CartPage() {
  // TODO: validate if user is logged in, use discounts and shippings, here we're using a test
  let products;
  let cart;
  try {
    cart = await getCartById(2);

    if (cart.message === "Cart is empty") {
      products = [];
    } else {
      products = await Promise.all(
        Object.entries(cart).map(async ([id, quantity]) => {
          console.log(id)
          const productBySku = await loadProduct(id);
          console.log(productBySku)
          return { ...productBySku, quantity, id };
        }))
    }
  } catch (error) {
    console.error("Error fetching cart:", error.message);
  }

  return <Content products={products}></Content>;
}

export default CartPage;
