// frontend/src/pages/AddCar/AddCar.jsx
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import './AddCar.css';

function AddCar() {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        model_name: '', // New field
        trim_level: '',
        year: '',
        vin: '',
        price: '',
        color: '',
        bodywork: '',
        engine: '',
        engine_capacity: '',
        fuel: '',
        image: '',
        description_1: '',
        description_2: ''
    });
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [error, setError] = useState('');

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsSubmitting(true);
        setError('');
    
        // Валидация перед отправкой
        if (!formData.model_name || !formData.trim_level || !formData.vin || !formData.price) {
            setError('Заполните все обязательные поля');
            setIsSubmitting(false);
            return;
        }
    
        try {
            // Подготовка данных изображения
            let imagePath = formData.image;
    
            const response = await axios.post('http://localhost:8000/cars', {
                model_name: formData.model_name, // New field
                trim_level: formData.trim_level,
                year: parseInt(formData.year) || 2023,
                vin: formData.vin,
                price: parseFloat(formData.price) || 0,
                color: formData.color,
                bodywork: formData.bodywork,
                engine: parseFloat(formData.engine) || 2.0,
                engine_capacity: formData.engine_capacity || '150 л.с.',
                fuel: formData.fuel || 'Бензин',
                image: imagePath || '',
                description_1: formData.description_1,
                description_2: formData.description_2
            }, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
    
            if (response.status === 201) {
                alert('Автомобиль успешно добавлен!');
                navigate('/cars');
            }
        } catch (err) {
            const serverError = err.response?.data?.message || err.message;
            setError(`Ошибка: ${serverError}`);
            console.error('Детали ошибки:', {
                error: err,
                request: err.config,
                response: err.response
            });
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="add-car-container">
            <h2>Добавить новый автомобиль</h2>
            {error && <div className="error-message">{error}</div>}
            
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>Модель автомобиля:</label>
                    <select
                        name="model_name"
                        value={formData.model_name}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите модель</option>
                        <option value="ES 250">ES 250</option>
                        <option value="NX 200">NX 200</option>
                        <option value="RX 350">RX 350</option>
                        <option value="GX">GX</option>
                        <option value="LX">LX</option>
                        <option value="LM">LM</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Комплектация:</label>
                    <select
                        name="trim_level"
                        value={formData.trim_level}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите комплектацию</option>
                        <option value="PROGRESSIVE">PROGRESSIVE</option>
                        <option value="PREMIUM">PREMIUM</option>
                        <option value="EXECUTIVE">EXECUTIVE</option>
                        <option value="LUXURY">LUXURY</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Год выпуска:</label>
                    <input
                        type="number"
                        name="year"
                        value={formData.year}
                        onChange={handleChange}
                        min="2020"
                        max={new Date().getFullYear() + 1}
                        required
                    />
                </div>

                <div className="form-group">
                    <label>VIN номер:</label>
                    <input
                        type="text"
                        name="vin"
                        value={formData.vin}
                        onChange={handleChange}
                        required
                        pattern="[A-HJ-NPR-Z0-9]{17}"
                        title="Введите 17-значный VIN номер"
                    />
                </div>

                <div className="form-group">
                    <label>Цена (руб):</label>
                    <input
                        type="number"
                        name="price"
                        value={formData.price}
                        onChange={handleChange}
                        step="1000"
                        min="0"
                        required
                    />
                </div>

                <div className="form-group">
                    <label>Цвет:</label>
                    <select
                        name="color"
                        value={formData.color}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите цвет</option>
                        <option value="ПЛАТИНОВЫЙ МЕТАЛИК">ПЛАТИНОВЫЙ МЕТАЛЛИК</option>
                        <option value="ЧЕРНЫЙ НЕМЕТАЛЛИК">ЧЕРНЫЙ НЕМЕТАЛЛИК</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Тип кузова:</label>
                    <select
                        name="bodywork"
                        value={formData.bodywork}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите тип кузова</option>
                        <option value="Седан">Седан</option>
                        <option value="Внедорожник">Внедорожник</option>
                        <option value="Кроссовер">Кроссовер</option>
                        <option value="Минивэн">Минивэн</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Объем двигателя (л):</label>
                    <input
                        type="number"
                        name="engine"
                        value={formData.engine}
                        onChange={handleChange}
                        step="0.1"
                        min="1.0"
                        max="6.0"
                        required
                    />
                </div>

                <div className="form-group">
                    <label>Мощность двигателя (л.с.):</label>
                    <input
                        type="number"
                        name="engine_capacity"
                        value={formData.engine_capacity}
                        onChange={handleChange}
                        min="100"
                        max="1000"
                        required
                    />
                </div>

                <div className="form-group">
                    <label>Тип топлива:</label>
                    <select
                        name="fuel"
                        value={formData.fuel}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите тип топлива</option>
                        <option value="Бензин">Бензин</option>
                        <option value="Дизель">Дизель</option>
                        <option value="Гибрид">Гибрид</option>
                        <option value="Электро">Электро</option>
                    </select>
                </div>

                <div className="form-group">
                    <label>Описание 1:</label>
                    <textarea
                        name="description_1"
                        value={formData.description_1}
                        onChange={handleChange}
                        rows="3"
                        placeholder="Введите первое описание автомобиля"
                    />
                </div>

                <div className="form-group">
                    <label>Описание 2:</label>
                    <textarea
                        name="description_2"
                        value={formData.description_2}
                        onChange={handleChange}
                        rows="3"
                        placeholder="Введите второе описание автомобиля"
                    />
                </div>

                <div className="form-group">
                    <label>Путь к изображению:</label>
                    <input
                        type="text"
                        name="image"
                        value={formData.image}
                        onChange={handleChange}
                        placeholder="Например: /assets/imgcars/es/es250platinumFront.png"
                    />
                    <small className="hint">
                        Используйте относительные пути из папки /assets/imgcars/
                    </small>
                    {formData.image && (
                        <div className="image-preview">
                            <img 
                                src={formData.image.startsWith('/') ? 
                                    process.env.PUBLIC_URL + formData.image : 
                                    formData.image
                                } 
                                alt="Предпросмотр"
                                onError={(e) => {
                                    e.target.style.display = 'none';
                                    setError('Изображение по указанному пути не найдено');
                                }}
                            />
                        </div>
                    )}
                </div>

                <div className="form-actions">
                    <button 
                        type="submit" 
                        disabled={isSubmitting}
                        className={isSubmitting ? 'submitting' : ''}
                    >
                        {isSubmitting ? 'Добавление...' : 'Добавить автомобиль'}
                    </button>
                    <button 
                        type="button" 
                        className="cancel-btn"
                        onClick={() => navigate('/cars')}
                    >
                        Отмена
                    </button>
                </div>
            </form>
        </div>
    );
}

export default AddCar;