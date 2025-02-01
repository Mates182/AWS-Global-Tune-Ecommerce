"use client";
import { useRouter } from "next/navigation";
function AddToCartForm({ product, productId }) {
  const router = useRouter();
  const onSubmit = (e) => {
    e.preventDefault();
    const quantity = e.target.units.value;
    router.push(`/product/${productId}?quantity=${quantity}`);
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
      </form>
    </>
  );
}

export default AddToCartForm;
