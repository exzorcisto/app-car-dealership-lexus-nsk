import React, { useState, useEffect, useMemo } from 'react';
import { Link } from 'react-router-dom';
import './AvailableCars.css';
import NavBar from '../../components/UI/NavBar/NavBar';
import axios from 'axios';
import ButtonCustom from '../../components/UI/ButtonCustom/ButtonCustom';
import { FaCircle, FaFilter, FaTimes } from "react-icons/fa";

function AvailableCars() {
    const [cars, setCars] = useState([]);
    const [filters, setFilters] = useState({
        priceMin: '',
        priceMax: '',
        color: '',
        engine: '',
        bodywork: '',
        fuel: ''
    });
    const [sortOption, setSortOption] = useState('price_asc');
    const [showFilters, setShowFilters] = useState(false);

    useEffect(() => {
        const fetchCars = async () => {
            try {
                const response = await axios.get('http://localhost:8000/cars');
                setCars(response.data);
            } catch (error) {
                console.error('Ошибка при загрузке автомобилей:', error);
            }
        };
        fetchCars();
    }, []);

    const filterOptions = useMemo(() => ({
        colors: [...new Set(cars.map(car => car.color))],
        engines: [...new Set(cars.map(car => car.engine.toString()))],
        bodyworks: [...new Set(cars.map(car => car.bodywork))],
        fuels: [...new Set(cars.map(car => car.fuel))]
    }), [cars]);

    const filteredAndSortedCars = useMemo(() => {
        let result = [...cars];

        if (filters.priceMin) result = result.filter(car => car.price >= Number(filters.priceMin) * 1000000);
        if (filters.priceMax) result = result.filter(car => car.price <= Number(filters.priceMax) * 1000000);
        if (filters.color) result = result.filter(car => car.color === filters.color);
        if (filters.engine) result = result.filter(car => car.engine.toString() === filters.engine);
        if (filters.bodywork) result = result.filter(car => car.bodywork === filters.bodywork);
        if (filters.fuel) result = result.filter(car => car.fuel === filters.fuel);

        switch (sortOption) {
            case 'price_asc': return result.sort((a, b) => a.price - b.price);
            case 'price_desc': return result.sort((a, b) => b.price - a.price);
            case 'name_asc': return result.sort((a, b) => a.model_name.localeCompare(b.model_name));
            case 'year_desc': return result.sort((a, b) => b.year - a.year);
            default: return result;
        }
    }, [cars, filters, sortOption]);

    const handleFilterChange = (e) => {
        const { name, value } = e.target;
        setFilters(prev => ({ ...prev, [name]: value }));
    };

    const resetFilters = () => {
        setFilters({
            priceMin: '',
            priceMax: '',
            color: '',
            engine: '',
            bodywork: '',
            fuel: ''
        });
    };

    return (
        <div className="available-cars-container">
            <NavBar>Контакты</NavBar>
            
            <div className="main-content">
                {/* Панель фильтров слева */}
                {showFilters && (
                    <div className="filters-sidebar">
                        <div className="filters-header">
                            <h3>ФИЛЬТРЫ</h3>
                            <button 
                                className="close-filters" 
                                onClick={() => setShowFilters(false)}
                                aria-label="Закрыть фильтры"
                            >
                                <FaTimes />
                            </button>
                        </div>
                        
                        <FilterGroup title="ЦЕНА, ₽">
                            <div className="price-range">
                                <div className="range-inputs">
                                    <input
                                        type="number"
                                        name="priceMin"
                                        value={filters.priceMin}
                                        onChange={handleFilterChange}
                                        placeholder="от"
                                        min="0"
                                        step="0.1"
                                    />
                                    <span>-</span>
                                    <input
                                        type="number"
                                        name="priceMax"
                                        value={filters.priceMax}
                                        onChange={handleFilterChange}
                                        placeholder="до"
                                        min="0"
                                        step="0.1"
                                    />
                                </div>
                            </div>
                        </FilterGroup>

                        <FilterSelect 
                            title="ЦВЕТ"
                            name="color"
                            value={filters.color}
                            options={filterOptions.colors}
                            onChange={handleFilterChange}
                            placeholder="Все цвета"
                        />

                        <FilterSelect 
                            title="ОБЪЁМ ДВИГАТЕЛЯ, л"
                            name="engine"
                            value={filters.engine}
                            options={filterOptions.engines}
                            onChange={handleFilterChange}
                            placeholder="Все"
                        />

                        <FilterSelect 
                            title="ТИП КУЗОВА"
                            name="bodywork"
                            value={filters.bodywork}
                            options={filterOptions.bodyworks}
                            onChange={handleFilterChange}
                            placeholder="Все"
                        />

                        <FilterSelect 
                            title="ТИП ТОПЛИВА"
                            name="fuel"
                            value={filters.fuel}
                            options={filterOptions.fuels}
                            onChange={handleFilterChange}
                            placeholder="Все"
                        />

                        <div className="filter-actions">
                            <button className="reset-btn" onClick={resetFilters}>
                                Сбросить все
                            </button>
                            <button 
                                className="apply-btn" 
                                onClick={() => setShowFilters(false)}
                            >
                                Применить
                            </button>
                        </div>
                    </div>
                )}

                {/* Основной контент справа */}
                <div className="cars-list-container">
                    <div className="cars-header">
                        <h2>АВТОМОБИЛИ LEXUS</h2>
                        <div className="controls">
                            <div className="sort-options">
                                <select 
                                    value={sortOption} 
                                    onChange={(e) => setSortOption(e.target.value)}
                                    className="sort-select"
                                >
                                    <option value="price_asc">По возрастанию цены</option>
                                    <option value="price_desc">По убыванию цены</option>
                                    <option value="name_asc">По названию</option>
                                    <option value="year_desc">По году выпуска</option>
                                </select>
                            </div>
                            <button 
                                className={`filter-toggle ${showFilters ? 'active' : ''}`}
                                onClick={() => setShowFilters(!showFilters)}
                            >
                                {showFilters ? <FaTimes /> : <FaFilter />}
                                Фильтры
                            </button>
                        </div>
                    </div>

                    <div className="cars-count">
                        Найдено: <strong>{filteredAndSortedCars.length}</strong> автомобилей
                    </div>

                    {filteredAndSortedCars.length > 0 ? (
                        filteredAndSortedCars.map(car => (
                            <CarCard key={`car-${car.carid}`} car={car} />
                        ))
                    ) : (
                        <div className="no-results">
                            По выбранным фильтрам автомобилей не найдено
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

const CarCard = ({ car }) => {
    if (!car) return null;
    
    return (
        <div className="car-card">
            <div className='car-card-left'>
            <Link to={`/availablecars/${car.carid}`} className="car-link">
                {car.image && (
                    <div className="car-image-container">
                        <img 
                            src={car.image.startsWith('/assets/') 
                                ? process.env.PUBLIC_URL + car.image 
                                : !car.image.startsWith('/') 
                                    ? process.env.PUBLIC_URL + '/assets/' + car.image 
                                    : process.env.PUBLIC_URL + car.image} 
                            alt={`${car.model_name} ${car.trim_level} ${car.year}`}
                            onError={(e) => e.target.style.display = 'none'}
                        />
                    </div>
                )}
                <div className="car-details">
                    <div className="car-model">{car.model_name} {car.trim_level} {car.year}</div>
                    <div className="car-available"><FaCircle /> Доступен к заказу</div>
                    <div className="car-color">{car.color}</div>
                    <div className="car-specs">
                        <span>{car.engine} л</span>
                        <span>{car.bodywork}</span>
                        <span>{car.fuel}</span>
                    </div>
                </div>
            </Link>
        </div>
        <div className="car-card-right">
                <div className="car-price">{car.price.toLocaleString('ru-RU')} ₽</div>
                <ButtonCustom 
                    className='btn-style-1' 
                    onClick={() => window.location.href = `/availablecars/${car.carid}`}
                >
                    ПОЛУЧИТЬ ПРЕДЛОЖЕНИЕ
                </ButtonCustom>
            </div>
        </div>
    );
};

const FilterGroup = ({ title, children }) => (
    <div className="filter-group">
        <h4>{title}</h4>
        {children}
    </div>
);

const FilterSelect = ({ title, name, value, options, onChange, placeholder }) => (
    <FilterGroup title={title}>
        <select
            name={name}
            value={value}
            onChange={onChange}
            className="filter-select"
        >
            <option value="">{placeholder}</option>
            {options.map(option => (
                <option key={option} value={option}>{option}</option>
            ))}
        </select>
    </FilterGroup>
);

export default AvailableCars;