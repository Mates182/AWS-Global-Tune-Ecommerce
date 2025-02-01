"use client";
import React, { useState, Suspense } from "react";
import Link from "next/link";
import BillingDetails from "@/components/BillingDetails";
import { BsFillTrashFill } from "react-icons/bs";
import { useRouter } from "next/navigation";

function CartTable({ products, cart, redirect }) {
  const router = useRouter();
  if (redirect) {
    router.push("/cart");
  }
  const isHidden = false;
  const [productsTemp, setProductsTemp] = useState(
    products.map((product) => {
      return { ...product, isHidden };
    })
  );

  if (productsTemp.length == 0) {
    return (
      <div className="mt-3 d-flex flex-column justify-content-center">
        <h2 className="text-center">Your cart seems to be Empty</h2>
        <div className="m-auto">
          <img
            src="https://img.freepik.com/premium-photo/guitar-shopping-cart-white-background_256339-2827.jpg?w=1060"
            alt="Empty Cart"
            style={{ maxWidth: "444px" }}
          />
        </div>

        <p className="mb-4 text-center">
          But the music never stops. Explore our collection of instruments and
          find the one that will make you feel the vibe.
        </p>
        <Link href="/category/products" className="btn btn-outline-success">
          Explore instruments
        </Link>
      </div>
    );
  }
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
    let delList = [];
    let postList = [];
    setIsTableModified(false);
    setProductsTemp(
      productsTemp.map((product, i) => {
        let productTemp = product;
        productTemp.quantity = e.target[i * 2].value;
        if (e.target[i * 2].min == 0) {
          delList.push(product.id);
        } else {
          postList.push({
            product_id: product.id,
            quantity: parseInt(product.quantity),
          });
        }
        return productTemp;
      })
    );
    setSubtotal(
      Math.ceil(
        productsTemp.reduce((acc, el) => acc + el.price * el.quantity, 0) * 100
      ) / 100
    );
    setTotal(subtotal + shipping);
    router.push(
      `/cart/?${
        delList.length > 0
          ? postList.length > 0
            ? `del=${JSON.stringify(delList)}&post=${JSON.stringify(postList)}`
            : `del=${JSON.stringify(delList)}`
          : postList.length > 0
          ? `post=${JSON.stringify(postList)}`
          : ""
      }`
    );
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
              {productsTemp.map((product, i) => (
                <tr
                  className={`align-middle ${product.isHidden ? "d-none" : ""}`}
                  key={i}
                >
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
                      className={`form-control`}
                      defaultValue={product.quantity}
                      min={product.isHidden ? "0" : "1"}
                      onChange={(e) => {
                        setIsTableModified(true);
                      }}
                    />
                  </td>
                  <td>
                    ${Math.ceil(product.quantity * product.price * 100) / 100}
                  </td>
                  <td>
                    <button
                      className="btn btn-outline-danger"
                      onClick={(e) => {
                        e.preventDefault();
                        setIsTableModified(true);
                        setProductsTemp(
                          productsTemp.map((productTemp, j) => {
                            if (j == i) {
                              productTemp.isHidden = true;
                            }
                            return productTemp;
                          })
                        );
                      }}
                    >
                      <BsFillTrashFill />
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
          <BillingDetails id="btd"></BillingDetails>
        </>
      )}
    </>
  );
}

export default CartTable;
