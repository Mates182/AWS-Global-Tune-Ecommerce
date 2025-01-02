export const loadProducts = async () => {
    const res = await fetch(process.env.GRAPHQL_PRODUCTS_ENDPOINT, {
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
    const res = await fetch(process.env.GRAPHQL_PRODUCTS_ENDPOINT, {
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
    const {data} = await res.json();
    const {products} = data
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
    const res = await fetch(process.env.GRAPHQL_PRODUCTS_ENDPOINT, {
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
        alert("Error: " + error.message);
      });
    const { data } = res;
    const {productBySku} = data
    return productBySku;
  };