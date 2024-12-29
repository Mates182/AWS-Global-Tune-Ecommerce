import React from "react";

function CartPage() {
  // TODO: implement cart logic, validate if user is logged in
  const cartItems = 0;
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
            <tr className="align-middle">
              <th scope="row">img</th>
              <td>Product Name</td>
              <td>$50</td>
              <td>
                <input
                  type="number"
                  className="form-control"
                  defaultValue="1"
                  min="1"
                />
              </td>
              <td>$500</td>
              <td>
                <button className="btn btn-outline-danger">X</button>
              </td>
            </tr>
            <tr className="align-middle">
              <th scope="row">img</th>
              <td>Product Name</td>
              <td>$50</td>
              <td>
                <input
                  type="number"
                  className="form-control"
                  defaultValue="1"
                  min="1"
                />
              </td>
              <td>$500</td>
              <td>
                <button className="btn btn-outline-danger">X</button>
              </td>
            </tr>
            <tr className="align-middle">
              <th scope="row">img</th>
              <td>Product Name</td>
              <td>$50</td>
              <td>
                <input
                  type="number"
                  className="form-control"
                  defaultValue="1"
                  min="1"
                />
              </td>
              <td>$500</td>
              <td>
                <button className="btn btn-outline-danger">X</button>
              </td>
            </tr>
            <tr className="align-middle">
              <th scope="row">img</th>
              <td>Product Name</td>
              <td>$50</td>
              <td>
                <input
                  type="number"
                  className="form-control"
                  defaultValue="1"
                  min="1"
                />
              </td>
              <td>$500</td>
              <td>
                <button className="btn btn-outline-danger">X</button>
              </td>
            </tr>
          </tbody>
        </table>
      </section>
      <section className="col-4 border p-3">
        <h5 className="text-center">CART TOTAL</h5>
        <p className="d-flex" style={{justifyContent:'space-between'}}>
          <span style={{ fontWeight: "700" }}>Subtotal </span>
          $00000
        </p>
        <hr />
        <p className="d-flex" style={{justifyContent:'space-between'}}>
          <span style={{ fontWeight: "700" }}>Shipping details </span>
          $00000
        </p>
        <p>Shipping to YOUR LOCATION</p>
        <p>Estimated delivery date: </p>
        <hr />
        <p className="d-flex" style={{justifyContent:'space-between'}}>
          <span style={{ fontWeight: "700" }}>TOTAL </span>
          $00000
        </p>
        {/* Redirect to billing details */}
        <button className="btn btn-warning w-100">Complete Purchase</button>
      </section>
    </div>
  );
}

export default CartPage;
