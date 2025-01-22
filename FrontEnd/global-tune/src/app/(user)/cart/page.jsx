import React from "react";
import Content from "@/components/CartPageContent";
import { loadProduct } from "@/controllers/products";

import { CartDAO } from "@/DAOs/CartDAO.js";
import { CartDTO } from "@/DTOs/CartDTO.js";

async function CartPage({ searchParams }) {
  const { del, post } = await searchParams;
  // TODO: implement the redirect
  //const redirect = (del || post) ? true : false
  const redirect = false;
  if (post) {
    // TODO: implement the cart id, the id:"1" is for testing
    const cart = new CartDTO({
      id: "1",
      items: JSON.parse(post),
    });
    const res = await CartDAO.createOrUpdateCart(cart);
    console.log(res);
  }
  const delList = del ? JSON.parse(del) : [];
  await Promise.all(
    delList.map(async (element) => {
      await CartDAO.deleteCartItem(1, element);
    })
  );
  // TODO: validate if user is logged in, use discounts and shippings, here we're using a test cart
  let products;
  let cart;
  try {
    cart = await CartDAO.getCartById(1);
    if (cart.message === "Cart is empty") {
      products = [];
    } else {
      products = await Promise.all(
        Object.entries(cart.items).map(async ([id, quantity]) => {
          const productBySku = await loadProduct(id);
          return { ...productBySku, quantity, id };
        })
      );
    }
  } catch (error) {
    console.error("Error fetching cart:", error.message);
  }

  return <Content products={products} redirect={redirect}></Content>;
}

export default CartPage;
