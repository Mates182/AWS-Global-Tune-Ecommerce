import React from "react";
import Content from '@/components/CartPageContent'


// TODOs: Consume the api for the cart
const loadCart = async (id) => {
  const res = await fetch(`https://dummyjson.com/carts/${id}`); //using dummyjson cart api for placeholders
  const data = await res.json();
  return data;
};

async function CartPage() {
  // TODO: validate if user is logged in, use discounts and shippings
  const { products } = await loadCart(2);
  
  return (
    <Content products={products}></Content>
  );
}

export default CartPage;
