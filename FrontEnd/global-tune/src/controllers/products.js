const PRODUCTS_CREATE_ENDPOINT = `http://98.81.173.125.30:80/create/`;
const PRODUCTS_LIST_ENDPOINT = process.env.PRODUCTS_LIST_ENDPOINT;
const PRODUCTS_GET_ENDPOINT = process.env.PRODUCTS_GET_ENDPOINT;
const PRODUCTS_UPLOAD_ENDPOINT = `http://54.159.208.108/upload/`;

export const loadProducts = async () => {
  const res = await fetch(PRODUCTS_LIST_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
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

  if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);

  const { data } = await res.json();
  return data;
};

export const loadCategories = async () => {
  const res = await fetch(PRODUCTS_LIST_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{ products { category } }`,
    }),
  });

  if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);

  const { data } = await res.json();
  const { products } = data;

  const categoryCount = products.reduce((acc, product) => {
    acc[product.category] = (acc[product.category] || 0) + 1;
    return acc;
  }, {});

  return Object.entries(categoryCount).map(([name, count]) => ({
    name,
    count,
  }));
};

export const loadProduct = async (productId) => {
  const res = await fetch(PRODUCTS_GET_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{ productBySku(sku: "${productId}") {
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
  });

  if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);

  const { data } = await res.json();
  return data.productBySku;
};

// Upload an image and return its URL
export const uploadImage = async (file) => {
  console.log("Upload Endpoint:xd", PRODUCTS_UPLOAD_ENDPOINT);
  console.log("Create Endpoint:xd", PRODUCTS_CREATE_ENDPOINT);
  console.log("list Endpoint:xd", PRODUCTS_LIST_ENDPOINT);

  const formData = new FormData();
  formData.append("image", file);
  formData.append("filename", file.name);

  const response = await fetch(PRODUCTS_UPLOAD_ENDPOINT, {
    method: "POST",
    body: formData,
    headers: { "Content-Type": "application/json" },
    redirect: "follow",
  });

  if (!response.ok) throw new Error("Image upload failed");

  const data = await response.json();
  return data.image_url;
};

// Create a new product
export const createProduct = async (productData) => {
  const formattedProduct = {
    Product: {
      Tags: productData.tags || [],
      Weight: parseFloat(productData.weight) || 0,
      Description: productData.description || "",
      Price: parseFloat(productData.price) || 0,
      Category: productData.category || "",
      Brand: productData.brand || "",
      Sku: productData.sku || "",
      Warranty: productData.warranty || "",
      _id: productData.id || "",
      Images: productData.images || [],
      Thumbnail: productData.thumbnail || "",
      Title: productData.title || "",
      Stock: parseInt(productData.stock) || 0,
      Active: productData.active ?? true,
    },
  };

  const response = await fetch(PRODUCTS_CREATE_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(formattedProduct),
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error(`Error on create: ${response.statusText}`);
  }

  return response.json();
};
