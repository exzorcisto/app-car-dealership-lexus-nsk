import React, { useState, useEffect } from 'react'
import './AvailableCars.css'
import NavBar from '../../components/UI/NavBar/NavBar'
import axios from 'axios';

function AvailableCars() {
    const [cars, setCars] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8000/cars')
            .then(response => {
                console.log("Data received:", response.data); // Add this!
                setCars(response.data);
            })
            .catch(error => {
                console.error('Error fetching cars:', error);
            });
    }, []);

    return (
        <div>
            <NavBar>Контакты</NavBar>
            <h2>Car List</h2>
            <ul>
                {cars.map(car => (
                    <li key={car.car_id}> {/* Use car.car_id as the key */}
                        {car.year} {car.trim_level} {car.color} - ${car.price} {/* Access correct properties */}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default AvailableCars;

