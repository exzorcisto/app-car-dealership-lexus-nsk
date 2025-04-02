import {Route, Routes} from 'react-router-dom';

import {Layout} from './components/Layout';
import Homepage from "./pages/Homepage/Homepage";
import ModelRange from "./pages/ModelRange/ModelRange";
import AvailableCars from "./pages/AvailableCars/AvailableCars";
import LexusWorld from "./pages/LexusWorld/LexusWorld";
import Contacts from "./pages/Contacts/Contacts";
import './index.css';

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route index element={<Homepage />} />
          <Route path='modelrange' element={<ModelRange />} />
          <Route path='availablecars' element={<AvailableCars />} />
          <Route path='lexusworld' element={<LexusWorld />} />
          <Route path='contacts' element={<Contacts />} />
        </Route>
      </Routes>
    </>
  );
}

export default App;
