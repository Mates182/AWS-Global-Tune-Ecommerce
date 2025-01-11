const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export const getCartById = async (id) => {
  const response = await fetch(`${API_BASE_URL}/get/${id}`);
  if (!response.ok) {
    throw new Error(`Error fetching cart: ${response.statusText}`);
  }
  return response.json();
};

export const createOrUpdateCart = async (cart) => {
  const response = await fetch(`${API_BASE_URL}/set/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(cart),
  });
  if (!response.ok) {
    throw new Error(`Error creating/updating cart: ${response.statusText}`);
  }
  return response.json();
};

export const deleteCartProducts = async (id, productId) => {
  if (!id || !productId ) {
    throw new Error("Cart ID and at least one product ID are required.");
  }
  const response = await fetch(`${API_BASE_URL}/delete/${id}?product=${productId}`, {
    method: "DELETE",
  });

  if (!response.ok) {
    throw new Error(
      `Error deleting products from cart: ${response.statusText}`
    );
  }

  return response.json();
};
