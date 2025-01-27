import React from "react";
import ProductCard from "@/components/ProductCard";
import Link from "next/link";
import { loadProducts, loadCategories } from "@/controllers/products";

// TODO: Display only the products with the category
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
              {`${category.name} (${category.count})`}
            </Link>
          ))}
        </ul>
      </aside>
      <section className="col-8">
        <div className="row row-cols-4">
          {products.map((product, i) => (
            <ProductCard product={product} key={i} />
          ))}
        </div>
      </section>
    </div>
  );
}

export default CategoryPage;
