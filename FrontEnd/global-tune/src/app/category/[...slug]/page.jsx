import React from "react";

function CategoryPage({ params }) {
  const { slug } = params; // url segments
  return (
    <div>
      <h1>{decodeURIComponent(slug[slug.length - 1].split("-")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(" "))}</h1>
    </div>
  );
}

export default CategoryPage;
