"use client";
import React, { useState, Suspense } from "react";
import Link from "next/link";
import BillingDetails from "@/components/BillingDetails";

function CartTable({ products }) {
  const [productsTemp, setProductsTemp] = useState([...products]);
  const [focusedImg, setFocusedImg] = useState(-1);
  const [subtotal, setSubtotal] = useState(
    Math.ceil(
      productsTemp.reduce((acc, el) => acc + el.price * el.quantity, 0) * 100
    ) / 100
  );
  const [shipping, setShipping] = useState(100);
  const [total, setTotal] = useState(subtotal + shipping);

  const [isTableModified, setIsTableModified] = useState(false);
  const [showBillingDetails, setShowBillingDetails] = useState(false);

  const onSubmit = (e) => {
    e.preventDefault();
    setIsTableModified(false);
    setProductsTemp(
      productsTemp.map((product, i) => {
        let productTemp = product;
        productTemp.quantity = e.target[i * 2].value;
        return productTemp;
      })
    );
    setSubtotal(
      Math.ceil(
        productsTemp.reduce((acc, el) => acc + el.price * el.quantity, 0) * 100
      ) / 100
    );
    setTotal(subtotal + shipping);
    // TODO: send a request to change values on cart on confirm changes Link
  };

  return (
    <>
      <form onSubmit={onSubmit} className="row">
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
              {productsTemp.map((product) => (
                <tr className="align-middle" key={product.id}>
                  <th scope="row">
                    <Link href={`/product/${product.id}`}>
                      <img
                        src={product.thumbnail}
                        className={
                          focusedImg == product.id
                            ? "border  border-warning"
                            : ""
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
                      id={`product${product.id}`}
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
                    {/**
                   // TODO: implement delete function
                   */}
                    <button
                      className="btn btn-outline-danger"
                      onClick={(e) => {
                        e.preventDefault();
                        alert("not implemented yet");
                      }}
                    >
                      X
                    </button>
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
         *  use next/Link on the revert changes <a>*/}
          {isTableModified ? (
            <>
              <button className="btn btn-outline-warning w-100 mb-2">
                Confirm Changes
              </button>
              <a href="/cart" className="btn btn-outline-danger w-100">
                Revert Changes
              </a>
            </>
          ) : (
            <button
              className={`btn btn-success w-100 ${
                showBillingDetails ? "d-none" : ""
              }`}
              onClick={(e) => {
                e.preventDefault();
                setShowBillingDetails(true);
              }}
            >
              Complete Purchase
            </button>
          )}
        </section>
      </form>
      {showBillingDetails && (
        <>
          <hr />
          <BillingDetails id='btd'></BillingDetails>
      
        </>
      )}
    </>
  );
}

export default CartTable;
