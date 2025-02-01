const PRODUCTS_ENDPOINT = process.env.PRODUCTS_ENDPOINT;

export const loadProducts = async () => {
  const res = await fetch(`${PRODUCTS_ENDPOINT}/all/`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      query: `
          {
            products {
              brand
              category
              id
              thumbnail
              title
              price
              stock
              sku
            }
          }
        `,
    }),
  });

  if (!res.ok) {
    throw new Error(`HTTP error! Status: ${res.status}`);
  }

  const { data } = await res.json();
  return data;
};

export const loadCategories = async () => {
  const res = await fetch(`${PRODUCTS_ENDPOINT}/all/`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      query: `
          {
            products {
              category
            }
          }
        `,
    }),
  });

  if (!res.ok) {
    throw new Error(`HTTP error! Status: ${res.status}`);
  }
  const { data } = await res.json();
  const { products } = data;
  const categoryCount = products.reduce((acc, product) => {
    const category = product.category;
    acc[category] = (acc[category] || 0) + 1;
    return acc;
  }, {});

  const categories = Object.entries(categoryCount).map(([name, count]) => ({
    name,
    count,
  }));
  return categories;
};

export const loadProduct = async (productId) => {
  const res = await fetch(`${PRODUCTS_ENDPOINT}/id/`, {
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
      console.log("Error: " + error.message);
    });
  const { data } = res;
  const { productBySku } = data;
  return productBySku;
};

export const createProduct = async (body) => {
  const response = await fetch(`${PRODUCTS_ENDPOINT}/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
    credentials: "include",
  });
  if (!response.ok) {
    throw new Error(`Error on create: ${response.statusText}`);
  }
  return response.json();
};
