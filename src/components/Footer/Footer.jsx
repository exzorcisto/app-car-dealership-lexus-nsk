import React from 'react'
import { Link } from 'react-router-dom';
import './Footer.css';
import logo from '../../img/logo.png';

export default function Footer() {
  return (
    <footer>
        <div className='map-nav'>
          <ul className='map-up'>
            <li><img className='logo' src={logo} alt="logo"/></li>
            <ul className='map-up-nav'>
              <li><a href='#'>Новосибирск, ул. Большевистская, 276/2</a></li>
              <li><a href='#'>+7 (383) 246-00-00</a></li>
              <li><a href='#'>Лексус - Новосибирск</a></li>
            </ul>
          </ul>
          <ul className='map-mid'>
            <li><Link to='/modelrange'>Модельный ряд</Link></li>
            <li><Link to='/availablecars'>Автомобили в наличии</Link></li>
            <li><Link to='/lexusworld'>Мир Lexus</Link></li>
            <li><Link to='/contacts'>Контакты</Link></li>
          </ul>
          <div>
            Вся представленная на сайте информация, касающаяся стоимости автомобилей, 
            аксессуаров и сервисного обслуживания, носит информационный характер и не является 
            публичной офертой, определяемой положениями ст. 437 (2) ГК РФ. Для получения подробной 
            информации обращайтесь в официальные дилерские центры. Опубликованная на данном сайте 
            информация может быть изменена в любое время без предварительного уведомления.
          </div>
          <div>&copy; 2025, ООО Лексус Моторс</div>
        </div>
    </footer>
  )
}