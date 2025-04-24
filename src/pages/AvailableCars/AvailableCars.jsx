import React, { useState, useEffect } from 'react';
import './AvailableCars.css';
import NavBar from '../../components/UI/NavBar/NavBar';
import axios from 'axios';

function AvailableCars() {
    const [cars, setCars] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8000/cars')
            .then(response => {
                console.log("Cars data received:", response.data);
                setCars(response.data);
            })
            .catch(error => {
                console.error('Error fetching cars:', error);
            });
    }, []);

    // Функция для корректного формирования пути к изображению
    const getImagePath = (imagePath) => {
        if (!imagePath) return '';
        
        // Если путь уже абсолютный (начинается с /assets)
        if (imagePath.startsWith('/assets/')) {
            return process.env.PUBLIC_URL + imagePath;
        }
        
        // Если путь относительный (без ведущего слеша)
        if (!imagePath.startsWith('/')) {
            return process.env.PUBLIC_URL + '/assets/' + imagePath;
        }
        
        // Для других случаев
        return process.env.PUBLIC_URL + imagePath;
    };

    return (
        <div className="available-cars-container">
            <NavBar>Контакты</NavBar>
            <h2>НОВЫЕ АВТОМОБИЛИ LEXUS</h2>
            <div className="cars-list">
                {cars.map(car => (
                    <div key={car.car_id} className="car-card">
                        {car.image && (
                            <div className="car-image-container">
                                <img 
                                    src={getImagePath(car.image)} 
                                    alt={`${car.trim_level} ${car.year}`}
                                    onError={(e) => {
                                        e.target.style.display = 'none';
                                        console.error('Failed to load image:', car.image);
                                    }}
                                />
                            </div>
                        )}
                        <div className="car-details">
                            <h3>{car.trim_level} {car.year}</h3>
                            <div className="car-price">{car.price}&#8381;</div>
                            <div className="car-color">Цвет: {car.color}</div>
                            <div>Тип кузова: {car.bodywork}</div>
                            <div>Двигатель: {car.engine} л, {car.engine_capacity}</div>
                            <div>Топливо: {car.fuel}</div>
                            {car.description_1 && <div className="car-description">{car.description_1}</div>}
                            {car.description_2 && <div className="car-description">{car.description_2}</div>}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default AvailableCars;