import React from "react";
import Link from "next/link";

// TODOs: switch to 2 different tipes of cards, to toggle the display of the product info
// Redirect to the product page on the Link
function ProductCard({ product }) {
  return (
    <div className="card" key={product.id}>
      <img
        className="card-img-top"
        src={product.images[0]}
        alt={`${product.title} image`}
        style={{ maxHeight: "300px" }}
      />
      <div className="card-body">
        <h5 className="card-title">{product.title}</h5>
        <p className="card-text">${product.price}</p>
        <Link href="#" className="btn btn-primary">
          VIEW PRODUCT
        </Link>
      </div>
    </div>
  );
}

export default ProductCard;
