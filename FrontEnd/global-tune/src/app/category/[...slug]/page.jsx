import React from "react";
import ProductCard from "@/components/ProductCard";
import Link from "next/link";

// TODOs: Consume the api for our products, here we're using a local graphql api
// Display only the products with the category
const loadProducts = async () => {
  const res = await fetch("http://localhost:27017/graphql/", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      query: `{products {
    brand
    category
    id
    thumbnail
    title
    price
    stock
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
  const {data} = res
  return data;
};
const loadCategories = async () => {
  const res = await fetch("https://dummyjson.com/products/categories"); //using dummyjson products api for placeholders
  const data = await res.json();
  return data;
};
// TODO: Add discounts and brands, display only x number of products
async function CategoryPage({ params }) {
  const { slug } = await params; // url segments
  const { products } = await loadProducts();
  const categories = await loadCategories();
  return (
    <div className="container row">
      <h1>
        {decodeURIComponent(
          slug[slug.length - 1]
            .split("-")
            .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
            .join(" ")
        )}
      </h1>
      <aside className="col">
        <ul className="list-group list-group-flush">
          <h3>Categories</h3>
          {categories.map((category, i) => (
            <Link
              href={`/category/musical-instruments/${category.name
                .toLowerCase()
                .replace(/\s/g, "-")}`}
              key={i}
              className="list-group-item"
            >
              {category.name}
            </Link>
          ))}
        </ul>
      </aside>
      <section className="col-8">
        <div className="row row-cols-4">
          {products.map((product) => (
            <ProductCard product={product} />
          ))}
        </div>
      </section>
    </div>
  );
}

export default CategoryPage;
