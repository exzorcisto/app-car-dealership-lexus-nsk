import React, { useState } from 'react'
import './Header.css';
import logo from '../../img/logo.png';
import { FaAngleDown } from "react-icons/fa";

export default function Header() {
  const [faAngleDown, setFaAngleDown] = useState(false)

  return (
    <header>
          <ul className='nav-up'>
            <li><img className='logo' src={logo} alt="logo"/></li>
            <ul className='nav-up-info'>
              <li>+7 (383) 246-00-00<FaAngleDown onClick={() => setFaAngleDown(!faAngleDown)} className={`FaAngleDown ${faAngleDown ? 'active' : ''}`}/></li>
              <li>Лексус - Новосибирск</li>
            </ul>
          </ul>
          <ul className='nav-menu'>
            <li><a href='#'>Модельный ряд</a></li>
            <li><a href='#'>Автомобили в наличии</a></li>
            <li><a href='#'>Мир Lexus</a></li>
            <li><a href='#'>Контакты</a></li>
          </ul>
    </header>
  )
}
