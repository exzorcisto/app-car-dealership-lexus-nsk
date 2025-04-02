import React from 'react'
import './Homepage.css';
import {CgArrowLongRight} from "react-icons/cg";
import ButtonCustom from '../../components/UI/ButtonCustom/ButtonCustom';


export default function Homepage() {
    return (
        <main>
            <div className="container-greetings">
                <div className='text-greetings'>
                    <span>Стремлениек к <br/> совершенству <br/> вместе с LEXUS</span>
                    <ButtonCustom onClick={() => console.log('clicking')}>ЗАКАЗАТЬ ОБРАТНЫЙ
                        ЗВОНОК <CgArrowLongRight className='CgArrowLongRight'/></ButtonCustom>
                </div>
            </div>
            <div className='container-greetings-2'>
                <span className='text-greeting-2'>
                    Рады видеть вас в Лексус - Новосибирск – официальном дилерском центре легендарной марки Lexus в <br/>
                    Новосибирске. Оцените настоящее японское качество стандартов продажи и обслуживания ООО «Тойота Мотор»!
                </span>
                <div className='block-buttons'>
                    <ButtonCustom className='btn-style-1' onClick={() => console.log('clicking')}>АВТОМОБИЛИ В
                        НАЛИЧИИ <CgArrowLongRight className='CgArrowLongRight'/></ButtonCustom>
                    <ButtonCustom className='btn-style-2' onClick={() => console.log('clicking')}>ЗАПИСАТЬСЯ НА
                        ТО <CgArrowLongRight className='CgArrowLongRight'/></ButtonCustom>
                </div>
            </div>
        </main>
    )
}
