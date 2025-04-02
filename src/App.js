import {Route, Routes} from 'react-router-dom';

import {Layout} from './components/Layout';
import Homepage from "./pages/Homepage/Homepage";
import ModelRange from "./pages/ModelRange/ModelRange";
import AvailableCars from "./pages/AvailableCars/AvailableCars"; // Импортируйте компонент
import LexusWorld from "./pages/LexusWorld/LexusWorld"; // Импортируйте компонент
import Contacts from "./pages/Contacts/Contacts"; // Импортируйте компонент
import './index.css';

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route index element={<Homepage />} />
          <Route path='modelrange' element={<ModelRange />} />
          <Route path='availablecars' element={<AvailableCars />} /> {/* Добавьте Route */}
          <Route path='lexusworld' element={<LexusWorld />} /> {/* Добавьте Route */}
          <Route path='contacts' element={<Contacts />} /> {/* Добавьте Route */}
        </Route>
      </Routes>
    </>
  );
}

export default App;
