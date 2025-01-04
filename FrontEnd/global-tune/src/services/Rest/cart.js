const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export const getCartById = async (id) => {
  const response = await fetch(`${API_BASE_URL}/cart/${id}`);
  if (!response.ok) {
    throw new Error(`Error fetching cart: ${response.statusText}`);
  }
  return response.json();
};

export const createOrUpdateCart = async (cart) => {
  const response = await fetch(`${API_BASE_URL}/cart/`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(cart),
  });
  if (!response.ok) {
    throw new Error(`Error creating/updating cart: ${response.statusText}`);
  }
  return response.json();
};
