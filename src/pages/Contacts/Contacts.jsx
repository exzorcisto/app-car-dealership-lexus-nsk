import React from 'react'
import './Contacts.css'
import NavBar from '../../components/UI/NavBar/NavBar'
import AddCar from '../../components/AddCar/AddCar'


export default function Contacts() {
    return (
        <main>
            <NavBar>Контакты</NavBar>
            <AddCar />
        </main>
    )
}
