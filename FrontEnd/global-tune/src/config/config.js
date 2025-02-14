const config = {
    PRODUCTS_UPLOAD_ENDPOINT: process.env.NEXT_PUBLIC_PRODUCTS_UPLOAD_ENDPOINT,
    PRODUCTS_CREATE_ENDPOINT: process.env.NEXT_PUBLIC_PRODUCTS_CREATE_ENDPOINT,
    PRODUCTS_LIST_ENDPOINT: process.env.NEXT_PUBLIC_PRODUCTS_LIST_ENDPOINT,
    PRODUCTS_GET_ENDPOINT: process.env.NEXT_PUBLIC_PRODUCTS_GET_ENDPOINT,
  };
  
  // Debug: Check if env variables are loaded correctly
  console.log("Config Loaded:", config);
  
  export default config;
  