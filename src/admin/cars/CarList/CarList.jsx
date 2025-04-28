import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './CarList.css';

function CarList() {
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCars = async () => {
      try {
        const response = await axios.get('http://localhost:8000/cars', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('adminToken')}`
          }
        });
        setCars(response.data);
      } catch (err) {
        console.error('Error fetching cars:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchCars();
  }, []);

  if (loading) return <div>Loading...</div>;

  return (
    <div className="car-list">
      <div className="header">
        <h2>Cars Management</h2>
        <Link to="/admin/cars/add" className="add-btn">Add New Car</Link>
      </div>
      
      <table>
        <thead>
          <tr>
            <th>Model</th>
            <th>Trim</th>
            <th>Year</th>
            <th>Price</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {cars.map(car => (
            <tr key={car.carid}>
              <td>{car.model_name}</td>
              <td>{car.trimlevel}</td>
              <td>{car.year}</td>
              <td>{car.price} руб.</td>
              <td>
                <Link to={`/admin/cars/edit/${car.carid}`} className="edit-btn">Edit</Link>
                <button className="delete-btn">Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default CarList;