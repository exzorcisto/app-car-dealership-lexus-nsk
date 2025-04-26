import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import ButtonCustom from '../../components/UI/ButtonCustom/ButtonCustom';
import './CarDetails.css';

const CarDetails = ({ cars }) => {
    const { id } = useParams();
    const navigate = useNavigate();
    const car = cars?.find(c => c.carid === parseInt(id));

    if (!car) {
        return (
            <div className="car-not-found">
                <h2>Автомобиль не найден</h2>
                <ButtonCustom onClick={() => navigate('/availablecars')}>
                    Вернуться к списку
                </ButtonCustom>
            </div>
        );
    }

    const exteriorFeatures = car.description_1 ? car.description_1.split('\n').filter(item => item.trim() !== '') : [];
    const comfortFeatures = car.description_2 ? car.description_2.split('\n').filter(item => item.trim() !== '') : [];

    return (
        <div className="car-details-container">
            {/* Шапка с изображением и основной информацией */}
            <div className="car-header-section">
                <div className="header-content">
                    <h1>LEXUS {car.model_name} {car.trim_level}</h1>
                    <div className="availability-date">Доступен к заказу • {new Date().toLocaleDateString('ru-RU')}</div>
                </div>
                
                <div className="price-action-container">
                    <div className="price-value">{car.price.toLocaleString('ru-RU')} ₽</div>
                    <ButtonCustom className="offer-btn">
                        ПОЛУЧИТЬ ПРЕДЛОЖЕНИЕ
                    </ButtonCustom>
                </div>
            </div>

            {/* Изображение автомобиля */}
            {car.image && (
                <div className="car-image-container">
                    <img 
                        src={car.image.startsWith('/assets/') 
                            ? process.env.PUBLIC_URL + car.image 
                            : !car.image.startsWith('/') 
                                ? process.env.PUBLIC_URL + '/assets/' + car.image 
                                : process.env.PUBLIC_URL + car.image} 
                        alt={`${car.model_name} ${car.trim_level} ${car.year}`}
                    />
                </div>
            )}

            {/* Характеристики */}
            <div className="specs-section">
                <h2>ХАРАКТЕРИСТИКИ</h2>
                <div className="specs-grid">
                    <div className="spec-row">
                        <span className="spec-name">Цвет кузова</span>
                        <span className="spec-value">{car.color}</span>
                    </div>
                    <div className="spec-row">
                        <span className="spec-name">Дата производства</span>
                        <span className="spec-value">{car.year}</span>
                    </div>
                    <div className="spec-row">
                        <span className="spec-name">Кузов</span>
                        <span className="spec-value">{car.bodywork}</span>
                    </div>
                    <div className="spec-row">
                        <span className="spec-name">Двигатель</span>
                        <span className="spec-value">{car.engine_capacity} л.с.</span>
                    </div>
                    <div className="spec-row">
                        <span className="spec-name">Объем двигателя</span>
                        <span className="spec-value">{car.engine} л</span>
                    </div>
                    <div className="spec-row">
                        <span className="spec-name">Топливо</span>
                        <span className="spec-value">{car.fuel}</span>
                    </div>
                </div>
            </div>

            {/* Комплектация */}
            <div className="features-section">
                <h2>КОМПЛЕКТАЦИЯ</h2>
                <div className="features-grid">
                    <div className="feature-group">
                        <h3>ЭКСТЕРЬЕР</h3>
                        <ul className="features-list">
                            {exteriorFeatures.map((feature, index) => (
                                <li key={`exterior-${index}`}>{feature}</li>
                            ))}
                        </ul>
                    </div>
                    <div className="feature-group">
                        <h3>КОМФОРТ</h3>
                        <ul className="features-list">
                            {comfortFeatures.map((feature, index) => (
                                <li key={`comfort-${index}`}>{feature}</li>
                            ))}
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default CarDetails;