import React from "react";
import ProductCard from "@/components/ProductCard";

// TODOs: Consume the api for our products
// Display only the products with the category
const loadProducts = async () => {
  const res = await fetch("https://dummyjson.com/products"); //using dummyjson products api for placeholders
  const data = await res.json();
  return data;
};
// TODO: Add categories, display only x number of products
async function CategoryPage({ params }) {
  const { slug } = params; // url segments
  const { products } = await loadProducts();
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
        <div>
          <h3>Categories</h3>
        </div>
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
