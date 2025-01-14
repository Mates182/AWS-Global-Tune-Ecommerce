import { CartDTO } from "@/DTOs/CartDTO.js";

const CART_DOMAIN_ENDPOINT = process.env.CART_DOMAIN_ENDPOINT;

export class CartDAO {
  static async getCartById(id) {
    const response = await fetch(`${CART_DOMAIN_ENDPOINT}/get/${id}`);
    if (!response.ok) {
      throw new Error(`Error fetching cart: ${response.statusText}`);
    }
    const data = await response.json();
    console.log(JSON.stringify(data));
    return new CartDTO({
      id: id,
      items: data,
    });
  }

  static async createOrUpdateCart(cart) {
    if (cart instanceof CartDTO) {
      const response = await fetch(`${CART_DOMAIN_ENDPOINT}/set/`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(cart),
      });
      if (!response.ok) {
        throw new Error(`Error creating/updating cart: ${response.statusText}`);
      }
      const data = await response.json();
      return data;
    } else {
      throw new Error(`Error parsing cart`);
    }
  }

  static async deleteCartItem(id, productId) {
    if (!id || !productId) {
      throw new Error("Cart ID and at least one product ID are required.");
    }
    const response = await fetch(
      `${CART_DOMAIN_ENDPOINT}/delete/${id}?product=${productId}`,
      {
        method: "DELETE",
      }
    );
    if (!response.ok) {
      throw new Error(
        `Error deleting products from cart: ${response.statusText}`
      );
    }
    return response.json();
  }
}
