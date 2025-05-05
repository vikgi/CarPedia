import React from 'react';
import No_Image_Available from '../data/No_Image_Available.jpg';

const CarItem = ({ model, brand, year, color, price, description, imgurl }) => {
  return (
    <div className="card my-3">
      <img
        src={imgurl || No_Image_Available}
        className="card-img-top"
        alt="Car"
        style={{ height: '180px', objectFit: 'cover' }}
      />
      <div className="card-body">
        <h5 className="card-title">{brand} {model}</h5>
        <p className="card-text">
          <strong>Year:</strong> {year}<br />
          <strong>Color:</strong> {color}<br />
          <strong>Price:</strong> ${price}
        </p>
        <p className="card-text">
          <small className="text-muted">{description}</small>
        </p>
      </div>
    </div>
  );
};

export default CarItem;