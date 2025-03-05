import React from 'react'
import './Main.css'
import { CgArrowLongRight } from "react-icons/cg";
import ButtonCustom from '../../UI/button/ButtonCustom';

export default function Main() {
  return (
    <main>
        <div className="container-greetings">
            <div className='text-greetings'>
                <span>Стремлениек к <br/> совершенству <br/> вместе с LEXUS</span>
                <ButtonCustom ButtonCustom onClick={() => console.log('clicking')}>ЗАКАЗАТЬ ОБРАТНЫЙ ЗВОНОК <CgArrowLongRight className='CgArrowLongRight'/></ButtonCustom>
            </div>
        </div>
        <div className='container-greetings-2'>
            <span className='text-greeting-2'>
                Рады видеть вас в Лексус - Новосибирск – официальном дилерском центре легендарной марки Lexus в <br /> 
                Новосибирске. Оцените настоящее японское качество стандартов продажи и обслуживания ООО «Тойота Мотор»!
            </span>
            <div className='block-buttons'>
                <ButtonCustom ButtonCustom onClick={() => console.log('clicking')}>АВТОМОБИЛИ В НАЛИЧИИ</ButtonCustom>
                <ButtonCustom className='cstmBtn-style-1' onClick={() => console.log('clicking')}>ЗАПИСАТЬСЯ НА ТО</ButtonCustom>
            </div>
        </div>
    </main>
  )
}
