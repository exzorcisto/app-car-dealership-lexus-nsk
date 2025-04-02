import React, { useState } from 'react'
import { Link } from 'react-router-dom';
import './Header.css';
import logo from '../../img/logo.png';
import { FaAngleDown } from "react-icons/fa";

export default function Header() {
  const [faAngleDown, setFaAngleDown] = useState(false)

  return (
    <header>
          <ul className='nav-up'>
            <li><Link to='/'><img className='logo' src={logo} alt="logo"/></Link></li>
            <ul className='nav-up-info'>
              <li>+7 (383) 246-00-00<FaAngleDown onClick={() => setFaAngleDown(!faAngleDown)} className={`FaAngleDown ${faAngleDown ? 'active' : ''}`}/></li>
              <li>Лексус - Новосибирск</li>
            </ul>
          </ul>
          <ul className='nav-menu'>
            <li><Link to='/modelrange'>Модельный ряд</Link></li>
            <li><Link to='/availablecars'>Автомобили в наличии</Link></li>
            <li><Link to='/lexusworld'>Мир Lexus</Link></li>
            <li><Link to='/contacts'>Контакты</Link></li>
          </ul>

    </header>
  )
}
