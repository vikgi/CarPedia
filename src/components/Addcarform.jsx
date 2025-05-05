import React, { useState } from 'react';
import axios from "axios";
import './Addcarform.css'
const Addcarform = () => {
  const [car, setCar] = useState({
    model: '',
    brand: '',
    year: '',
    color: '',
    price: '',
    description: '',
    imgurl: '',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCar({ ...car, [name]: name === "price" ? Number(value) : value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const userId = localStorage.getItem("userId");
    const token = localStorage.getItem("token");

    const carData = { ...car };

    try {
      const res = await axios.post(
        `http://localhost:5000/api/cars/user?userId=${userId}`,
        carData,
        { headers: { "Authorization": `Bearer ${token}` } }
      );
    } catch (err) {
      console.error("Error adding car", err);
    }

    setCar({
      model: '',
      brand: '',
      year: '',
      color: '',
      price: '',
      description: '',
      imgurl: '',
    });
  };

  return (
    <form
      className="p-3 rounded-4 bg-light-grey custom-shadow"
      onSubmit={handleSubmit}
      noValidate
    >
      <h4 className="mb-4 text-danger text-black fw-semibold text-center">Add a New Car</h4>

      <div className="row g-2 mb-2">
        <div className="col-md-6">
          <label className="form-label">Model</label>
          <input
            type="text"
            name="model"
            value={car.model}
            onChange={handleInputChange}
            className="form-control rounded- red-focus"
            placeholder="Enter car model" 
            required
          />
        </div>
        <div className="col-md-6">
          <label className="form-label">Brand</label>
          <input
            type="text"
            name="brand"
            value={car.brand}
            onChange={handleInputChange}
            className="form-control rounded-3 red-focus"
            placeholder="Enter brand"
            required
          />
        </div>
      </div>

      <div className="row g-3 mb-3">
        <div className="col-md-4">
          <label className="form-label">Year</label>
          <input
            type="text"
            name="year"
            value={car.year}
            onChange={handleInputChange}
            className="form-control rounded-3 red-focus"
            placeholder="e.g. 2023"
            required
          />
        </div>
        <div className="col-md-4">
          <label className="form-label">Color</label>
          <input
            type="text"
            name="color"
            value={car.color}
            onChange={handleInputChange}
            className="form-control rounded-3 red-focus"
            placeholder="e.g. Grey"
            required
          />
        </div>
        <div className="col-md-4">
          <label className="form-label">Price ($)</label>
          <input
            type="number"
            name="price"
            value={car.price}
            onChange={handleInputChange}
            className="form-control rounded-3 red-focus"
            placeholder="e.g. 25000"
            required
          />
        </div>
      </div>

      <div className="mb-3">
        <label className="form-label">Image URL</label>
        <input
          type="text"
          name="imgurl"
          value={car.imgurl}
          onChange={handleInputChange}
          className="form-control rounded-3 red-focus"
          placeholder="Paste image link"
          required
        />
      </div>

      <div className="mb-4">
        <label className="form-label">Description</label>
        <textarea
          name="description"
          value={car.description}
          onChange={handleInputChange}
          className="form-control rounded-3 red-focus"
          placeholder="Brief description..."
          rows="3"
        />
      </div>

      <div className="d-grid">
        <button type="submit" className="btn btn-danger rounded-3">
          Add Car
        </button>
      </div>
    </form>
  );
};

export default Addcarform;
