"use client";
import React, { useState } from "react";
import { uploadImage, createProduct } from "@/controllers/products";

const CreateProductForm = () => {
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    price: "",
    stock: "",
    category: "",
    brand: "",
    sku: "",
    warranty: "",
    weight: "",
    tags: "",
    thumbnail: null,
    images: [],
    active: true,
  });

  const [uploading, setUploading] = useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleFileChange = (e) => {
    const files = e.target.files;
    if (files.length > 0) {
      setFormData((prev) => ({
        ...prev,
        [e.target.name]: e.target.name === "thumbnail" ? files[0] : [...files],
      }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setUploading(true);
    
    try {
      let uploadedThumbnail = "";
      let uploadedImages = [];

      if (formData.thumbnail) {
        uploadedThumbnail = await uploadImage(formData.thumbnail);
      }

      if (formData.images.length > 0) {
        uploadedImages = await Promise.all(
          formData.images.map((file) => uploadImage(file))
        );
      }

      const productData = {
        title: formData.title,
        description: formData.description,
        price: parseFloat(formData.price),
        stock: parseInt(formData.stock),
        category: formData.category,
        brand: formData.brand,
        sku: formData.sku,
        warranty: formData.warranty,
        weight: parseFloat(formData.weight),
        tags: formData.tags.split(",").map((tag) => tag.trim()), // Convert CSV tags to an array
        thumbnail: uploadedThumbnail,
        images: uploadedImages,
        active: formData.active === "true",
      };

      await createProduct(productData);
      alert("Product created successfully!");
    } catch (error) {
      console.error(error);
      alert("Error creating product");
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="container">
      <h1>Create a New Product</h1>
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <div className="form-group">
          <label>Title</label>
          <input type="text" name="title" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Description</label>
          <textarea name="description" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Price</label>
          <input type="number" name="price" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Stock</label>
          <input type="number" name="stock" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Category</label>
          <input type="text" name="category" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Brand</label>
          <input type="text" name="brand" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>SKU</label>
          <input type="text" name="sku" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Warranty</label>
          <input type="text" name="warranty" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Weight (kg)</label>
          <input type="number" step="0.1" name="weight" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Tags (comma-separated)</label>
          <input type="text" name="tags" onChange={handleChange} required />
        </div>

        <div className="form-group">
          <label>Active</label>
          <select name="active" onChange={handleChange} required>
            <option value="true">Yes</option>
            <option value="false">No</option>
          </select>
        </div>

        <div className="form-group">
          <label>Thumbnail</label>
          <input type="file" name="thumbnail" accept="image/*" onChange={handleFileChange} required />
        </div>

        <div className="form-group">
          <label>Product Images</label>
          <input type="file" name="images" accept="image/*" multiple onChange={handleFileChange} />
        </div>

        <button type="submit" disabled={uploading}>
          {uploading ? "Uploading..." : "Create Product"}
        </button>
      </form>
    </div>
  );
};

export default CreateProductForm;
