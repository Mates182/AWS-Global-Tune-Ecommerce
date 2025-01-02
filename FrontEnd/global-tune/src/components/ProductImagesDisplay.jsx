"use client";
import React, { useState } from "react";

function ProductImagesDisplay({ images, title}) {
  const [selectedImg, setSelectedImg] = useState(0);
  const [focusedImg, setFocusedImg] = useState(-1);
  return (
    <>
      <img
        src={images[selectedImg]}
        alt={`${title} image`}
        className="mw-100 border"
      />
      <hr />
      <div
        className="d-flex"
        style={{ overflowX: "auto", whiteSpace: "nowrap" }}
      >
        {images.map((image, i) => (
          <img
            key={i}
            src={images[i]}
            alt={`${title} image`}
            className={`border ${
              i == selectedImg
                ? "border-dark"
                : i == focusedImg
                ? "border-success"
                : ""
            } me-2`}
            style={{ height: "100px", width: "auto", cursor: "pointer" }}
            onClick={() => {
              setSelectedImg(i);
            }}
            onMouseEnter={()=>{
                setFocusedImg(i)
            }}
            onMouseLeave={()=>{
                setFocusedImg(-1)
            }}
            
          />
        ))}
      </div>
    </>
  );
}

export default ProductImagesDisplay;
