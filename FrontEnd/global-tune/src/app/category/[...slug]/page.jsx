import React from "react";
import ProductCard from "@/components/ProductCard";
import Link from 'next/link'

// TODOs: Consume the api for our products
// Display only the products with the category
const loadProducts = async () => {
  const res = await fetch("https://dummyjson.com/products"); //using dummyjson products api for placeholders
  const data = await res.json();
  return data;
};
const loadCategories = async () => {
  const res = await fetch("https://dummyjson.com/products/categories"); //using dummyjson products api for placeholders
  const data = await res.json();
  console.log(data)
  return data;
};
// TODO: Add discounts and brands, display only x number of products
async function CategoryPage({ params }) {
  const { slug } = params; // url segments
  const { products } = await loadProducts();
  const categories = await loadCategories()
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
        {categories.map((category, i)=>(
          <Link href={`/category/musical-instruments/${category.name.toLowerCase().replace(/\s/g, '-')}`} key={i} className="list-group-item">{category.name}</Link>
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
