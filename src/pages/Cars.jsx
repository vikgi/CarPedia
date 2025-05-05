import React, { useEffect, useState } from 'react';
import CarItem from '../components/CarItem';
import axios from 'axios';
import Navbar from '../components/Navbar';

const CarListPage = () => {
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchCars = async () => {
    try {
      const userId = localStorage.getItem("userId");
      const res = await axios.get(`http://localhost:5000/api/cars/user?userId=${userId}`);
      if(res.data == null){
        setCars([]);
        return;
      }
      setCars(res.data);
    } catch (err) {
      console.error("Error fetching cars:", err.response?.data?.error || err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCars();
  }, []);

  return (
    <>
    <Navbar />
      <div className="container mt-5">
        {loading ? (
          <p>Loading...</p>
        ) : cars.length === 0 ? (
          <p>No cars added yet.</p>
        ) : (
          <div className="row">
            {cars.map((car) => (
              <div className="col-md-4" key={car.id}>
                <CarItem
                  model={car.model}
                  brand={car.brand}
                  year={car.year}
                  color={car.color}
                  price={car.price}
                  description={car.description}
                  imgurl={car.imgurl}
                />
              </div>
            ))}
          </div>
        )}
      </div>
    </>
  );
};

export default CarListPage;