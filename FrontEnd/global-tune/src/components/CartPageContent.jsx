"use client";
import React, { useState } from "react";
import Link from "next/link";

function CartTable({ products }) {
  const [focusedImg, setFocusedImg] = useState(-1);
  const [subtotal, setSubtotal] = useState(
    Math.ceil(products.reduce((acc, el) => acc + el.price, 0) * 100) / 100
  );
  const [shipping, setShipping] = useState(100);
  const [total, setTotal] = useState(subtotal + shipping);

  const [isTableModified, setIsTableModified] = useState(false);

  return (
    <div className="row">
      <section className="col-8">
        <table className="table">
          <thead className="table-dark">
            <tr>
              <th scope="col">Preview</th>
              <th scope="col">Product</th>
              <th scope="col">Unit Price</th>
              <th scope="col">Units</th>
              <th scope="col">Subtotal</th>
              <th scope="col"></th>
            </tr>
          </thead>
          <tbody>
            {products.map((product) => (
              <tr className="align-middle" key={product.id}>
                <th scope="row">
                  <Link href={`/product/${product.id}`}>
                    <img
                      src={product.thumbnail}
                      className={
                        focusedImg == product.id ? "border  border-warning" : ""
                      }
                      alt={`${product.title} preview`}
                      style={{ maxWidth: "100px" }}
                      onMouseEnter={() => setFocusedImg(product.id)}
                      onMouseLeave={() => setFocusedImg(-1)}
                    />
                  </Link>
                </th>
                <td>
                  <Link
                    href={`/product/${product.id}`}
                    className="link-offset-2 link-underline link-underline-opacity-0 link-opacity-25-hover"
                  >
                    {product.title}
                  </Link>
                </td>
                <td>${product.price}</td>
                <td>
                  <input
                    type="number"
                    className="form-control"
                    defaultValue={product.quantity}
                    min="1"
                    onChange={(e) => {
                      setIsTableModified(true);
                    }}
                  />
                </td>
                <td>
                  ${Math.ceil(product.quantity * product.price * 100) / 100}
                </td>
                <td>
                  <button className="btn btn-outline-danger">X</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>
      <section className="col-4 border p-3">
        <h5 className="text-center">CART TOTAL</h5>
        <p className="d-flex" style={{ justifyContent: "space-between" }}>
          <span style={{ fontWeight: "700" }}>Subtotal </span>${subtotal}
        </p>
        <hr />
        <p className="d-flex" style={{ justifyContent: "space-between" }}>
          <span style={{ fontWeight: "700" }}>Shipping details </span>$
          {shipping}
        </p>
        <p>Shipping to YOUR LOCATION</p>
        <p>Estimated delivery date: dd-mm-yyyy</p>
        <hr />
        <p className="d-flex" style={{ justifyContent: "space-between" }}>
          <span style={{ fontWeight: "700" }}>TOTAL </span>${total}
        </p>
        {/**
         // TODOs: 
         *  use next/Link on the revert changes button
         *  send a request to change values on cart on confirm changes Link
         *  Redirect to billing details on Complete Purchase button*/}
        {isTableModified ? (
          <>
            <Link
              href="/cart"
              className="btn btn-outline-warning w-100 mb-2"
              onClick={() => {
                setIsTableModified(false);
              }}
            >
              Confirm Changes
            </Link>
            <a href="/cart">
              <button className="btn btn-outline-danger w-100">
                Revert Changes
              </button>
            </a>
          </>
        ) : (
          <button className="btn btn-success w-100">Complete Purchase</button>
        )}
      </section>
    </div>
  );
}

export default CartTable;
