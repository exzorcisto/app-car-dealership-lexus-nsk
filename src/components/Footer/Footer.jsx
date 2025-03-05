import React from 'react'
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
            <li><a href='#'>Модельный ряд</a></li>
            <li><a href='#'>Автомобили в наличии</a></li>
            <li><a href='#'>Мир Lexus</a></li>
            <li><a href='#'>Контакты</a></li>
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