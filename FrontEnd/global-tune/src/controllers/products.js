const PRODUCTS_CREATE_ENDPOINT = `http://172.31.23.6/create`;
const PRODUCTS_LIST_ENDPOINT = process.env.PRODUCTS_LIST_ENDPOINT;
const PRODUCTS_GET_ENDPOINT = process.env.PRODUCTS_GET_ENDPOINT;
const PRODUCTS_UPLOAD_ENDPOINT = `http://54.159.208.108/upload`;

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

export const uploadImage = async (file) => {
  const formData = new FormData();
  formData.append("image", file);  // File
  formData.append("filename", file.name);  // Custom filename
  formData.append("title", "Sample Image");  // Example Title
  formData.append("description", "An example image");  // Example Description

  const response = await fetch(PRODUCTS_UPLOAD_ENDPOINT, {
    method: "POST",
    body: formData,
    headers: {},
    redirect: "follow",
  });

  if (!response.ok) {
    const {data} = response
    console.log(data)
    throw new Error(`Image upload failed: ${response.status} ${response.statusText}` );
  }

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
