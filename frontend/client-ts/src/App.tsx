import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { useRoutes } from './route';


const App: React.FC = () => {


  const routes = useRoutes(false)
  return (
    <BrowserRouter>
          {routes}
    </BrowserRouter>  
  );
}

export default App;