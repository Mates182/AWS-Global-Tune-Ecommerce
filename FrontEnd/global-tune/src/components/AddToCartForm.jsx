"use client";
import React, { useState } from "react";
import Link from "next/link";

function AddToCartForm({ product }) {
  const [hidden, setHidden] = useState(true);
  const onSubmit = (e) => {
    e.preventDefault();
    setHidden(false);
    // TODO: Add the product units to the cart => e.target.units.value
  };

  const { stock } = product;
  return (
    <>
      <form onSubmit={onSubmit} style={{ maxWidth: "400px" }}>
        <div className="input-group mb-3">
          <span className="input-group-text" id="basic-addon1">
            Units:{" "}
          </span>
          <input
            type="number"
            defaultValue="1"
            className="form-control"
            max={stock}
            min="1"
            id="units"
          />

          <button className="btn btn-outline-info">ADD TO CART!</button>
        </div>
        {!hidden && (
          <div className="modal-overlay" onClick={() => setHidden(true)}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
              <h2>Added to your Cart!</h2>
              <p>
                You can continue shopping or go to your cart and continue with
                the payment
              </p>
              <button
                onClick={(e) => {
                  e.preventDefault();
                  setHidden(true);
                }}
                className="btn btn-outline-primary mb-1"
              >
                Continue Shopping
              </button>
              {/* TODO: Validate if the user is logged in */}
              <Link
                href="/cart"
                onClick={() => {
                  setHidden(true);
                }}
                className="btn btn-outline-success"
              >
                View Cart
              </Link>
            </div>
          </div>
        )}
        {/* This style works here, on global.css dont show the background color */}
        <style jsx>{`
          .modal-overlay {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 0, 0, 0.5);
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 1000;
          }

          .modal-content {
            background-color: #fff;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            max-width: 500px;
            width: 100%;
            text-align: center;
          }
        `}</style>
      </form>
    </>
  );
}

export default AddToCartForm;
